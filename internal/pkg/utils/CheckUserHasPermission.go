package utils

import (
	"github.com/mattmoran/fyp/api/pkg/database"
	"go.uber.org/zap"
)

func CheckUserHasPermission(fileId string, userId string) bool {
	zap.S().Debug("fileId: ", fileId, " userId: ", userId)
	file, err := database.FileRepo.GetFileByID(fileId)
	if err != nil {
		zap.S().Debug("Error getting file by id to check user permission ", err)
		return false
	}

	if file.OwnerId != userId {
		zap.S().Debug("User does not have permission to access file ", file.OwnerId, userId)
		return false
	}
	// TODO: Check shared files

	return true
}
