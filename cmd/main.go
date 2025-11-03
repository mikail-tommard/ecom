package main

import (
	"database/sql"
	"log"

	"github.com/mikail-tommard/ecom/cmd/api"
	"github.com/mikail-tommard/ecom/config"
	"github.com/mikail-tommard/ecom/db"
)



func main() {
	db, err := db.NewPostgreSQLStorage(db.Config{
		Host: config.Envs.DBHost,
		Port: config.Envs.DBPort,
		User: config.Envs.DBUser,
		Password: config.Envs.DBPassword,
		DBName: config.Envs.DBName,
		SSLMode: config.Envs.DBSSLMode,
	})
	if err != nil {
		log.Fatal(err)
	}

	initStorage(db)

	server := api.NewAPIServer(":8080", nil)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}

func initStorage(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("DB: Successfuly connected!")
}