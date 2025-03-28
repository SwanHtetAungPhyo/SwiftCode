
--- AUTHOR : SWAN HTET AUNG PHYO
--- This is th presentation in plain sql  with transaction
BEGIN;

CREATE TABLE countries (
                           id BIGSERIAL PRIMARY KEY,
                           country_iso2_code VARCHAR(2) UNIQUE NOT NULL,
                           name TEXT UNIQUE NOT NULL,
                           time_zone TEXT,
                           created_at TIMESTAMPTZ DEFAULT now(),
                           updated_at TIMESTAMPTZ DEFAULT now()
);


CREATE TABLE towns (
                       id BIGSERIAL PRIMARY KEY,
                       name TEXT NOT NULL,
                       country_id BIGINT NOT NULL,
                       created_at TIMESTAMPTZ DEFAULT now(),
                       updated_at TIMESTAMPTZ DEFAULT now(),
                       CONSTRAINT fk_towns_country FOREIGN KEY (country_id) REFERENCES countries(id) ON DELETE CASCADE
);
CREATE INDEX idx_towns_country_id ON towns(country_id);

CREATE TABLE swiftcodes (
                            id BIGSERIAL PRIMARY KEY,
                            name TEXT NOT NULL,
                            address TEXT,
                            swift_code VARCHAR(11) UNIQUE NOT NULL,
                            is_headquarter BOOLEAN DEFAULT FALSE,
                            country_id BIGINT NOT NULL,
                            town_name_id BIGINT NOT NULL,
                            code_type TEXT,
                            created_at TIMESTAMPTZ DEFAULT now(),
                            updated_at TIMESTAMPTZ DEFAULT now(),
                            CONSTRAINT fk_swiftcodes_country FOREIGN KEY (country_id) REFERENCES countries(id) ON DELETE CASCADE,
                            CONSTRAINT fk_swiftcodes_town FOREIGN KEY (town_name_id) REFERENCES towns(id) ON DELETE CASCADE
);
CREATE INDEX idx_swiftcodes_country_id ON swiftcodes(country_id);
CREATE INDEX idx_swiftcodes_town_name_id ON swiftcodes(town_name_id);


-- Insert Countries
INSERT INTO public.countries (country_iso2_code, name, time_zone)
VALUES
    ('AL', 'ALBANIA', 'Europe/Tirane'),
    ('BG', 'BULGARIA', 'Europe/Sofia'),
    ('UY', 'URUGUAY', 'America/Montevideo'),
    ('MC', 'MONACO', 'Europe/Monaco'),
    ('PL', 'POLAND', 'Europe/Warsaw'),
    ('LV', 'LATVIA', 'Europe/Riga'),
    ('MT', 'MALTA', 'Europe/Malta'),
    ('AW', 'ARUBA', 'America/Aruba'),
    ('CL', 'CHILE', 'Pacific/Easter');

-- Insert Towns
INSERT INTO public.towns (name, country_id)
VALUES
    ('TIRANA', (SELECT id FROM public.countries WHERE country_iso2_code = 'AL')),
    ('VARNA', (SELECT id FROM public.countries WHERE country_iso2_code = 'BG')),
    ('SOFIA', (SELECT id FROM public.countries WHERE country_iso2_code = 'BG')),
    ('MONTEVIDEO', (SELECT id FROM public.countries WHERE country_iso2_code = 'UY')),
    ('MONACO', (SELECT id FROM public.countries WHERE country_iso2_code = 'MC')),
    ('WROCLAW', (SELECT id FROM public.countries WHERE country_iso2_code = 'PL')),
    ('RIGA', (SELECT id FROM public.countries WHERE country_iso2_code = 'LV')),
    ('ST. JULIAN', (SELECT id FROM public.countries WHERE country_iso2_code = 'MT')),
    ('ORANJESTAD', (SELECT id FROM public.countries WHERE country_iso2_code = 'AW')),
    ('SANTIAGO', (SELECT id FROM public.countries WHERE country_iso2_code = 'CL'));

