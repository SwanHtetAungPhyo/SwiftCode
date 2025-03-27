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


