package models

type GroupShareDataset struct {
	Base
	GroupID     string    `json:"group_id" gorm:"foreignKey:Group(ID);size:191"`
	DatasetID   string    `json:"dataset_id" gorm:"foreignKey:Dataset(ID);size:191"`
	WriteAccess bool      `json:"write_access"`
	UserGroup   UserGroup `json:"group" gorm:"foreignKey:GroupID"`
}
