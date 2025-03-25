import { check } from 'k6';
import http from 'k6/http';

export const options = {
    vus: 10,
    duration: '30s',
    thresholds: {
        http_req_failed: ['rate<0.01'],
        http_req_duration: ['p(95)<500'],
    },
};

export default function () {
    const countryCode = 'AL';
    const url = `http://localhost:8080/v1/swift-codes/country/${countryCode}`;

    const response = http.get(url, {
        tags: { name: 'GetSwiftCodesByCountry' },
    });

    check(response, {
        'is status 200': (r) => r.status === 200,
        'has correct content type': (r) => r.headers['Content-Type']?.includes('application/json'),
    });

    if (response.status === 200) {
        let body;
        try {
            body = response.json();
        } catch (e) {
            console.error(`JSON parse error: ${e}, response body: ${response.body}`);
            return;
        }

        if (!body || typeof body !== 'object') {
            console.error('Invalid response body format');
            return;
        }


        check(body, {
            'has country ISO code': (b) => b.countryISO2 === countryCode,
            'has country name': (b) => b.countryName === 'ALBANIA',
            'has swift codes array': (b) => Array.isArray(b.swiftCodes) && b.swiftCodes.length > 0,
        });
        if (body.swiftCodes && body.swiftCodes.length > 0) {
            const firstEntry = body.swiftCodes[0];

            check(firstEntry, {
                'swift code has address': (e) => typeof e.address === 'string',
                'swift code has bank name': (e) => typeof e.bankName === 'string',
                'swift code has country code': (e) => e.countryISO2 === countryCode,
                'swift code has valid SWIFT code': (e) => /^[A-Z]{6}[A-Z0-9]{2}([A-Z0-9]{3})?$/.test(e.swiftCode),
            });
        }
    }
}