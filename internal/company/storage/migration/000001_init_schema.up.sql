CREATE TABLE company
(
    id      serial
        constraint company_pk
            primary key,
    name    varchar,
    code    varchar,
    country varchar,
    website varchar,
    phone   varchar
);