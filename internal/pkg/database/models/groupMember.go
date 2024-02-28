package models

type GroupMember struct {
	Base
	GroupID string `json:"group_id" gorm:"foreignKey:Group(ID);size:191"`
	UserID  string `json:"user_id" gorm:"foreignKey:User(id);size:191"`
	User    User   `json:"user" gorm:"foreignKey:UserID"`
}
