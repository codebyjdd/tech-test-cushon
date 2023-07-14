package db

import (
	"database/sql"

	sqlc "github.com/codebyjdd/tech-test-cushon/internal/db/sqlc/generated"
	_ "github.com/go-sql-driver/mysql"
)

func New() (Db, error) {
	conn, err := sql.Open("mysql", "root:rootpass@/cushon")
	if err != nil {
		return nil, err
	}

	return db{
		db: sqlc.New(conn),
	}, nil
}
