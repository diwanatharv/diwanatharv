package models

import "time"

type Organization struct {
	Id                   uint      `json:"id" gorm:"primary_key"`
	OrganizationName     string    `json:"organization_name"`
	AdminEmail           string    `json:"admin_email"`
	SiteURL              string    `json:"site_url"`
	CreatedAt            time.Time `json:"created_at"`
	UpdatedAt            time.Time `json:"updated_at"`
	Status               string    `json:"status"`
	AuthenticationMethod string    `json:"authentication_method"`
}
type LoginCredentials struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password"`
}
type Tenant struct {
	Id             uint      `json:"id" gorm:"primary_key"`
	TenantName     string    `json:"tenant_name"`
	AdminEmail     string    `json:"admin_email"`
	SiteURL        string    `json:"site_url"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
	OrganizationId int       `json:"organization_id"`
	Status         string    `json:"status"`
}

type User struct {
	UserId           int     `gorm:"column:user_id;primary_key"`
	EmailAddress     string  `gorm:"column:email_address"`
	PhoneNumber      string  `gorm:"column:phone_number"`
	LogOnName        string  `gorm:"column:logon_name"`
	City             string  `gorm:"column:city"`
	Country          string  `gorm:"column:country"`
	Industry         string  `gorm:"column:industry"`
	Organization     string  `gorm:"column:organization"`
	CompanyHeadcount string  `gorm:"column:company_headcount"`
	FirstName        string  `gorm:"column:firstname"`
	LastName         string  `gorm:"column:lastname"`
	Address          string  `gorm:"column:address"`
	Password         string  `gorm:"column:user_password"`
	DomainId         []uint8 `gorm:"column:domain_id"`
	Status           string  `gorm:"column:status"`
	OtpMethod        string  `gorm:"column:otp_method"`
	Metadata         string  `gorm:"column:metadata"`
	Dn               string  `gorm:"column:dn"`
	UserRoleID       int     `gorm:"column:user_role_id"`
	OrgID            int     `gorm:"column:org_id"`
}

func (User) TableName() string {
	return "users"
}
