package apitesting

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/mattmoran/fyp/api/pkg/database"
	"github.com/mattmoran/fyp/api/pkg/database/models"
	"github.com/mattmoran/fyp/api/pkg/utils"
	"os"
	"testing"
)

func FileGroupTest(t *testing.T, r Request) {
	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := "fyp_test"
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	database.ConnectToDatabase(dbUsername, dbPassword, dbName, dbHost, dbPort)

	for owner := 0; owner < 2; owner++ {
		for public := 0; public < 2; public++ {
			for usw := 0; usw < 2; usw++ {
				for usr := 0; usr < 2; usr++ {
					for gsw := 0; gsw < 2; gsw++ {
						for gsr := 0; gsr < 2; gsr++ {
							fmt.Printf(
								"\n\n\nTesting [datasetGroup] - owner: %v public: %v usw: %v usr: %v gsw: %v gsr: %v",
								owner, public, usw, usr, gsw, gsr,
							)
							database.UserRepo.CreateUser(models.User{
								Base: models.Base{
									ID:        "testuserowner",
									CreatedAt: 100,
									UpdatedAt: 100,
								},
								Email:     "test@test.com",
								FirstName: "test",
								LastName:  "testson",
							})
							database.UserRepo.CreateUser(models.User{
								Base: models.Base{
									ID:        "testusercollaborator",
									CreatedAt: 100,
									UpdatedAt: 100,
								},
								Email:     "test@test.com",
								FirstName: "test",
								LastName:  "testson",
							})
							database.UserRepo.CreateUser(models.User{
								Base: models.Base{
									ID:        "testusernotowner",
									CreatedAt: 100,
									UpdatedAt: 100,
								},
								Email:     "test@test.com",
								FirstName: "test",
								LastName:  "testson",
							})

							userID := "testuserowner"
							if owner == 0 {
								userID = "testusernotowner"
							}

							database.DatasetRepo.CreateDataset(models.Dataset{
								Base: models.Base{
									ID:        "testdataset",
									CreatedAt: 100,
									UpdatedAt: 100,
								},
								DatasetName: "test",
								IsPublic:    public != 0,
								OwnerID:     "testuserowner",
							})
							database.StaredDatasetRepo.StarDataset(userID, "testdataset")
							database.DatasetCollaboratorRepo.CreateDatabaseCollaborator(models.DatasetCollaborator{
								Base:      models.Base{ID: "testcollaborator"},
								UserID:    "testusernotowner",
								DatasetID: "testdataset",
							})
							if r.Method != "DELETE" {
								database.FileRepo.CreateFile(models.File{
									Base: models.Base{
										ID:        "testfile1",
										UpdatedAt: 100,
										CreatedAt: 100,
									},
									Name:      "Test file1.tif",
									OwnerID:   "testuserowner",
									Status:    "processed",
									DatasetID: "testdataset",
								})
								database.FileRepo.CreateFile(models.File{
									Base: models.Base{
										ID:        "testfile2",
										UpdatedAt: 100,
										CreatedAt: 100,
									},
									Name:      "Test file2.tif",
									OwnerID:   "testuserowner",
									Status:    "processed",
									DatasetID: "testdataset",
								})
							}

							if usw != 0 {
								database.UserShareDatasetRepo.Create(models.UserShareDataset{
									UserID:      userID,
									DatasetID:   "testdataset",
									WriteAccess: true,
								})
							}

							if usr != 0 {
								database.UserShareDatasetRepo.Create(models.UserShareDataset{
									UserID:      userID,
									DatasetID:   "testdataset",
									WriteAccess: false,
								})
							}

							if gsw != 0 || gsr != 0 {
								database.UserRepo.CreateUser(models.User{
									Base: models.Base{
										ID:        "testusergroupowner",
										CreatedAt: 100,
										UpdatedAt: 100,
									},
									Email:     "test@test.com",
									FirstName: "test",
									LastName:  "testson",
								})
								database.GroupRepo.Create(models.UserGroup{
									Base: models.Base{
										ID:        "testgroup",
										CreatedAt: 100,
										UpdatedAt: 100,
									},
									GroupName: "test",
									OwnerID:   "testusergroupowner",
								})
								database.GroupMemberRepo.AddUserToGroup(userID, "testgroup")
							}
							if gsw != 0 {
								database.GroupShareDatasetRepo.Create(models.GroupShareDataset{
									GroupID:     "testgroup",
									DatasetID:   "testdataset",
									WriteAccess: true,
								})
							}

							if gsr != 0 {
								database.GroupShareDatasetRepo.Create(models.GroupShareDataset{
									GroupID:     "testgroup",
									DatasetID:   "testdataset",
									WriteAccess: false,
								})
							}

							fmt.Printf("jwt: %s\n", userID)
							jwt, _ := utils.CreateJWT(utils.JWTPayload{
								UserId:         userID,
								Email:          "test@test.com",
								FirstName:      "test",
								LastName:       "teston",
								StandardClaims: jwt.StandardClaims{},
							})

							newRequest := r

							if r.ResponseBody != nil {
								_, ok := r.ResponseBody["access"]
								if ok {
									if owner == 1 || usw == 1 || gsw == 1 {
										newRequest.ResponseBody = map[string]interface{}{"access": "write"}
									} else if usr == 1 || gsr == 1 || public == 1 {
										newRequest.ResponseBody = map[string]interface{}{"access": "read"}
									} else {
										newRequest.ResponseBody = map[string]interface{}{"access": "none"}
									}
								}
							}

							if owner == 1 || usw == 1 || gsw == 1 {
								newRequest.ResponseCode = r.ResponseCode
							} else if (usr == 1 || gsr == 1 || public == 1) && r.Method == "GET" {
								newRequest.ResponseCode = r.ResponseCode
							} else if (usr == 1 || gsr == 1 || public == 1) && r.Method != "GET" {
								newRequest.ResponseCode = 403
								newRequest.ResponseBody = map[string]interface{}{"error": "Forbidden"}
							} else {
								newRequest.ResponseCode = 404
								newRequest.ResponseBody = map[string]interface{}{"error": "Dataset not found"}
							}
							AuthGroupTest(t, newRequest, jwt)

							database.DatasetCollaboratorRepo.RemoveCollaborator("testdataset", "testusercollaborator")
							database.DatasetCollaboratorRepo.RemoveCollaborator("testdataset", "testusernotowner")
							database.DatasetRepo.DeleteDatasetByID("testdataset")
							if usw != 0 || usr != 0 {
								database.UserShareDatasetRepo.DeleteUserShareDataset("testdataset", userID)
							}
							if gsr != 0 || gsw != 0 {
								database.GroupShareDatasetRepo.DeleteGroupShareDataset("testdataset", "testgroup")
								database.GroupMemberRepo.RemoveUserFromGroup(userID, "testgroup")
								database.GroupRepo.DeleteGroup("testgroup")
								database.UserRepo.DeleteUserByID("testusergroupowner")
							}
							database.StaredDatasetRepo.UnstarDataset(userID, "testdataset")
							if r.Method != "DELETE" {
								database.FileRepo.DeleteFile("testfile1")
								database.FileRepo.DeleteFile("testfile2")
							}
							database.UserRepo.DeleteUserByID("testuserowner")
							database.UserRepo.DeleteUserByID("testusernotowner")
							database.UserRepo.DeleteUserByID("testusercollaborator")
						}
					}
				}
			}
		}
	}
}
