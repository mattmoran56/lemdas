package models

type UserGroup struct {
	Base
	GroupName string `json:"group_name"`
	OwnerID   string `json:"owner_id" gorm:"foreignKey:User(id);size:191"`
}
