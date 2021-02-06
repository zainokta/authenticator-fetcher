package main

import (
	"authenticator/internal/infrastructure"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	server := infrastructure.NewServer(os.Getenv("APP_PORT"), os.Getenv("GIN_MODE"))
	server.Start()
}
