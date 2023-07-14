// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.1

package sqlc

import (
	"database/sql"
)

type CompaniesIsaFund struct {
	ID        uint32
	CompanyID uint32
	FundID    uint32
}

type Company struct {
	ID      uint32
	Name    sql.NullString
	Enabled uint32
}

type Customer struct {
	ID        uint32
	CompanyID uint32
	Name      sql.NullString
}

type CustomersInvestment struct {
	ID         uint32
	CustomerID uint32
	FundID     uint32
	Created    sql.NullTime
	Updated    sql.NullTime
	Amount     sql.NullString
}

type IsaFund struct {
	ID   uint32
	Name sql.NullString
}