package web

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/mattmoran/fyp/api/pkg/apitesting"
	"github.com/mattmoran/fyp/api/pkg/database"
	"github.com/mattmoran/fyp/api/pkg/database/models"
	"github.com/mattmoran/fyp/api/pkg/utils"
	"os"
	"testing"
)

// Body types
type SimpleSearchBody struct {
	Query string `json:"query"`
}

func TestInitiateServer(t *testing.T) {
	token, _ := utils.CreateJWT(utils.JWTPayload{
		UserId:         "test1",
		Email:          "",
		FirstName:      "",
		LastName:       "",
		StandardClaims: jwt.StandardClaims{},
	})

	token2, _ := utils.CreateJWT(utils.JWTPayload{
		UserId:         "test2",
		Email:          "",
		FirstName:      "",
		LastName:       "",
		StandardClaims: jwt.StandardClaims{},
	})

	tests := []apitesting.Test{
		{
			Name:         "Simple search - no auth provided",
			Method:       "POST",
			Path:         "/simpleSearch",
			Auth:         "",
			Body:         map[string]interface{}{"query": "test"},
			ResponseCode: 401,
			ResponseBody: map[string]interface{}{"error": "JWT token is missing or in an invalid format"},
		},
		{
			Name:         "Simple search - no body",
			Method:       "POST",
			Path:         "/simpleSearch",
			Auth:         token,
			Body:         nil,
			ResponseCode: 400,
			ResponseBody: map[string]interface{}{"error": "No search query found"},
		},
		{
			Name:         "Simple search - Return all user's files and datasets",
			Method:       "POST",
			Path:         "/simpleSearch",
			Auth:         token,
			Body:         map[string]interface{}{"query": "test"},
			ResponseCode: 200,
			ResponseBody: map[string]interface{}{"files": []any{
				map[string]any{"id": "testfile1", "created_at": float64(100), "updated_at": float64(100), "name": "testfile1", "owner_id": "test1", "status": "processed", "dataset_id": "test1", "is_public": false},
				map[string]any{"id": "testfile2", "created_at": float64(100), "updated_at": float64(100), "name": "testfile2", "owner_id": "test1", "status": "processed", "dataset_id": "test1", "is_public": true},
			}, "datasets": []any{
				map[string]any{"id": "test1", "created_at": float64(100), "updated_at": float64(100), "dataset_name": "test dataset", "owner_id": "test1", "is_public": false},
			},
			},
		},
		{
			Name:         "Simple search - Only return public or user's files",
			Method:       "POST",
			Path:         "/simpleSearch",
			Auth:         token2,
			Body:         map[string]interface{}{"query": "test"},
			ResponseCode: 200,
			ResponseBody: map[string]interface{}{"files": []any{
				map[string]any{"id": "testfile2", "created_at": float64(100), "updated_at": float64(100), "name": "testfile2", "owner_id": "test1", "status": "processed", "dataset_id": "test1", "is_public": true},
			}, "datasets": []any{}},
		},
	}

	gin.SetMode(gin.TestMode)

	// Connect to the database
	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := "fyp_test"
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	database.ConnectToDatabase(dbUsername, dbPassword, dbName, dbHost, dbPort)

	// Initiate the database
	database.UserRepo.CreateUser(models.User{
		Base: models.Base{
			ID: "test1",
		},
		Email:     "test1@test.com",
		FirstName: "test1",
		LastName:  "test",
	})
	database.UserRepo.CreateUser(models.User{
		Base: models.Base{
			ID: "test2",
		},
		Email:     "test2@test.com",
		FirstName: "test2",
		LastName:  "test",
	})

	database.DatasetRepo.CreateDataset(models.Dataset{
		Base: models.Base{
			ID:        "test1",
			CreatedAt: 100,
			UpdatedAt: 100,
		},
		DatasetName: "test dataset",
		OwnerID:     "test1",
		IsPublic:    false,
	})

	database.FileRepo.CreateFile(models.File{
		Base: models.Base{
			ID:        "testfile1",
			CreatedAt: 100,
			UpdatedAt: 100,
		},
		Name:      "testfile1",
		OwnerId:   "test1",
		Status:    "processed",
		DatasetID: "test1",
		IsPublic:  false,
	})
	database.FileRepo.CreateFile(models.File{
		Base: models.Base{
			ID:        "testfile2",
			CreatedAt: 100,
			UpdatedAt: 100,
		},
		Name:      "testfile2",
		OwnerId:   "test1",
		Status:    "processed",
		DatasetID: "test1",
		IsPublic:  true,
	})

	router := InitiateServer()

	apitesting.TestServer(t, tests, router)
}
