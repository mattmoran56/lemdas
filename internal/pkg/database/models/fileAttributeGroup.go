package models

type FileAttributeGroup struct {
	Base
	AttributeGroupName string               `json:"attribute_group_name"`
	FileID             string               `json:"file_id"`
	ParentGroupID      *string              `json:"parent_group_id" gorm:"foreignKey:ID;size:191;null"`
	Children           []FileAttributeGroup `json:"children" gorm:"foreignKey:ParentGroupID"`
	Attributes         []FileAttribute      `json:"attributes" gorm:"foreignKey:AttributeGroupID"`
}
