package models

type File struct {
	Base
	Name      string
	OwnerId   string `gorm:"foreignKey:User(ID)"`
	Status    string `gorm:"default:'uploaded'"`
	DatasetID string `gorm:"foreignKey:Dataset(ID)"`
	IsPublic  bool   `gorm:"default:false"`
}
