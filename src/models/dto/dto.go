package dto

type OrganizationRequest struct {
	FirstName       string `json:"firstname" validate:"required"`
	LastName        string `json:"lastname" validate:"required"`
	Email           string `json:"email" validate:"required,email"`
	OrgName         string `json:"orgname"`
	Password        string `json:"password" validate:"required"`
	ConfrimPassword string `json:"confirmpassword" validate:"required"`
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
