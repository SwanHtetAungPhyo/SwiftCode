create table countries
(
    id              bigserial
        primary key,
    countryiso2code char(2)      not null
        constraint uni_countries_countryiso2code
            unique,
    name            varchar(255) not null
        constraint uni_countries_name
            unique,
    timezone        varchar(255) not null
);


create table bank_details
(
    id             bigserial
        primary key,
    name           varchar(255) not null,
    address        varchar(255) not null,
    town_name      varchar(255) not null,
    swift_code     varchar(12)  not null,
    is_headquarter boolean      not null,
    countryid      bigint       not null
        constraint fk_countries_bank_details
            references countries
            on update cascade on delete cascade
);


