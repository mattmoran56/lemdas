package models

type User struct {
	Base
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Avatar    string `json:"avatar"`
	Bio       string `json:"bio"`
}
