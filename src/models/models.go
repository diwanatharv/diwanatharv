package models

type User struct {
	Id        uint   ` json:"id"   gorm:"primary_key"`
	FirstName string `json:"firstname" validate:"required"`
	LastName  string `json:"lastname" validate:"required"`
	Email     string `json:"email" validate:"required,email" gorm:"uniqueIndex:email"`
	OrgName   string `json:"orgname"`
	Password  string `json:"password" validate:"required"`
}
type DatabaseConfig struct {
	DbHost   string `yaml:"db.host"`
	DbPort   int    `yaml:"db.port"`
	User     string `yaml:"db.user"`
	Password string `yaml:"db.password"`
	DbName   string `yaml:"db.name"`
}
