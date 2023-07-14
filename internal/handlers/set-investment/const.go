package handlerSetInvestment

const (
	ErrResp       = "error writing http response"
	ErrNoCust     = "investment request with no supplied customer id"
	ErrNoFund     = "investment request with no supplied fund id"
	ErrNoAmt      = "investment request with no supplied investment amount"
	ErrDb         = "error storing investment in the database"
	ErrRespCust   = "you must supply a customer Id with your request"
	ErrRespDb     = "downstream storage error"
	ErrRespFund   = "you must supply the id of the fund you wish to invest in"
	ErrRespAmt    = "you must supply a valid investment amount in GBP ranging from 0.01 to 9,999,999.00"
	InfoSetInvest = "set fund investment for customer"
)
