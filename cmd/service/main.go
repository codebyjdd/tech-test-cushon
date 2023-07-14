package main

import (
	"fmt"
	"net/http"
	"os"

	handlerAvailableFunds "github.com/codebyjdd/tech-test-cushon/internal/handlers/available-funds"
	handlerGetInvestment "github.com/codebyjdd/tech-test-cushon/internal/handlers/get-investment"
	handlerSetInvestment "github.com/codebyjdd/tech-test-cushon/internal/handlers/set-investment"
	"github.com/codebyjdd/tech-test-cushon/internal/log"
	"github.com/julienschmidt/httprouter"
)

const (
	Port = 8080
)

func main() {
	logger := log.New()
	getAvailableFundsHandler, err := handlerAvailableFunds.New()
	if err != nil {
		logger.Errorf("could not set up get available funds handler : %s", err.Error())
		os.Exit(0)
	}

	getInvestmentHandler, err := handlerGetInvestment.New()
	if err != nil {
		logger.Errorf("could not set up get investment handler : %s", err.Error())
		os.Exit(0)
	}

	setInvestmentHandler, err := handlerSetInvestment.New()
	if err != nil {
		logger.Errorf("could not set up set investment handler : %s", err.Error())
		os.Exit(0)
	}

	router := httprouter.New()
	router.GET("/funds", getAvailableFundsHandler.Handle)
	router.GET("/investment/customer/:customer/fund/:fund", getInvestmentHandler.Handle)
	router.GET("/invest/customer/:customer/fund/:fund/amount/:amount", setInvestmentHandler.Handle)

	logger.Infof("Listening on port %d", Port)
	if err = http.ListenAndServe(fmt.Sprintf(":%d", Port), router); err != nil {
		logger.Errorf("error serving requests : %s", err.Error())
	}
}
