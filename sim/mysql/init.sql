CREATE DATABASE cushon;

CREATE TABLE cushon.companies
(
    id int unsigned auto_increment primary key,
    name text,
    enabled tinyint unsigned not null default 1
);

CREATE TABLE cushon.isa_funds
(
    id int unsigned auto_increment primary key,
    name text
);

CREATE TABLE cushon.companies_isa_funds
(
    id int unsigned auto_increment primary key,
    company_id int unsigned not null,
    fund_id int unsigned not null,
    INDEX(company_id, fund_id)
);

CREATE TABLE cushon.customers
(
    id int unsigned auto_increment primary key,
    company_id int unsigned not null,
    name text,
    INDEX(company_id)
);

CREATE TABLE cushon.customers_investments
(
    id int unsigned auto_increment primary key,
    customer_id int unsigned not null,
    fund_id int unsigned not null,
    created datetime,
    updated datetime,
    amount decimal(9,2),
    INDEX(customer_id, fund_id)
);

INSERT INTO cushon.companies
    (name)
VALUES
    ('McDriscoll'), ('BigCompany'), ('GeneralPublic');

INSERT INTO cushon.isa_funds
    (name)
VALUES
    ('Cushon Equities Fund'), ('No risk, no reward'), ('experimental technologies 3');

INSERT INTO cushon.companies_isa_funds
    (company_id, fund_id)
VALUES
    (1, 1), (1, 2), (1, 3), (2, 1), (2, 2), (3, 1);

INSERT INTO cushon.customers
    (company_id, name)
VALUES
    (3, 'John Driscoll');