package dto

import "github.com/authnull0/user-service/src/models"

type OrganizationRequest struct {
	FirstName            string `json:"firstname" validate:"required"`
	LastName             string `json:"lastname" validate:"required"`
	Email                string `json:"email" validate:"required,email"`
	SiteURL              string `json:"siteurl" validate:"required"`
	OrgName              string `json:"orgname"`
	Password             string `json:"password" validate:"required"`
	ConfrimPassword      string `json:"confirmpassword" validate:"required"`
	AuthenticationMethod string `json:"authentication_method" validate:"required"`
}

type OrganizationResponse struct {
	Code    int    `json:"code"`
	Status  string `json:"status"`
	Message string `json:"message"`
}

type CreateTenantRequest struct {
	TenantName string `json:"tenantname" validate:"required"`
	Email      string `json:"email" validate:"required,email"`
	Url        string `json:"url" validate:"required"`
	CreatedBy  string `json:"createdby" validate:"required"`
}

type CreateTenantResponse struct {
	Code    int    `json:"code"`
	Status  string `json:"status"`
	Message string `json:"message"`
}

type LoginResponse struct {
	Code    int    `json:"code"`
	Status  string `json:"status"`
	Message string `json:"message"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}
type Params struct {
	Memory      uint32
	Iterations  uint32
	Parallelism uint8
	SaltLength  uint32
	KeyLength   uint32
}

type VerifyEmailResponse struct {
	Code    int    `json:"code"`
	Status  string `json:"status"`
	Message string `json:"message"`
}

type GetTenantListResponse struct {
	Code    int             `json:"code"`
	Status  string          `json:"status"`
	Message string          `json:"message"`
	Data    []models.Tenant `json:"data"`
}

type DashboardRequest struct {
	Email string `json:"email" validate:"required,email"`
}

type DashboardResponse struct {
	Code    int    `json:"code"`
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    int64  `json:"data"`
}

type GetTenantListRequest struct {
	Email string `json:"email" validate:"required,email"`
}

type GetUserListRequest struct {
	Email string `json:"email" validate:"required,email"`
}

type GetUserListResponse struct {
	Code    int           `json:"code"`
	Status  string        `json:"status"`
	Message string        `json:"message"`
	Data    []models.User `json:"data"`
}

type GetEndpointListRequest struct {
	Email string `json:"email" validate:"required,email"`
}

type GetEndpointListResponse struct {
	Code    int                 `json:"code"`
	Status  string              `json:"status"`
	Message string              `json:"message"`
	Data    []models.EpmMachine `json:"data"`
}
type GetOrgListRequest struct {
	Email string `json:"email" validate:"required,email"`
}
type GetOrgListResponse struct {
	Code    int                 `json:"code"`
	Status  string              `json:"status"`
	Message string              `json:"message"`
	Data    models.Organization `json:"data"`
}

type ValidateEmailAndOrgNameRequest struct {
	Email   string `json:"email"`
	OrgName string `json:"orgname"`
}
