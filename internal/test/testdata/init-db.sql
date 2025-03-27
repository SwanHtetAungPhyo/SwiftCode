create table countries
(
    id                bigserial
        primary key,
    country_iso2_code varchar(2)
        constraint uni_countries_country_iso2_code
            unique,
    name              text
        constraint uni_countries_name
            unique,
    time_zone         text,
    created_at        timestamp with time zone,
    updated_at        timestamp with time zone
);


create table towns
(
    id         bigserial
        primary key,
    name       text,
    country_id bigint,
    created_at timestamp with time zone,
    updated_at timestamp with time zone
);



create index idx_towns_country_id
    on towns (country_id);

create table swiftcodes
(
    id             bigserial
        primary key,
    name           text,
    address        text,
    swift_code     varchar(11)
        constraint uni_swiftcodes_swift_code
            unique,
    is_headquarter boolean,
    country_id     bigint,
    town_name_id   bigint,
    code_type      text,
    created_at     timestamp with time zone,
    updated_at     timestamp with time zone
);



create index idx_swiftcodes_town_name_id
    on swiftcodes (town_name_id);

create index idx_swiftcodes_country_id
    on swiftcodes (country_id);


--
-- -- Inserting countries data
-- INSERT INTO countries (id, country_iso2_code, name, time_zone) VALUES
--                                                                    (1, 'AL', 'ALBANIA', 'Europe/Tirane'),
--                                                                    (2, 'BG', 'BULGARIA', 'Europe/Sofia'),
--                                                                    (3, 'UY', 'URUGUAY', 'America/Montevideo'),
--                                                                    (4, 'MC', 'MONACO', 'Europe/Monaco'),
--                                                                    (5, 'PL', 'POLAND', 'Europe/Warsaw'),
--                                                                    (6, 'LV', 'LATVIA', 'Europe/Riga'),
--                                                                    (7, 'MT', 'MALTA', 'Europe/Malta'),
--                                                                    (8, 'AW', 'ARUBA', 'America/Aruba'),
--                                                                    (9, 'CL', 'CHILE', 'Pacific/Easter');
--
-- -- Inserting towns data
-- INSERT INTO towns (id, name, country_id) VALUES
--                                              (1, 'TIRANA', 1),
--                                              (2, 'VARNA', 2),
--                                              (3, 'SOFIA', 2),
--                                              (4, 'MONTEVIDEO', 3),
--                                              (5, 'MONACO', 4),
--                                              (6, 'WROCLAW', 5),
--                                              (7, 'RIGA', 6),
--                                              (8, 'ST. JULIAN', 7),
--                                                 (9, 'ORANJESTAD', 8),
-- (10, 'SANTIAGO', 9);
--
-- -- Inserting bank details data
-- INSERT INTO bank_details (id, name, address, swift_code, is_headquarter, country_id, town_name_id) VALUES
-- (1, 'UNITED BANK OF ALBANIA SH.A.', 'HYRJA 3 RR. DRITAN HOXHA ND. 11 TIRANA, TIRANA, 1023', 'AAISALTRXXX', TRUE, 1, 1),
-- (2, 'ABV INVESTMENTS LTD', 'TSAR ASEN 20 VARNA, VARNA, 9002', 'ABIEBGS1XXX', FALSE, 2, 2),
-- (3, 'ADAMANT CAPITAL PARTNERS AD', 'JAMES BOURCHIER BLVD 76A HILL TOWER SOFIA, SOFIA, 1421', 'ADCRBGS1XXX', FALSE, 2, 3),
-- (4, 'AFINIDAD A.F.A.P.S.A.', 'PLAZA INDEPENDENCIA 743 MONTEVIDEO, MONTEVIDEO, 11000', 'AFAAUYM1XXX', FALSE, 3, 4),
-- (5, 'CREDIT AGRICOLE MONACO (CRCA PROVENCE COTE D AZUR MONACO)', '23 BOULEVARD PRINCESSE CHARLOTTE MONACO, MONACO, 98000', 'AGRIMCM1XXX', TRUE, 4, 5),
--                                               (6, 'SANTANDER CONSUMER BANK SPOLKA AKCYJNA', 'STRZEGOMSKA 42C WROCLAW, DOLNOSLASKIE, 53-611', 'AIPOPLP1XXX', FALSE, 5, 6),
--                                               (7, 'ABLV BANK, AS IN LIQUIDATION', 'MIHAILA TALA STREET 1 RIGA, RIGA, LV-1045', 'AIZKLV22XXX', FALSE, 6, 7),
--                                               (9, 'ALIOR BANK SPOLKA AKCYJNA', 'WARSZAWA, MAZOWIECKIE', 'ALBPPLP1BMW', FALSE, 5, 6),
--                                               (10, 'ALIOR BANK SPOLKA AKCYJNA', 'LOPUSZANSKA BUSINESS PARK LOPUSZANSKA 38 D WARSZAWA, MAZOWIECKIE, 02-232', 'ALBPPLPWXXX', FALSE, 5, 6),
--                                               (11, 'AIB BANK NV', 'WILHELMINASTRAAT 36 - ORANJESTAD ORANJESTAD-WEST AND ORANJESTAD-EAST ,ORANJESTAD', 'ANIBAWA1XXX', FALSE, 8, 9),
--                                               (12, 'ANTEL', 'GUATEMALA 1075 MONTEVIDEO, MONTEVIDEO, 11800', 'ANTLUYM1XXX', FALSE, 3, 4),
--                                               (13, 'APS BANK PLC.', 'APS CENTRE TOWER STREET BIRKIRKARA, BIRKIRKARA, BKR 4012', 'APSBMTMTXXX', FALSE, 7, 8),
--                                               (14, 'ARUBA BANK, LTD', 'CAMACURI 12 - ORANJESTAD ORANJESTAD-WEST AND ORANJESTAD-EAST ,ORANJESTAD', 'ARUBAWAXXXX', FALSE, 8, 9),
--                                               (15, 'JESMOND MIZZI FINANCIAL ADVISORS LIMITED', 'ABATE RIGORD STREET TAXBIEX', 'XBIEX, XBX 1120', 'ATISMTM1XXX', FALSE, 7, 8),
--                                               (16, 'AVAL IN JSC', 'TODOR ALEKSANDROV BLVD 73 FLOOR 1 SOFIA, SOFIA, 1303', 'AVJCBGS1XXX', FALSE, 2, 3),
--                                               (17, 'BANK JULIUS BAER (MONACO) S.A.M.', '12 BOULEVARD DES MOULINS MONACO, MONACO, 98000', 'BAERMCMCXXX', TRUE, 4, 5),
--                                               (18, 'BALKAN ADVISORY COMPANY-IP EAD', 'SQ. POZITANO 9 SOFIA, SOFIA, 1000', 'BAPDBGS1XXX', FALSE, 2, 3),
--                                               (19, 'BARCLAYS BANK S.A', '31 AVENUE DE LA COSTA MONACO, MONACO, 98000', 'BARCMCC1XXX', TRUE, 4, 5),
--                                               (20, 'BARCLAYS BANK PLC MONACO', '31 AVENUE DE LA COSTA MONACO, MONACO, 98000', 'BARCMCMXXXX', FALSE, 4, 5),
--                                               (21, 'BANCO BILBAO VIZCAYA ARGENTARIA URUGUAY S.A.', '25 DE MAYO, ESQUINA ZAQBALA 401 MONTEVIDEO, MONTEVIDEO, 11000', 'BBVAUYMMXXX', FALSE, 3, 4),
--                                               (22, 'BCI CORREDOR DE BOLSA S.A.', 'SANTIAGO', 'BCCSCLR1XXX', FALSE, 9, 10),
--                                               (23, 'BANCO CENTRAL DE CHILE', 'SANTIAGO', 'BCECCLRFXXX', FALSE, 9, 10),
--                                               (24, 'BANCO CENTRAL DE CHILE', 'SANTIAGO', 'BCECCLRMXXX', FALSE, 9, 10),
--                                               (25, 'BNP PARIBAS WEALTH MANAGEMENT MONACO', '27 BOULEVARD PRINCESSE CHARLOTTE MONACO, MONACO, 98000', 'BCGMMCM1XXX', FALSE, 4, 5),
--                                               (26, 'BANCO DE CHILE', 'VINA DEL MAR', 'BCHICLR10R2', FALSE, 9, 10);
