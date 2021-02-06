package config

import (
	"database/sql"
	"fmt"
	"os"
)

type Database struct {
	Conn *sql.DB
}

func NewDatabaseConnection() Database {
	psqlInfo := getPostgreInfo()

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	return Database{Conn: db}
}

func getPostgreInfo() string {
	return fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_DATABASE"))
}
