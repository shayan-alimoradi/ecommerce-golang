package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func NewPostgreStorage(addr string, maxOpenConns, maxIdleConns string, maxIdleTime string) (*sql.DB, error) {
	db, err := sql.Open("postgres", addr)
	if err != nil {
		return nil, err
	}

	return db, nil
}
