package handlerGetInvestment

import (
	"errors"
	"fmt"
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
	customer, fund, err := h.validateParams(p)
	if err != nil {
		var ourMsg, theirMsg string

		switch err.Error() {
		case "customer":
			ourMsg = ErrNoCust
			theirMsg = ErrRespCust
		case "fund":
			ourMsg = ErrNoFund
			theirMsg = ErrRespFund
		}

		h.log.Error(ourMsg)
		w.WriteHeader(http.StatusBadRequest)
		if _, err = w.Write([]byte(theirMsg)); err != nil {
			h.log.Errorf("%s : %s", ErrResp, err.Error())
		}
		return
	}

	investment, err := h.db.GetCustomerInvestment(customer, fund)
	if err != nil {
		h.log.Error(ErrDb)
		w.WriteHeader(http.StatusBadRequest)
		if _, err = w.Write([]byte(ErrRespDb)); err != nil {
			h.log.Errorf("%s : %s", ErrResp, err.Error())
		}
		return
	}
	w.Write([]byte(fmt.Sprintf("%f", investment)))
}

func (h handler) validateParams(p httprouter.Params) (entities.Customer, entities.Fund, error) {
	customer := entities.Customer{}
	fund := entities.Fund{}

	if len(p.ByName("customer")) < 1 {
		return customer, fund, errors.New("customer")
	}

	custId, err := strconv.Atoi(p.ByName("customer"))
	if err != nil {
		return customer, fund, errors.New("customer")
	}

	customer, err = h.db.GetCustomer(custId)
	if err != nil {
		return customer, fund, errors.New("customer")
	}

	if len(p.ByName("fund")) < 1 {
		return customer, fund, errors.New("fund")
	}

	fundId, err := strconv.Atoi(p.ByName("fund"))
	if err != nil {
		return entities.Customer{}, fund, errors.New("fund")
	}

	fund, err = h.db.GetFund(fundId)
	if err != nil {
		return entities.Customer{}, fund, errors.New("fund")
	}
	return customer, fund, nil
}
