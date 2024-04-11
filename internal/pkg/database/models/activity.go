package models

type Activity struct {
	Base
	Type    string `json:"type"`
	UserID  string `json:"user_id"gorm:"foreignKey:ID;references:User(ID)"`
	Details string `json:"details"`
}
