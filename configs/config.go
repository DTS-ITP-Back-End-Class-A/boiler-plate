package configs

import (
	"log"
	"os"
	"testing/models"

	"github.com/alexsasharegan/dotenv"
)

func SetENV() models.Config {
	err := dotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	mysqlHost := os.Getenv("MYSQL_HOST")
	mysqlUsername := os.Getenv("MYSQL_USERNAME")
	mysqlPassword := os.Getenv("MYSQL_PASSWORD")
	mysqlPort := os.Getenv("MYSQL_PORT")
	mysqlDatabase := os.Getenv("MYSQL_DATABASE")
	mysqlPool := os.Getenv("MYSQL_POOL")
	mysqlTimeOut := os.Getenv("MYSQL_TIMEOUT")

	return models.Config{
		Host:     mysqlHost,
		Username: mysqlUsername,
		Password: mysqlPassword,
		Port:     mysqlPort,
		Database: mysqlDatabase,
		Pool:     mysqlPool,
		Timeout:  mysqlTimeOut,
	}
}
