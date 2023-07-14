package handlerSetInvestment

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/codebyjdd/tech-test-cushon/internal/db"
	"github.com/codebyjdd/tech-test-cushon/internal/entities"
	"github.com/julienschmidt/httprouter"
	"github.com/sirupsen/logrus"
)

type Handler interface {
	Handle(w http.ResponseWriter, r *http.Request, p httprouter.Params)
}

type handler struct {
	db  db.Db
	log *logrus.Logger
}

func (h handler) Handle(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	customer, fund, amount, err := h.validateParams(p)
	if err != nil {
		var ourMsg, theirMsg string

		switch err.Error() {
		case "customer":
			ourMsg = ErrNoCust
			theirMsg = ErrRespCust
		case "fund":
			ourMsg = ErrNoFund
			theirMsg = ErrRespFund
		case "amount":
			ourMsg = ErrNoAmt
			theirMsg = ErrRespAmt
		}

		h.log.Error(ourMsg)
		w.WriteHeader(http.StatusBadRequest)
		if _, err := w.Write([]byte(theirMsg)); err != nil {
			h.log.Errorf("%s : %s", ErrResp, err.Error())
		}
		return
	}
	if err = h.db.SetCustomerInvestment(customer, fund, amount); err != nil {
		h.log.Errorf("%s : %s", ErrDb, err.Error())
		w.WriteHeader(http.StatusBadRequest)
		if _, err = w.Write([]byte(ErrRespDb)); err != nil {
			h.log.Errorf("%s : %s", ErrResp, err.Error())
		}
		return
	}
	h.log.Infof("%s : %s : %s : %d", InfoSetInvest, customer.Name, fund.Name, amount)
}

func (h handler) validateParams(p httprouter.Params) (entities.Customer, entities.Fund, float64, error) {
	customer := entities.Customer{}
	fund := entities.Fund{}
	amount := 0.0

	if len(p.ByName("customer")) < 1 {
		return customer, fund, amount, errors.New("customer")
	}

	custId, err := strconv.Atoi(p.ByName("customer"))
	if err != nil {
		return customer, fund, amount, errors.New("customer")
	}

	customer, err = h.db.GetCustomer(custId)
	if err != nil {
		return customer, fund, amount, errors.New("customer")
	}

	if len(p.ByName("fund")) < 1 {
		return customer, fund, amount, errors.New("fund")
	}

	fundId, err := strconv.Atoi(p.ByName("fund"))
	if err != nil {
		return entities.Customer{}, fund, amount, errors.New("fund")
	}

	fund, err = h.db.GetFund(fundId)
	if err != nil {
		return entities.Customer{}, fund, amount, errors.New("fund")
	}

	amount, err = strconv.ParseFloat(p.ByName("amount"), 64)
	if err != nil {
		return entities.Customer{}, entities.Fund{}, 0.0, errors.New("amount")
	}
	return customer, fund, amount, nil
}
