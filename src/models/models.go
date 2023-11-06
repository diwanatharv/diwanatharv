package models

import "time"

type Organization struct {
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
