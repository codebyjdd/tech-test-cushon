-- name: GetCustomer :one
SELECT
    c.id,
    c.name,
    c.company_id,
    c2.name as 'company'
FROM customers c

INNER JOIN companies c2
ON (c.company_id = c2.id AND c2.enabled = 1)

WHERE
    c.id = ?;

-- name: GetFund :one
SELECT *
FROM isa_funds
WHERE
    id = ?;

-- name: GetPublicFunds :many
SELECT
    ifnd.id,
    ifnd.name

FROM isa_funds ifnd

INNER JOIN companies_isa_funds cif
ON (cif.fund_id = ifnd.id)

INNER JOIN companies c
ON (c.id = cif.company_id)

WHERE
    c.name = 'GeneralPublic'
AND c.enabled = 1;

-- name: GetCustomerFundInvestment :one
SELECT
    ci.amount

FROM customers_investments ci
WHERE
    customer_id = ?
AND fund_id = ?;

-- name: UpdateCustomerFundInvestment :exec
UPDATE customers_investments
SET
    amount = ?,
    updated = ?
WHERE
    customer_id = ?
AND fund_id = ?;

-- name: CreateCustomerFundInvestment :exec
INSERT INTO customers_investments
    (customer_id, fund_id, created, updated, amount)
VALUES
    (?, ?, ?, ?, ?)