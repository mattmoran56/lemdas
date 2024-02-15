package utils

import "github.com/mattmoran/fyp/api/pkg/database"

func CheckUserHasPermission(fileId string, userId string) bool {
	file, err := database.FileRepo.GetFileByID(fileId)
	if err != nil {
		return false
	}

	if file.OwnerId != userId {
		return false
	}
	// TODO: Check shared files

	return true
}
