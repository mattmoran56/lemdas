package models

type FileAttribute struct {
	Base
	FileID           string `gorm:"foreignKey:File(ID)" json:"file_id"`
	AttributeName    string `json:"attribute_name"`
	AttributeValue   string `json:"attribute_value"`
	AttributeGroupID string `json:"attribute_group_id" gorm:"foreignKey:FileAttributeGroup(ID);size:191"`
}
