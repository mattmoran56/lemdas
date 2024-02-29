package models

type UserShareDataset struct {
	Base
	UserID      string `json:"user_id" gorm:"foreignKey:User(ID);size:191"`
	DatasetID   string `json:"dataset_id" gorm:"foreignKey:Dataset(ID);size:191"`
	WriteAccess bool   `json:"write_access"`
	User        User   `json:"user" gorm:"foreignKey:UserID"`
}
