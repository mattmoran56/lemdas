package models

type DatasetCollaborator struct {
	Base
	UserID    string `json:"user_id" gorm:"foreignKey: User(ID)"`
	DatasetID string `json:"dataset_id" gorm:"foreignKey: Dataset(ID)"`
}
