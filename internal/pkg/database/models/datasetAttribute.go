package models

type DatasetAttribute struct {
	Base
	DatasetID      string `gorm:"foreignKey:Dataset(ID)" json:"dataset_id"`
	AttributeName  string `json:"attribute_name"`
	AttributeValue string `json:"attribute_value"`
}
