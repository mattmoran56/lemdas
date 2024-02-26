package models

type StaredDataset struct {
	Base
	UserID    string `json:"user_id" gorm:"foreignKey: User(ID)"`
	DatasetID string `json:"dataset_id" gorm:"foreignKey: Dataset(ID)"`
}
