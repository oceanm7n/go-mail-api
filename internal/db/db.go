package db

// Package db provides database related functions for the application. Database is postgresql.

import (
	"database/sql"
	"fmt"
	"log"
	"mailganer-task/internal/config"

	_ "github.com/lib/pq"
)

// DB is a database connection
var DB *sql.DB

// Connect connects to the database
func Connect() *sql.DB {
	var err error
	config := config.GetConfig()
	DB, err = sql.Open("postgres", "host="+config.DB.Host+" port="+fmt.Sprint(config.DB.Port)+" user="+config.DB.User+" password="+config.DB.Password+" sslmode=disable")

	if err != nil {
		log.Fatal(err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatal(err)
	}

	return DB
}
