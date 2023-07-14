package handlerGetInvestment

import (
	"github.com/codebyjdd/tech-test-cushon/internal/db"
	"github.com/codebyjdd/tech-test-cushon/internal/log"
)

func New() (Handler, error) {
	conn, err := db.New()
	if err != nil {
		return nil, err
	}
	return handler{
		db:  conn,
		log: log.New(),
	}, nil
}
