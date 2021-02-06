package main

import (
	"authenticator/config"
	"authenticator/internal/core"
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

	db := config.NewDatabaseConnection()

	server := infrastructure.NewServer(os.Getenv("APP_PORT"), os.Getenv("GIN_MODE"))

	repositories := core.InitRepository(db.Conn)

	infrastructure.NewRouter(repositories).SetRoutes(server.Router)

	server.Start()
}
