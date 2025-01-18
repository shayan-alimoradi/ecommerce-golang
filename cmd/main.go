package main

import (
	"database/sql"
	"log"

	"github.com/shayan-alimoradi/ecommerce-golang/cmd/api"
	"github.com/shayan-alimoradi/ecommerce-golang/config"
	"github.com/shayan-alimoradi/ecommerce-golang/db"
)

func main() {
	db, err := db.NewPostgreStorage(
		config.Envs.ADDR,
		config.Envs.MaxOpenConns,
		config.Envs.MaxIdleConns,
		config.Envs.MaxIdleTime,
	)
	if err != nil {
		log.Fatal(err)
	}

	initStorage(db)

	server := api.NewAPIServer(":8081", db)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}

func initStorage(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("DB successfully connected")
}