-- Insert Swiftcodes
INSERT INTO public.swiftcodes (name, address, swift_code, is_headquarter, country_id, town_name_id, code_type)
VALUES
    ('UNITED BANK OF ALBANIA SH.A.', 'HYRJA 3 RR. DRITAN HOXHA ND. 11 TIRANA, TIRANA, 1023', 'AAISALTRXXX', TRUE,
     (SELECT id FROM public.countries WHERE country_iso2_code = 'AL'),
     (SELECT id FROM public.towns WHERE name = 'TIRANA'), 'BIC11'),
    ('ABV INVESTMENTS LTD', 'TSAR ASEN 20 VARNA, VARNA, 9002', 'ABIEBGS1XXX', FALSE,
     (SELECT id FROM public.countries WHERE country_iso2_code = 'BG'),
     (SELECT id FROM public.towns WHERE name = 'VARNA'), 'BIC11'),
    ('ADAMANT CAPITAL PARTNERS AD', 'JAMES BOURCHIER BLVD 76A HILL TOWER SOFIA, SOFIA, 1421', 'ADCRBGS1XXX', FALSE,
     (SELECT id FROM public.countries WHERE country_iso2_code = 'BG'),
     (SELECT id FROM public.towns WHERE name = 'SOFIA'), 'BIC11'),
    ('AFINIDAD A.F.A.P.S.A.', 'PLAZA INDEPENDENCIA 743 MONTEVIDEO, MONTEVIDEO, 11000', 'AFAAUYM1XXX', FALSE,
     (SELECT id FROM public.countries WHERE country_iso2_code = 'UY'),
     (SELECT id FROM public.towns WHERE name = 'MONTEVIDEO'), 'BIC11'),
    ('CREDIT AGRICOLE MONACO (CRCA PROVENCE COTE D AZUR MONACO)', '23 BOULEVARD PRINCESSE CHARLOTTE MONACO, MONACO, 98000',
     'AGRIMCM1XXX', FALSE, (SELECT id FROM public.countries WHERE country_iso2_code = 'MC'),
     (SELECT id FROM public.towns WHERE name = 'MONACO'), 'BIC11'),
    ('SANTANDER CONSUMER BANK SPOLKA AKCYJNA', 'STRZEGOMSKA 42C WROCLAW, DOLNOSLASKIE, 53-611', 'AIPOPLP1XXX', FALSE,
     (SELECT id FROM public.countries WHERE country_iso2_code = 'PL'),
     (SELECT id FROM public.towns WHERE name = 'WROCLAW'), 'BIC11'),
    ('ABLV BANK, AS IN LIQUIDATION', 'MIHAILA TALA STREET 1 RIGA, RIGA, LV-1045', 'AIZKLV22XXX', FALSE,
     (SELECT id FROM public.countries WHERE country_iso2_code = 'LV'),
     (SELECT id FROM public.towns WHERE name = 'RIGA'), 'BIC11'),
    ('AKBANK T.A.S. (MALTA BRANCH)', 'FLOOR 6, PORTOMASO BUSINESS TOWER 01 PORTOMASO PTM - ST. JULIAN\ ST. JULIA, STJ 4011',
     'AKBKMTMTXXX', FALSE, (SELECT id FROM public.countries WHERE country_iso2_code = 'MT'),
     (SELECT id FROM public.towns WHERE name = 'ST. JULIAN'), 'BIC11');

COMMIT ;


BEGIN;


--- To get the information by swiftcode
SELECT
    s.address,
    s.name,
    c.name,
    c.country_iso2_code,
    s.is_headquarter
FROM
    swiftcodes s
JOIN
        countries c on c.id = s.country_id
WHERE
    s.swift_code = 'AGRIMCM1XXX';


--- To get the information by country ISO 2 Code
SELECT
    c.name,
    c.country_iso2_code,
    s.address,
    s.name,
    s.swift_code,
    s.is_headquarter
FROM
    countries c
JOIN
        swiftcodes s on c.id = s.country_id
WHERE country_iso2_code = 'BG';


COMMIT ;