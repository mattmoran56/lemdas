package models

type Dataset struct {
	Base
	DatasetName string `json:"dataset_name"`
	OwnerID     string `json:"owner_id" gorm:"foreignKey:User(ID)"`
	IsPublic    bool   `json:"is_public" default:"false" gorm:"type:boolean""`
}
