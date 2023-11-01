package models

type User struct {
	Id              uint   ` json:"id"   gorm:"primary_key"`
	FirstName       string `json:"firstname" validate:"required"`
	LastName        string `json:"lastname" validate:"required"`
	Role            string `json:"role" validate:"required"`
	Email           string `json:"email" validate:"required,email" gorm:"uniqueIndex:email"`
	Url             string `json:"url" validate:"required,url"`
	OrgName         string `json:"orgname"`
	TenantName      string `json:"tenantname"`
	Password        string `json:"password" validate:"required"`
	ConfirmPassword string `json:"confirmpassword" validate:"required,eqfield=Password"`
}
type DatabaseConfig struct {
	DbHost   string `yaml:"db.host"`
	DbPort   int    `yaml:"db.port"`
	User     string `yaml:"db.user"`
	Password string `yaml:"db.password"`
	DbName   string `yaml:"db.name"`
}
