package models

type File struct {
	Base
	Name      string `json:"name"`
	OwnerID   string `gorm:"foreignKey:User(ID)" json:"owner_id"`
	Status    string `gorm:"default:'uploaded'" json:"status"`
	DatasetID string `gorm:"foreignKey:Dataset(ID)" json:"dataset_id"`
}
