package db

import (
	"github.com/codebyjdd/tech-test-cushon/internal/entities"
)

type Db interface {
	GetAvailablePublicFunds() ([]entities.Fund, error)
	GetCustomer(customerId int) (entities.Customer, error)
	GetCustomerInvestment(customer entities.Customer, fund entities.Fund) (float64, error)
	GetFund(fundId int) (entities.Fund, error)
	SetCustomerInvestment(customer entities.Customer, fund entities.Fund, amount float64) error
}
