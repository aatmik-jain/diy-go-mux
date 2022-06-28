package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

func ConnectDB() *sql.DB {

	USER := os.Getenv("APP_DB_USERNAME")
	PASSWORD := os.Getenv("APP_DB_PASSWORD")
	DBNAME := os.Getenv("APP_DB_NAME")

	connectionString := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", USER, PASSWORD, DBNAME)

	sqlDB, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}
	return sqlDB
}
