package handlerGetInvestment

const (
	ErrResp     = "error writing http response"
	ErrNoCust   = "investment request with no supplied customer id"
	ErrNoFund   = "investment request with no supplied fund id"
	ErrDb       = "error getting customer fund investment from db"
	ErrRespCust = "you must supply a customer Id with your request"
	ErrRespFund = "you must supply the id of the fund you wish to invest in"
	ErrRespDb   = "investment not found"
)
