package models

type User struct {
	Id        uint   ` json:"id"   gorm:"primary_key"`
	FirstName string `json:"firstname" validate:"required"`
	LastName  string `json:"lastname" validate:"required"`
	Email     string `json:"email" validate:"required,email" gorm:"uniqueIndex:email"`
	OrgName   string `json:"orgname"`
	Password  string `json:"password" validate:"required"`
}
type LoginCredentials struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password"`
}
