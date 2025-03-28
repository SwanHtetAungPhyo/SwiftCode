#!/usr/bin/env bash
curl -X GET "http://localhost:8080/v1/swift-codes/country/PL" | jq . > ./result_json/response_get_by_country.json

curl -X GET "http://localhost:8080/v1/swift-codes/BSCHCLRMXXX" | jq . > ./result_json/response_get_by_swift_code.json

curl -X POST "http://localhost:8080/v1/swift-codes" \
-H "Content-Type: application/json" \
-d '{
      "address": "123 Bank St.",
      "bankName": "Some Bank",
      "countryISO2": "US",
      "countryName": "United States",
      "isHeadquarter": true,
      "swiftCode": "SOMEBANK01"
    }' | jq . > ./result_json/response_post_create_swift_code.json
