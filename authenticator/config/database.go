package config

import (
	"database/sql"
	"fmt"
	"os"
	"time"
)

type Database struct {
	Conn *sql.DB
}

func NewDatabaseConnection() Database {
	var err error
	var db *sql.DB

	psqlInfo := getPostgreInfo()

	for i := 0; i < 10; i++ {
		db, err = sql.Open("postgres", psqlInfo)

		if err != nil {
			fmt.Printf("Unable to Open DB: %s... Retrying\n", err.Error())
			time.Sleep(time.Second * 6)
		} else if err = db.Ping(); err != nil {
			fmt.Printf("Unable to Ping DB: %s... Retrying\n", err.Error())
			time.Sleep(time.Second * 6)
		} else {
			err = nil
			break
		}
	}

	fmt.Println("Database connected")

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
