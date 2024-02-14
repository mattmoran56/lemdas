package models

type FileAttribute struct {
	Base
	FileID         string `gorm:"foreignKey:File(ID)" json:"file_id"`
	AttributeName  string `json:"attribute_name"`
	AttributeValue string `json:"attribute_value"`
}
