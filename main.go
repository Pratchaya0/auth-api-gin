package main

import (
	"os"

	"github.com/Pratchaya0/auth-api-gin/configs"
	"github.com/Pratchaya0/auth-api-gin/modules/servers"
	"github.com/Pratchaya0/auth-api-gin/pkg/databases"
	"github.com/joho/godotenv"
)

func main() {
	// Load dotenv config
	if err := godotenv.Load(".env"); err != nil {
		panic(err.Error())
	}
	cfg := new(configs.Configs)

	// Fiber configs
	cfg.App.Host = os.Getenv("GIN_HOST")
	cfg.App.Port = os.Getenv("GIN_PORT")

	// Database Configs
	cfg.PostgreSQL.Host = os.Getenv("DB_HOST")
	cfg.PostgreSQL.Port = os.Getenv("DB_PORT")
	cfg.PostgreSQL.Protocol = os.Getenv("DB_PROTOCOL")
	cfg.PostgreSQL.Username = os.Getenv("DB_USERNAME")
	cfg.PostgreSQL.Password = os.Getenv("DB_PASSWORD")
	cfg.PostgreSQL.Database = os.Getenv("DB_DATABASE")

	databases.SetupDatabase()

	s := servers.NewServer(cfg, databases.DB())
	s.Start()
}
