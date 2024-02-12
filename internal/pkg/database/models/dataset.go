package models

type Dataset struct {
	Base
	DatasetName string `json:"dataset_name"`
	OwnerID     string `json:"owner_id" gorm:"foreignKey:User(ID)"`
}
