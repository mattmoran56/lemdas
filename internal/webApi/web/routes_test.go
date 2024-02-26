package web

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/mattmoran/fyp/api/pkg/apitesting"
	"github.com/mattmoran/fyp/api/pkg/database"
	"github.com/mattmoran/fyp/api/pkg/database/models"
	"github.com/mattmoran/fyp/api/pkg/utils"
	"os"
	"testing"
)

func TestInitiateServer(t *testing.T) {
	token1, _ := utils.CreateJWT(utils.JWTPayload{
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

	// All tests should be written for the end point ommitting the slash at the end of the path
	// All test should be written for the end point with auth provided
	handleGetDatasetsTests := []apitesting.Test{
		{
			Name:         "Get datasets - get all user's datasets",
			Method:       "GET",
			Path:         "/dataset",
			Auth:         token1,
			Body:         nil,
			ResponseCode: 200,
			ResponseBody: map[string]interface{}{
				"datasets": []any{
					map[string]any{"id": "test1", "created_at": float64(100), "updated_at": float64(100), "dataset_name": "test dataset", "owner_id": "test1", "is_public": false},
					map[string]any{"id": "test2", "created_at": float64(100), "updated_at": float64(100), "dataset_name": "test dataset", "owner_id": "test1", "is_public": false},
					map[string]any{"id": "test3", "created_at": float64(100), "updated_at": float64(100), "dataset_name": "test dataset", "owner_id": "test1", "is_public": true},
					map[string]any{"id": "test4", "created_at": float64(100), "updated_at": float64(100), "dataset_name": "test dataset", "owner_id": "test1", "is_public": false},
				},
			},
		},
		{
			Name:         "Get datasets - get no other user's datasets",
			Method:       "GET",
			Path:         "/dataset",
			Auth:         token2,
			Body:         nil,
			ResponseCode: 200,
			ResponseBody: map[string]interface{}{
				"datasets": []any{},
			},
		},
		// TODO: test order by
	}
	handleGetDatasetTests := []apitesting.Test{
		{
			Name:         "Get dataset - get user's dataset",
			Method:       "GET",
			Path:         "/dataset/test1",
			Auth:         token1,
			Body:         nil,
			ResponseCode: 200,
			ResponseBody: map[string]interface{}{"owner_name": "test1 test", "id": "test1", "created_at": float64(100), "updated_at": float64(100), "dataset_name": "test dataset", "owner_id": "test1", "is_public": false},
		},
		{
			Name:         "Get dataset - get other user's dataset",
			Method:       "GET",
			Path:         "/dataset/test1",
			Auth:         token2,
			Body:         nil,
			ResponseCode: 404,
			ResponseBody: map[string]interface{}{"error": "Dataset not found"},
		},
	}
	handleGetFilesForDatasetTests := []apitesting.Test{
		{
			Name:         "Get files for dataset - get all user's files",
			Method:       "GET",
			Path:         "/dataset/test1/files",
			Auth:         token1,
			Body:         nil,
			ResponseCode: 200,
			ResponseBody: map[string]interface{}{"files": []any{
				map[string]any{"id": "testfile1", "created_at": float64(100), "updated_at": float64(100), "name": "testfile1", "owner_id": "test1", "status": "processed", "dataset_id": "test1", "is_public": false},
				map[string]any{"id": "testfile2", "created_at": float64(100), "updated_at": float64(100), "name": "testfile2", "owner_id": "test1", "status": "processed", "dataset_id": "test1", "is_public": true},
			},
			},
		},
		{
			Name:         "Get files for dataset - don't get other user's files (only public or shared files)",
			Method:       "GET",
			Path:         "/dataset/test3/files",
			Auth:         token2,
			Body:         nil,
			ResponseCode: 200,
			ResponseBody: map[string]interface{}{"files": []any{
				map[string]any{"id": "testfile4", "created_at": float64(100), "updated_at": float64(100), "name": "testfile2", "owner_id": "test1", "status": "processed", "dataset_id": "test3", "is_public": true},
				map[string]any{"id": "testfile5", "created_at": float64(100), "updated_at": float64(100), "name": "testfile5", "owner_id": "test1", "status": "processed", "dataset_id": "test3", "is_public": true},
			},
			},
		},
		{
			Name:         "Get files for dataset -invalid dataset id",
			Method:       "GET",
			Path:         "/dataset/invalid/files",
			Auth:         token2,
			Body:         nil,
			ResponseCode: 404,
			ResponseBody: map[string]interface{}{"error": "Dataset not found"},
		},
	}
	handleCreateDatasetTests := []apitesting.Test{
		{
			Name:              "Create dataset - create dataset",
			Method:            "POST",
			Path:              "/dataset",
			Auth:              token1,
			Body:              map[string]interface{}{"dataset_name": "test dataset 2", "is_public": false},
			ResponseCode:      201,
			ResponseBody:      map[string]interface{}{"id": "test dataset 2", "created_at": float64(100), "updated_at": float64(100), "dataset_name": "test dataset 2", "owner_id": "test1", "is_public": false},
			ManualCompareBody: true,
		},
		{
			Name:              "Create dataset - create dataset with no is public",
			Method:            "POST",
			Path:              "/dataset",
			Auth:              token1,
			Body:              map[string]interface{}{"dataset_name": "test dataset 3"},
			ResponseCode:      201,
			ResponseBody:      map[string]interface{}{"id": "test dataset 3", "created_at": float64(100), "updated_at": float64(100), "dataset_name": "test dataset 3", "owner_id": "test1", "is_public": false},
			ManualCompareBody: true,
		},
		{
			Name:         "Create dataset - create dataset with is public and no name",
			Method:       "POST",
			Path:         "/dataset",
			Auth:         token1,
			Body:         map[string]interface{}{"is_public": true},
			ResponseCode: 400,
			ResponseBody: map[string]interface{}{"error": "Invalid request body"},
		},
		{
			Name:         "Create dataset - create dataset with no body",
			Method:       "POST",
			Path:         "/dataset",
			Auth:         token1,
			Body:         nil,
			ResponseCode: 400,
			ResponseBody: map[string]interface{}{"error": "Invalid request body"},
		},
	}
	handleDeleteDatasetTests := []apitesting.Test{
		{
			Name:         "Delete dataset - delete dataset that is public but not owned by user",
			Method:       "DELETE",
			Path:         "/dataset/test3",
			Auth:         token2,
			Body:         nil,
			ResponseCode: 403,
			ResponseBody: map[string]interface{}{"error": "Forbidden"},
		},
		{
			Name:         "Delete dataset - delete dataset that is private but not owned by user",
			Method:       "DELETE",
			Path:         "/dataset/test2",
			Auth:         token2,
			Body:         nil,
			ResponseCode: 404,
			ResponseBody: map[string]interface{}{"error": "Dataset not found"},
		},
		{
			Name:         "Delete dataset - delete dataset that has files in it",
			Method:       "DELETE",
			Path:         "/dataset/test1",
			Auth:         token1,
			Body:         nil,
			ResponseCode: 400,
			ResponseBody: map[string]interface{}{"error": "Dataset has files"},
		},
		{
			Name:              "Delete dataset - delete dataset",
			Method:            "DELETE",
			Path:              "/dataset/test4",
			Auth:              token1,
			Body:              nil,
			ResponseCode:      204,
			ResponseBody:      map[string]interface{}{},
			ManualCompareBody: true,
		},
		{
			Name:         "Delete dataset - delete dataset that doesn't exist",
			Method:       "DELETE",
			Path:         "/dataset/invalid",
			Auth:         token1,
			Body:         nil,
			ResponseCode: 404,
			ResponseBody: map[string]interface{}{"error": "Dataset not found"},
		},
	}

	handleGetDatasetAttributesTests := []apitesting.Test{
		{
			Name:         "Get dataset attributes - get list of attributes for dataset",
			Method:       "GET",
			Path:         "/dataset/test1/attribute",
			Auth:         token1,
			Body:         nil,
			ResponseCode: 200,
			ResponseBody: map[string]interface{}{"attributes": []any{
				map[string]any{"id": "test1", "created_at": float64(100), "updated_at": float64(100), "dataset_id": "test1", "attribute_name": "test attribute", "attribute_value": "string"},
				map[string]any{"id": "test2", "created_at": float64(100), "updated_at": float64(100), "dataset_id": "test1", "attribute_name": "test2 attribute", "attribute_value": "string2"},
			}},
		},
		{
			Name:         "Get dataset attributes - get list of attributes for dataset with no attributes",
			Method:       "GET",
			Path:         "/dataset/test2/attribute",
			Auth:         token1,
			Body:         nil,
			ResponseCode: 200,
			ResponseBody: map[string]interface{}{"attributes": []any{}},
		},
		{
			Name:         "Get dataset attributes - get list of attributes for public dataset not owned by user",
			Method:       "GET",
			Path:         "/dataset/test3/attribute",
			Auth:         token2,
			Body:         nil,
			ResponseCode: 200,
			ResponseBody: map[string]interface{}{"attributes": []any{
				map[string]any{"id": "test13", "created_at": float64(100), "updated_at": float64(100), "dataset_id": "test3", "attribute_name": "test3 attribute", "attribute_value": "string3"},
				map[string]any{"id": "test23", "created_at": float64(100), "updated_at": float64(100), "dataset_id": "test3", "attribute_name": "test23 attribute", "attribute_value": "string23"},
			}},
		},
		{
			Name:         "Get dataset attributes - get list of attributes for private dataset not owned by user",
			Method:       "GET",
			Path:         "/dataset/test1/attribute",
			Auth:         token2,
			Body:         nil,
			ResponseCode: 404,
			ResponseBody: map[string]interface{}{"error": "Dataset not found"},
		},
		{
			Name:         "Get dataset attributes - get list of attributes for dataset that doesn't exist",
			Method:       "GET",
			Path:         "/dataset/invalid/attribute",
			Auth:         token1,
			Body:         nil,
			ResponseCode: 404,
			ResponseBody: map[string]interface{}{"error": "Dataset not found"},
		},
		// TODO: test order by
		//{
		//	Name:         "Get dataset attributes - get list of attributes for dataset test order by",
		//	Method:       "GET",
		//	Path:         "/dataset/test1/attribute?orderBy=attribute_name",
		//	Auth:         token1,
		//	Body:         nil,
		//	ResponseCode: 200,
		//	ResponseBody: map[string]interface{}{"attributes": []any{
		//		map[string]any{"id": "test1", "created_at": float64(100), "updated_at": float64(100), "dataset_id": "test1", "attribute_name": "test attribute", "attribute_value": "string"},
		//		map[string]any{"id": "test2", "created_at": float64(100), "updated_at": float64(100), "dataset_id": "test1", "attribute_name": "test2 attribute", "attribute_value": "string2"},
		//	}},
		//},
	}
	handleCreateDatasetAttributeTests := []apitesting.Test{
		{
			Name:              "Create dataset attribute - create attribute for dataset you have permission for",
			Method:            "POST",
			Path:              "/dataset/test1/attribute",
			Auth:              token1,
			Body:              map[string]interface{}{"attribute_name": "test attribute 2", "attribute_value": "string"},
			ResponseCode:      201,
			ResponseBody:      map[string]interface{}{"id": "test attribute 2", "created_at": float64(100), "updated_at": float64(100), "dataset_id": "test1", "attribute_name": "test attribute 2", "attribute_value": "string"},
			ManualCompareBody: true,
		},
		{
			Name:         "Create dataset attribute - create attribute for dataset you have permission for with no body",
			Method:       "POST",
			Path:         "/dataset/test1/attribute",
			Auth:         token1,
			Body:         nil,
			ResponseCode: 400,
			ResponseBody: map[string]interface{}{"error": "Invalid request body"},
		},
		{
			Name:         "Create dataset attribute - create attribute for dataset non-owned public dataset",
			Method:       "POST",
			Path:         "/dataset/test3/attribute",
			Auth:         token2,
			Body:         map[string]interface{}{"attribute_name": "test attribute 2", "attribute_value": "string"},
			ResponseCode: 403,
			ResponseBody: map[string]interface{}{"error": "Forbidden"},
		},
		{
			Name:         "Create dataset attribute - create attribute for dataset non-owned private dataset",
			Method:       "POST",
			Path:         "/dataset/test1/attribute",
			Auth:         token2,
			Body:         map[string]interface{}{"attribute_name": "test attribute 2", "attribute_value": "string"},
			ResponseCode: 404,
			ResponseBody: map[string]interface{}{"error": "Dataset not found"},
		},
		{
			Name:         "Create dataset attribute - create attribute for dataset invalid dataset",
			Method:       "POST",
			Path:         "/dataset/invalid/attribute",
			Auth:         token1,
			Body:         map[string]interface{}{"attribute_name": "test attribute 2", "attribute_value": "string"},
			ResponseCode: 404,
			ResponseBody: map[string]interface{}{"error": "Dataset not found"},
		},
	}
	handleUpdateDatasetAttributeTests := []apitesting.Test{
		{
			Name:              "Update dataset attribute - update attribute for dataset you have permission for",
			Method:            "PUT",
			Path:              "/dataset/test1/attribute/test1",
			Auth:              token1,
			Body:              map[string]interface{}{"attribute_name": "test attribute updated", "attribute_value": "string updated"},
			ResponseCode:      201,
			ResponseBody:      map[string]interface{}{"id": "test1", "created_at": float64(100), "updated_at": float64(100), "dataset_id": "test1", "attribute_name": "test attribute updated", "attribute_value": "string updated"},
			ManualCompareBody: true,
		},
		{
			Name:         "Update dataset attribute - update attribute for dataset you have permission for no body",
			Method:       "PUT",
			Path:         "/dataset/test1/attribute/test1",
			Auth:         token1,
			Body:         nil,
			ResponseCode: 400,
			ResponseBody: map[string]interface{}{"error": "Invalid request body"},
		},
		{
			Name:         "Update dataset attribute - update attribute for public dataset not owned by user",
			Method:       "PUT",
			Path:         "/dataset/test3/attribute/test13",
			Auth:         token2,
			Body:         map[string]interface{}{"attribute_name": "test attribute updates", "attribute_value": "string updated"},
			ResponseCode: 403,
			ResponseBody: map[string]interface{}{"error": "Forbidden"},
		},
		{
			Name:         "Update dataset attribute - update attribute for private dataset not owned by user",
			Method:       "PUT",
			Path:         "/dataset/test1/attribute/test1",
			Auth:         token2,
			Body:         map[string]interface{}{"attribute_name": "test attribute updates", "attribute_value": "string updated"},
			ResponseCode: 404,
			ResponseBody: map[string]interface{}{"error": "Dataset not found"},
		},
		{
			Name:         "Update dataset attribute - update attribute for invalid dataset",
			Method:       "PUT",
			Path:         "/dataset/invalid/attribute/test1",
			Auth:         token1,
			Body:         map[string]interface{}{"attribute_name": "test attribute updates", "attribute_value": "string updated"},
			ResponseCode: 404,
			ResponseBody: map[string]interface{}{"error": "Dataset not found"},
		},
		{
			Name:         "Update dataset attribute - update attribute for invalid attribute",
			Method:       "PUT",
			Path:         "/dataset/test1/attribute/invalid",
			Auth:         token1,
			Body:         map[string]interface{}{"attribute_name": "test attribute updates", "attribute_value": "string updated"},
			ResponseCode: 404,
			ResponseBody: map[string]interface{}{"error": "Attribute not found"},
		},
		{
			Name:         "Update dataset attribute - update attribute for invalid dataset and attribute",
			Method:       "PUT",
			Path:         "/dataset/invalid/attribute/invalid",
			Auth:         token1,
			Body:         map[string]interface{}{"attribute_name": "test attribute updates", "attribute_value": "string updated"},
			ResponseCode: 404,
			ResponseBody: map[string]interface{}{"error": "Dataset not found"},
		},
	}

	handleDeleteDatasetAttributeTests := []apitesting.Test{
		{
			Name:              "Delete dataset attribute - delete attribute for dataset you have permission for",
			Method:            "DELETE",
			Path:              "/dataset/test1/attribute/test1",
			Auth:              token1,
			Body:              nil,
			ResponseCode:      204,
			ResponseBody:      map[string]interface{}{},
			ManualCompareBody: true,
		},
		{
			Name:         "Delete dataset attribute - delete attribute that doesn't exist or is already deleted",
			Method:       "DELETE",
			Path:         "/dataset/test1/attribute/test1",
			Auth:         token1,
			Body:         nil,
			ResponseCode: 404,
			ResponseBody: map[string]interface{}{"error": "Attribute not found"},
		},
		{
			Name:         "Delete dataset attribute - delete attribute for public dataset not owned by user",
			Method:       "DELETE",
			Path:         "/dataset/test3/attribute/test13",
			Auth:         token2,
			Body:         nil,
			ResponseCode: 403,
			ResponseBody: map[string]interface{}{"error": "Forbidden"},
		},
		{
			Name:         "Delete dataset attribute - delete attribute for private dataset not owned by user",
			Method:       "DELETE",
			Path:         "/dataset/test1/attribute/test2",
			Auth:         token2,
			Body:         nil,
			ResponseCode: 404,
			ResponseBody: map[string]interface{}{"error": "Dataset not found"},
		},
		{
			Name:         "Delete dataset attribute - delete attribute for invalid dataset",
			Method:       "DELETE",
			Path:         "/dataset/invalid/attribute/test1",
			Auth:         token1,
			Body:         nil,
			ResponseCode: 404,
			ResponseBody: map[string]interface{}{"error": "Dataset not found"},
		},
		{
			Name:         "Delete dataset attribute - delete attribute for invalid attribute",
			Method:       "DELETE",
			Path:         "/dataset/test1/attribute/invalid",
			Auth:         token1,
			Body:         nil,
			ResponseCode: 404,
			ResponseBody: map[string]interface{}{"error": "Attribute not found"},
		},
		{
			Name:         "Delete dataset attribute - delete attribute for invalid dataset and attribute",
			Method:       "DELETE",
			Path:         "/dataset/invalid/attribute/invalid",
			Auth:         token1,
			Body:         nil,
			ResponseCode: 404,
			ResponseBody: map[string]interface{}{"error": "Dataset not found"},
		},
	}

	handleGetFileTests := []apitesting.Test{
		{
			Name:         "Get file - get user's file",
			Method:       "GET",
			Path:         "/file/testfile1",
			Auth:         token1,
			Body:         nil,
			ResponseCode: 200,
			ResponseBody: map[string]interface{}{"owner_name": "test1 test", "id": "testfile1", "created_at": float64(100), "updated_at": float64(100), "name": "testfile1", "owner_id": "test1", "status": "processed", "dataset_id": "test1", "dataset_name": "test dataset", "is_public": false},
		},
		{
			Name:         "Get file - get other user's private file",
			Method:       "GET",
			Path:         "/file/testfile1",
			Auth:         token2,
			Body:         nil,
			ResponseCode: 404,
			ResponseBody: map[string]interface{}{"error": "File not found"},
		},
		{
			Name:         "Get file - get public file",
			Method:       "GET",
			Path:         "/file/testfile2",
			Auth:         token2,
			Body:         nil,
			ResponseCode: 200,
			ResponseBody: map[string]interface{}{"owner_name": "test1 test", "id": "testfile2", "created_at": float64(100), "updated_at": float64(100), "name": "testfile2", "owner_id": "test1", "status": "processed", "dataset_id": "test1", "dataset_name": "test dataset", "is_public": true},
		},
		{
			Name:         "Get file - get file that doesn't exist",
			Method:       "GET",
			Path:         "/file/invalid",
			Auth:         token1,
			Body:         nil,
			ResponseCode: 404,
			ResponseBody: map[string]interface{}{"error": "File not found"},
		},
	}
	handleDeleteFileTests := []apitesting.Test{
		{
			Name:         "Delete file - delete file that is public but not owned by user",
			Method:       "DELETE",
			Path:         "/file/testfile2",
			Auth:         token2,
			Body:         nil,
			ResponseCode: 403,
			ResponseBody: map[string]interface{}{"error": "Forbidden"},
		},
		{
			Name:         "Delete file - delete file that is private but not owned by user",
			Method:       "DELETE",
			Path:         "/file/testfile1",
			Auth:         token2,
			Body:         nil,
			ResponseCode: 404,
			ResponseBody: map[string]interface{}{"error": "File not found"},
		},
		{
			Name:              "Delete file - delete file",
			Method:            "DELETE",
			Path:              "/file/testfile5",
			Auth:              token1,
			Body:              nil,
			ResponseCode:      204,
			ResponseBody:      map[string]interface{}{},
			ManualCompareBody: true,
		},
		{
			Name:         "Delete file - delete file that doesn't exist",
			Method:       "DELETE",
			Path:         "/file/invalid",
			Auth:         token1,
			Body:         nil,
			ResponseCode: 404,
			ResponseBody: map[string]interface{}{"error": "File not found"},
		},
	}
	handlePreviewTests := []apitesting.Test{
		// TODO: test preview
	}

	handleGetFileAttributesTests := []apitesting.Test{
		{
			Name:         "Get file attributes - get list of attributes for file",
			Method:       "GET",
			Path:         "/file/testfile1/attribute",
			Auth:         token1,
			Body:         nil,
			ResponseCode: 200,
			ResponseBody: map[string]interface{}{"attributes": []any{
				map[string]any{"id": "test1", "created_at": float64(100), "updated_at": float64(100), "file_id": "testfile1", "attribute_name": "test attribute", "attribute_value": "string"},
				map[string]any{"id": "test2", "created_at": float64(100), "updated_at": float64(100), "file_id": "testfile1", "attribute_name": "test2 attribute", "attribute_value": "string2"},
			},
			},
		},
		{
			Name:         "Get file attributes - get list of attributes for file with no attributes",
			Method:       "GET",
			Path:         "/file/testfile4/attribute",
			Auth:         token1,
			Body:         nil,
			ResponseCode: 200,
			ResponseBody: map[string]interface{}{"attributes": []any{}},
		},
		{
			Name:         "Get file attributes - get list of attributes for public file not owned by user",
			Method:       "GET",
			Path:         "/file/testfile2/attribute",
			Auth:         token2,
			Body:         nil,
			ResponseCode: 200,
			ResponseBody: map[string]interface{}{
				"attributes": []any{
					map[string]any{"id": "test12", "created_at": float64(100), "updated_at": float64(100), "file_id": "testfile2", "attribute_name": "test12 attribute", "attribute_value": "string12"},
					map[string]any{"id": "test22", "created_at": float64(100), "updated_at": float64(100), "file_id": "testfile2", "attribute_name": "test22 attribute", "attribute_value": "string22"},
				},
			},
		},
		{
			Name:         "Get file attributes - get list of attributes for private file not owned by user",
			Method:       "GET",
			Path:         "/file/testfile1/attribute",
			Auth:         token2,
			Body:         nil,
			ResponseCode: 404,
			ResponseBody: map[string]interface{}{"error": "File not found"},
		},
		{
			Name:         "Get file attributes - get list of attributes for file that doesn't exist",
			Method:       "GET",
			Path:         "/file/invalid/attribute",
			Auth:         token1,
			Body:         nil,
			ResponseCode: 404,
			ResponseBody: map[string]interface{}{"error": "File not found"},
		},
		// TODO: test order by
	}
	handleCreateFileAttributeTests := []apitesting.Test{
		{
			Name:              "Create file attribute - create attribute for file you have permission for",
			Method:            "POST",
			Path:              "/file/testfile1/attribute",
			Auth:              token1,
			Body:              map[string]interface{}{"attribute_name": "test attribute 3", "attribute_value": "string3"},
			ResponseCode:      201,
			ResponseBody:      map[string]interface{}{"id": "id", "created_at": float64(100), "updated_at": float64(100), "file_id": "testfile1", "attribute_name": "test attribute 3", "attribute_value": "string3"},
			ManualCompareBody: true,
		},
		{
			Name:         "Create file attribute - create attribute for file you have permission for with no body",
			Method:       "POST",
			Path:         "/file/testfile1/attribute",
			Auth:         token1,
			Body:         nil,
			ResponseCode: 400,
			ResponseBody: map[string]interface{}{"error": "Invalid request body"},
		},
		{
			Name:         "Create file attribute - create attribute for file non-owned public file",
			Method:       "POST",
			Path:         "/file/testfile2/attribute",
			Auth:         token2,
			Body:         map[string]interface{}{"attribute_name": "test attribute 2", "attribute_value": "string"},
			ResponseCode: 403,
			ResponseBody: map[string]interface{}{"error": "Forbidden"},
		},
		{
			Name:         "Create file attribute - create attribute for file non-owned private file",
			Method:       "POST",
			Path:         "/file/testfile1/attribute",
			Auth:         token2,
			Body:         map[string]interface{}{"attribute_name": "test attribute 2", "attribute_value": "string"},
			ResponseCode: 404,
			ResponseBody: map[string]interface{}{"error": "File not found"},
		},
		{
			Name:         "Create file attribute - create attribute for invalid file",
			Method:       "POST",
			Path:         "/file/invalid/attribute",
			Auth:         token1,
			Body:         map[string]interface{}{"attribute_name": "test attribute 2", "attribute_value": "string"},
			ResponseCode: 404,
			ResponseBody: map[string]interface{}{"error": "File not found"},
		},
	}
	handleUpdateFileAttributeTests := []apitesting.Test{
		{
			Name:              "Update file attribute - update attribute for file you have permission for",
			Method:            "PUT",
			Path:              "/file/testfile1/attribute/test1",
			Auth:              token1,
			Body:              map[string]interface{}{"attribute_name": "test attribute updated", "attribute_value": "string updated"},
			ResponseCode:      201,
			ResponseBody:      map[string]interface{}{"id": "test1", "created_at": float64(100), "updated_at": float64(100), "file_id": "testfile1", "attribute_name": "test attribute updated", "attribute_value": "string updated"},
			ManualCompareBody: true,
		},
		{
			Name:         "Update file attribute - update attribute for file you have permission for no body",
			Method:       "PUT",
			Path:         "/file/testfile1/attribute/test1",
			Auth:         token1,
			Body:         nil,
			ResponseCode: 400,
			ResponseBody: map[string]interface{}{"error": "Invalid request body"},
		},
		{
			Name:         "Update file attribute - update attribute for public file not owned by user",
			Method:       "PUT",
			Path:         "/file/testfile2/attribute/test12",
			Auth:         token2,
			Body:         map[string]interface{}{"attribute_name": "test attribute updates", "attribute_value": "string updated"},
			ResponseCode: 403,
			ResponseBody: map[string]interface{}{"error": "Forbidden"},
		},
		{
			Name:         "Update file attribute - update attribute for private file not owned by user",
			Method:       "PUT",
			Path:         "/file/testfile1/attribute/test1",
			Auth:         token2,
			Body:         map[string]interface{}{"attribute_name": "test attribute updates", "attribute_value": "string updated"},
			ResponseCode: 404,
			ResponseBody: map[string]interface{}{"error": "File not found"},
		},
		{
			Name:         "Update file attribute - update attribute for invalid file",
			Method:       "PUT",
			Path:         "/file/invalid/attribute/test1",
			Auth:         token1,
			Body:         map[string]interface{}{"attribute_name": "test attribute updates", "attribute_value": "string updated"},
			ResponseCode: 404,
			ResponseBody: map[string]interface{}{"error": "File not found"},
		},
		{
			Name:         "Update file attribute - update attribute for invalid attribute",
			Method:       "PUT",
			Path:         "/file/testfile1/attribute/invalid",
			Auth:         token1,
			Body:         map[string]interface{}{"attribute_name": "test attribute updates", "attribute_value": "string updated"},
			ResponseCode: 404,
			ResponseBody: map[string]interface{}{"error": "Attribute not found"},
		},
		{
			Name:         "Update file attribute - update attribute for invalid file and attribute",
			Method:       "PUT",
			Path:         "/file/invalid/attribute/invalid",
			Auth:         token1,
			Body:         map[string]interface{}{"attribute_name": "test attribute updates", "attribute_value": "string updated"},
			ResponseCode: 404,
			ResponseBody: map[string]interface{}{"error": "File not found"},
		},
	}
	handleDeleteFileAttributeTests := []apitesting.Test{
		{
			Name:              "Delete file attribute - delete attribute for file you have permission for",
			Method:            "DELETE",
			Path:              "/file/testfile1/attribute/test1",
			Auth:              token1,
			Body:              nil,
			ResponseCode:      204,
			ResponseBody:      map[string]interface{}{},
			ManualCompareBody: true,
		},
		{
			Name:         "Delete file attribute - delete attribute that doesn't exist or is already deleted",
			Method:       "DELETE",
			Path:         "/file/testfile1/attribute/test1",
			Auth:         token1,
			Body:         nil,
			ResponseCode: 404,
			ResponseBody: map[string]interface{}{"error": "Attribute not found"},
		},
		{
			Name:         "Delete file attribute - delete attribute for public file not owned by user",
			Method:       "DELETE",
			Path:         "/file/testfile2/attribute/test12",
			Auth:         token2,
			Body:         nil,
			ResponseCode: 403,
			ResponseBody: map[string]interface{}{"error": "Forbidden"},
		},
		{
			Name:         "Delete file attribute - delete attribute for private file not owned by user",
			Method:       "DELETE",
			Path:         "/file/testfile1/attribute/test2",
			Auth:         token2,
			Body:         nil,
			ResponseCode: 404,
			ResponseBody: map[string]interface{}{"error": "File not found"},
		},
		{
			Name:         "Delete file attribute - delete attribute for invalid file",
			Method:       "DELETE",
			Path:         "/file/invalid/attribute/test1",
			Auth:         token1,
			Body:         nil,
			ResponseCode: 404,
			ResponseBody: map[string]interface{}{"error": "File not found"},
		},
		{
			Name:         "Delete file attribute - delete attribute for invalid attribute",
			Method:       "DELETE",
			Path:         "/file/testfile1/attribute/invalid",
			Auth:         token1,
			Body:         nil,
			ResponseCode: 404,
			ResponseBody: map[string]interface{}{"error": "Attribute not found"},
		},
		{
			Name:         "Delete file attribute - delete attribute for invalid file and attribute",
			Method:       "DELETE",
			Path:         "/file/invalid/attribute/invalid",
			Auth:         token1,
			Body:         nil,
			ResponseCode: 404,
			ResponseBody: map[string]interface{}{"error": "File not found"},
		},
	}

	testsList := [][]apitesting.Test{
		handleGetDatasetsTests,
		handleGetDatasetTests,

		handleGetFilesForDatasetTests,
		handleCreateDatasetTests,
		handleDeleteDatasetTests,

		handleGetDatasetAttributesTests,
		handleCreateDatasetAttributeTests,

		handleUpdateDatasetAttributeTests,
		handleDeleteDatasetAttributeTests,

		handleGetFileTests,
		handleDeleteFileTests,
		handlePreviewTests,

		handleGetFileAttributesTests,
		handleCreateFileAttributeTests,

		handleUpdateFileAttributeTests,
		handleDeleteFileAttributeTests,
	}

	// TODO: Duplicate the tests to check with and without auth
	for _, tests := range testsList {
		additionalTests := make([]apitesting.Test, len(tests))
		for i, test := range tests {
			additionalTests[i] = test
			additionalTests[i].Name = test.Name + " - no auth provided"
			additionalTests[i].Auth = ""
			additionalTests[i].ResponseCode = 401
			additionalTests[i].ResponseBody = map[string]interface{}{"error": "JWT token is missing or in an invalid format"}
		}
		tests = append(tests, additionalTests...)
		fmt.Println(tests)
	}
	fmt.Printf("testList[0]: %v", testsList[0])

	// TODO: Duplicate the tests to check with and without a slash at the end of the path
	for i, tests := range testsList {
		additionalTests := make([]apitesting.Test, len(tests))
		for i, test := range tests {
			additionalTests[i] = test
			additionalTests[i].Name = test.Name + " - with slash"
			additionalTests[i].Path = test.Path + "/"
		}
		testsList[i] = append(tests, additionalTests...)
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
	{
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
		database.DatasetAttributeRepo.CreateDatasetAttribute(models.DatasetAttribute{
			Base: models.Base{
				ID:        "test1",
				CreatedAt: 100,
				UpdatedAt: 100,
			},
			DatasetID:      "test1",
			AttributeName:  "test attribute",
			AttributeValue: "string",
		})
		database.DatasetAttributeRepo.CreateDatasetAttribute(models.DatasetAttribute{
			Base: models.Base{
				ID:        "test2",
				CreatedAt: 100,
				UpdatedAt: 100,
			},
			DatasetID:      "test1",
			AttributeName:  "test2 attribute",
			AttributeValue: "string2",
		})

		database.DatasetRepo.CreateDataset(models.Dataset{
			Base: models.Base{
				ID:        "test3",
				CreatedAt: 100,
				UpdatedAt: 100,
			},
			DatasetName: "test dataset",
			OwnerID:     "test1",
			IsPublic:    true,
		})
		database.DatasetAttributeRepo.CreateDatasetAttribute(models.DatasetAttribute{
			Base: models.Base{
				ID:        "test13",
				CreatedAt: 100,
				UpdatedAt: 100,
			},
			DatasetID:      "test3",
			AttributeName:  "test3 attribute",
			AttributeValue: "string3",
		})
		database.DatasetAttributeRepo.CreateDatasetAttribute(models.DatasetAttribute{
			Base: models.Base{
				ID:        "test23",
				CreatedAt: 100,
				UpdatedAt: 100,
			},
			DatasetID:      "test3",
			AttributeName:  "test23 attribute",
			AttributeValue: "string23",
		})

		database.DatasetRepo.CreateDataset(models.Dataset{
			Base: models.Base{
				ID:        "test2",
				CreatedAt: 100,
				UpdatedAt: 100,
			},
			DatasetName: "test dataset",
			OwnerID:     "test1",
			IsPublic:    false,
		})

		database.DatasetRepo.CreateDataset(models.Dataset{
			Base: models.Base{
				ID:        "test4",
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
			OwnerID:   "test1",
			Status:    "processed",
			DatasetID: "test1",
			IsPublic:  false,
		})
		database.FileAttributeRepo.CreateFileAttribute(models.FileAttribute{
			Base: models.Base{
				ID:        "test1",
				CreatedAt: 100,
				UpdatedAt: 100,
			},
			FileID:         "testfile1",
			AttributeName:  "test attribute",
			AttributeValue: "string",
		})
		database.FileAttributeRepo.CreateFileAttribute(models.FileAttribute{
			Base: models.Base{
				ID:        "test2",
				CreatedAt: 100,
				UpdatedAt: 100,
			},
			FileID:         "testfile1",
			AttributeName:  "test2 attribute",
			AttributeValue: "string2",
		})
		database.FileRepo.CreateFile(models.File{
			Base: models.Base{
				ID:        "testfile2",
				CreatedAt: 100,
				UpdatedAt: 100,
			},
			Name:      "testfile2",
			OwnerID:   "test1",
			Status:    "processed",
			DatasetID: "test1",
			IsPublic:  true,
		})
		database.FileAttributeRepo.CreateFileAttribute(models.FileAttribute{
			Base: models.Base{
				ID:        "test12",
				CreatedAt: 100,
				UpdatedAt: 100,
			},
			FileID:         "testfile2",
			AttributeName:  "test12 attribute",
			AttributeValue: "string12",
		})
		database.FileAttributeRepo.CreateFileAttribute(models.FileAttribute{
			Base: models.Base{
				ID:        "test22",
				CreatedAt: 100,
				UpdatedAt: 100,
			},
			FileID:         "testfile2",
			AttributeName:  "test22 attribute",
			AttributeValue: "string22",
		})

		database.FileRepo.CreateFile(models.File{
			Base: models.Base{
				ID:        "testfile4",
				CreatedAt: 100,
				UpdatedAt: 100,
			},
			Name:      "testfile2",
			OwnerID:   "test1",
			Status:    "processed",
			DatasetID: "test3",
			IsPublic:  true,
		})
		database.FileRepo.CreateFile(models.File{
			Base: models.Base{
				ID:        "testfile5",
				CreatedAt: 100,
				UpdatedAt: 100,
			},
			Name:      "testfile5",
			OwnerID:   "test1",
			Status:    "processed",
			DatasetID: "test3",
			IsPublic:  true,
		})
	}

	router := InitiateServer()

	for _, tests := range testsList {
		apitesting.TestServer(t, tests, router)
	}
}
