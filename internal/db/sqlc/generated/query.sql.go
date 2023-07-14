// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.1
// source: query.sql

package sqlc

import (
	"context"
	"database/sql"
)

const createCustomerFundInvestment = `-- name: CreateCustomerFundInvestment :exec
INSERT INTO customers_investments
    (customer_id, fund_id, created, updated, amount)
VALUES
    (?, ?, ?, ?, ?)
`

type CreateCustomerFundInvestmentParams struct {
	CustomerID int32
	FundID     int32
	Created    sql.NullTime
	Updated    sql.NullTime
	Amount     sql.NullString
}

func (q *Queries) CreateCustomerFundInvestment(ctx context.Context, arg CreateCustomerFundInvestmentParams) error {
	_, err := q.db.ExecContext(ctx, createCustomerFundInvestment,
		arg.CustomerID,
		arg.FundID,
		arg.Created,
		arg.Updated,
		arg.Amount,
	)
	return err
}

const getCustomer = `-- name: GetCustomer :one
SELECT
    c.id,
    c.name,
    c.company_id,
    c2.name as 'company'
FROM customers c

INNER JOIN companies c2
ON (c.company_id = c2.id AND c2.enabled = 1)

WHERE
    c.id = ?
`

type GetCustomerRow struct {
	ID        uint32
	Name      sql.NullString
	CompanyID uint32
	Company   sql.NullString
}

func (q *Queries) GetCustomer(ctx context.Context, id uint32) (GetCustomerRow, error) {
	row := q.db.QueryRowContext(ctx, getCustomer, id)
	var i GetCustomerRow
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.CompanyID,
		&i.Company,
	)
	return i, err
}

const getCustomerFundInvestment = `-- name: GetCustomerFundInvestment :one
SELECT
    ci.amount

FROM customers_investments ci
WHERE
    customer_id = ?
AND fund_id = ?
`

type GetCustomerFundInvestmentParams struct {
	CustomerID uint32
	FundID     uint32
}

func (q *Queries) GetCustomerFundInvestment(ctx context.Context, arg GetCustomerFundInvestmentParams) (sql.NullString, error) {
	row := q.db.QueryRowContext(ctx, getCustomerFundInvestment, arg.CustomerID, arg.FundID)
	var amount sql.NullString
	err := row.Scan(&amount)
	return amount, err
}

const getFund = `-- name: GetFund :one
SELECT id, name
FROM isa_funds
WHERE
    id = ?
`

func (q *Queries) GetFund(ctx context.Context, id uint32) (IsaFund, error) {
	row := q.db.QueryRowContext(ctx, getFund, id)
	var i IsaFund
	err := row.Scan(&i.ID, &i.Name)
	return i, err
}

const getPublicFunds = `-- name: GetPublicFunds :many
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
AND c.enabled = 1
`

func (q *Queries) GetPublicFunds(ctx context.Context) ([]IsaFund, error) {
	rows, err := q.db.QueryContext(ctx, getPublicFunds)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []IsaFund
	for rows.Next() {
		var i IsaFund
		if err := rows.Scan(&i.ID, &i.Name); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateCustomerFundInvestment = `-- name: UpdateCustomerFundInvestment :exec
UPDATE customers_investments
SET
    amount = ?,
    updated = ?
WHERE
    customer_id = ?
AND fund_id = ?
`

type UpdateCustomerFundInvestmentParams struct {
	Amount     sql.NullString
	Updated    sql.NullTime
	CustomerID uint32
	FundID     uint32
}

func (q *Queries) UpdateCustomerFundInvestment(ctx context.Context, arg UpdateCustomerFundInvestmentParams) error {
	_, err := q.db.ExecContext(ctx, updateCustomerFundInvestment,
		arg.Amount,
		arg.Updated,
		arg.CustomerID,
		arg.FundID,
	)
	return err
}
