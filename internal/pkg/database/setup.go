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
	UserRepo                *repositories.UserRepo
	FileRepo                *repositories.FileRepository
	DatasetRepo             *repositories.DatasetRepository
	DatasetAttributeRepo    *repositories.DatasetAttributeRepository
	FileAttributeRepo       *repositories.FileAttributeRepository
	StaredDatasetRepo       *repositories.StaredDatasetRepository
	DatasetCollaboratorRepo *repositories.DatasetCollaboratorRepository
	GroupRepo               *repositories.GroupRepository
	GroupMemberRepo         *repositories.GroupMemberRepository
	UserShareDatasetRepo    *repositories.UserShareDatasetRepository
	GroupShareDatasetRepo   *repositories.GroupShareDatasetRepository
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
		&models.File{},
		&models.Dataset{},
		&models.DatasetAttribute{},
		&models.FileAttribute{},
		&models.StaredDataset{},
		&models.DatasetCollaborator{},
		&models.UserGroup{},
		&models.GroupMember{},
		&models.UserShareDataset{},
		&models.GroupShareDataset{},
	)
	if err != nil {
		zap.S().Warn("Error migrating tables ", err)
		return err
	}

	UserRepo = repositories.NewUserRepo(db)
	FileRepo = repositories.NewFileRepository(db)
	DatasetRepo = repositories.NewDatasetRepository(db)
	DatasetAttributeRepo = repositories.NewDatasetAttributeRepository(db)
	FileAttributeRepo = repositories.NewFileAttributeRepository(db)
	StaredDatasetRepo = repositories.NewStaredDatasetRepository(db)
	DatasetCollaboratorRepo = repositories.NewDatasetCollaboratorRepository(db)
	GroupRepo = repositories.NewGroupRepository(db)
	GroupMemberRepo = repositories.NewGroupMemberRepository(db)
	UserShareDatasetRepo = repositories.NewUserShareDatasetRepository(db)
	GroupShareDatasetRepo = repositories.NewGroupShareDatasetRepository(db)

	return nil

}
