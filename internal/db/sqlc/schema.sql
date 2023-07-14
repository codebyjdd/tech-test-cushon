CREATE TABLE companies
(
    id int unsigned auto_increment primary key,
    name text,
    enabled tinyint unsigned not null default 1
);

CREATE TABLE isa_funds
(
    id int unsigned auto_increment primary key,
    name text
);

CREATE TABLE companies_isa_funds
(
    id int unsigned auto_increment primary key,
    company_id int unsigned not null,
    fund_id int unsigned not null,
    INDEX(company_id, fund_id)
);

CREATE TABLE customers
(
    id int unsigned auto_increment primary key,
    company_id int unsigned not null,
    name text,
    INDEX(company_id)
);

CREATE TABLE customers_investments
(
    id int unsigned auto_increment primary key,
    customer_id int unsigned not null,
    fund_id int unsigned not null,
    created datetime,
    updated datetime,
    amount decimal(9,2)
);