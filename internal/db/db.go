package db

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"time"

	sqlc "github.com/codebyjdd/tech-test-cushon/internal/db/sqlc/generated"
	"github.com/codebyjdd/tech-test-cushon/internal/entities"
)

type db struct {
	db *sqlc.Queries
}

func (d db) GetAvailablePublicFunds() ([]entities.Fund, error) {
	res, err := d.db.GetPublicFunds(context.Background())
	if err != nil {
		return nil, err
	}

	results := make([]entities.Fund, 0)
	for _, result := range res {
		results = append(results, entities.Fund{
			Id:   int(result.ID),
			Name: result.Name.String,
		})
	}
	return results, nil
}

func (d db) GetCustomer(customerId int) (entities.Customer, error) {
	res, err := d.db.GetCustomer(context.Background(), uint32(customerId))
	if err != nil {
		return entities.Customer{}, err
	}
	return entities.Customer{
		Id:   int(res.ID),
		Name: res.Name.String,
		Company: entities.Company{
			Id:   int(res.CompanyID),
			Name: res.Company.String,
		},
	}, nil
}

func (d db) GetCustomerInvestment(customer entities.Customer, fund entities.Fund) (float64, error) {
	investment, err := d.db.GetCustomerFundInvestment(context.Background(), sqlc.GetCustomerFundInvestmentParams{
		CustomerID: uint32(customer.Id),
		FundID:     uint32(fund.Id),
	})
	if err != nil {
		return 0.0, err
	}
	amt, err := strconv.ParseFloat(investment.String, 64)
	if err != nil {
		return 0.0, err
	}
	return amt, nil
}

func (d db) GetFund(fundId int) (entities.Fund, error) {
	res, err := d.db.GetFund(context.Background(), uint32(fundId))
	if err != nil {
		return entities.Fund{}, err
	}
	return entities.Fund{
		Id:   int(res.ID),
		Name: res.Name.String,
	}, nil
}

func (d db) SetCustomerInvestment(customer entities.Customer, fund entities.Fund, amount float64) error {
	_, err := d.db.GetCustomerFundInvestment(context.Background(), sqlc.GetCustomerFundInvestmentParams{
		CustomerID: uint32(customer.Id),
		FundID:     uint32(fund.Id),
	})

	if err != nil {
		fmt.Println("fund fetch error", err.Error())
		return d.db.CreateCustomerFundInvestment(context.Background(), sqlc.CreateCustomerFundInvestmentParams{
			CustomerID: int32(customer.Id),
			FundID:     int32(fund.Id),
			Created:    sql.NullTime{Valid: true, Time: time.Now()},
			Updated:    sql.NullTime{Valid: true, Time: time.Now()},
			Amount:     sql.NullString{Valid: true, String: fmt.Sprintf("%f", amount)},
		})
	}
	return d.db.UpdateCustomerFundInvestment(context.Background(), sqlc.UpdateCustomerFundInvestmentParams{
		Amount:     sql.NullString{Valid: true, String: fmt.Sprintf("%f", amount)},
		Updated:    sql.NullTime{Valid: true, Time: time.Now()},
		CustomerID: uint32(customer.Id),
		FundID:     uint32(fund.Id),
	})
}
