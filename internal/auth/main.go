package main

import (
	"github.com/mattmoran/fyp/api/auth/web"
	"github.com/mattmoran/fyp/api/pkg/database"
	"github.com/mattmoran/fyp/api/pkg/logger"
	"go.uber.org/zap"
	"os"
)

func main() {
	logger.Init()
	defer zap.S().Sync()

	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")

	zap.S().Infof("Connecting to database: %s %s:%s", dbName, dbHost, dbPort)

	err := database.ConnectToDatabase(dbUsername, dbPassword, dbName, dbHost, dbPort)
	if err != nil {
		zap.S().Error("Failed to connect to database", err)
	}

	web.InitiateServer()
}
