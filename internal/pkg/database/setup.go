package database

import (
	"fmt"
	"github.com/mattmoran/fyp/api/pkg/database/models"
	"github.com/mattmoran/fyp/api/pkg/database/repositories"
	"github.com/mattmoran/fyp/api/pkg/logger"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	UserRepo *repositories.UserRepo
)

func ConnectToDatabase(username, password, dbName, host, port string) error {
	logger.Init()
	defer zap.S().Sync()

	zap.S().Info("Connecting to database")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", username, password, host, port, dbName)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		zap.S().Fatalf("Error connecting to database, %v", err)
		return err
	}

	zap.S().Info("Successfully connected to database")

	zap.S().Info("Migrating tables")
	err = db.AutoMigrate(
		&models.User{},
	)
	if err != nil {
		zap.S().Warn("Error migrating tables", err)
		return err
	}

	UserRepo = repositories.NewUserRepo(db)

	return nil

}
