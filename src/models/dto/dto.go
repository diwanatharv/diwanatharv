package dto

type UserRequest struct {
	FirstName       string `json:"firstname" validate:"required"`
	LastName        string `json:"lastname" validate:"required"`
	Email           string `json:"email" validate:"required,email"`
	OrgName         string `json:"orgname"`
	Password        string `json:"password" validate:"required"`
	ConfrimPassword string `json:"confirmpassword" validate:"required"`
}

type UserResponse struct {
	Code    int    `json:"code"`
	Status  string `json:"status"`
	Message string `json:"message"`
}

type CreateTenantRequest struct {
	TenantName string `json:"tenantname" validate:"required"`
	Email      string `json:"email" validate:"required,email"`
	Url        string `json:"url" validate:"required"`
}
