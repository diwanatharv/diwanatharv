package service

import (
	"fmt"
	"log"
	"time"

	"github.com/authnull0/user-service/src/models"
	"github.com/authnull0/user-service/src/models/dto"
	"github.com/authnull0/user-service/src/repository"
	"github.com/authnull0/user-service/utils"
)

func SignUp(user dto.OrganizationRequest) (*dto.OrganizationResponse, error) {
	manager := repository.Postgressmanager()
	isNotUnique, err := repository.IsFieldNotUnique(manager.Db, "email", user.Email)
	if err != nil {
		log.Print(err.Error())
		return &dto.OrganizationResponse{
			Code:    500,
			Status:  "failed",
			Message: "ERROR: Failed to check if email is unique",
		}, nil
	}

	// If the email field is not unique, return an error.
	if isNotUnique {
		return &dto.OrganizationResponse{
			Code:    400,
			Status:  "failed",
			Message: "email already exists",
		}, nil
	}

	hashedPassword, err := repository.HashPassword(user.Password)
	if err != nil {
		return &dto.OrganizationResponse{
			Code:    500,
			Status:  "failed",
			Message: "ERROR: Failed to hash password",
		}, nil
	}

	// Save the user to the database.

	user.Password = hashedPassword

	var users models.Organization

	users.FirstName = user.FirstName
	users.LastName = user.LastName
	users.Email = user.Email
	users.OrgName = user.OrgName
	users.Password = user.Password

	err = manager.Insert(&users).Error

	if err != nil {
		log.Print(err.Error())
		return &dto.OrganizationResponse{
			Code:    500,
			Status:  "failed",
			Message: "user creation failed",
		}, nil
	}

	// Send a welcome email to the user.

	message := fmt.Sprintf("<h1>Welcome to Authnull</h1><p>Hi, %s</p><p>Thank you for signing up with Authnull. We are excited to have you on board with us.</p><p>Regards,</p><p>Authnull Team</p>", user.Email)
	val := utils.ValidateEmail(user.Email, message)

	if !val {
		return &dto.OrganizationResponse{
			Code:    400,
			Status:  "failed",
			Message: "email sending failed",
		}, nil
	}

	return &dto.OrganizationResponse{
		Code:    200,
		Status:  "success",
		Message: "user created successfully",
	}, nil
}
func Login(loginRequest dto.LoginRequest) (*dto.LoginResponse, error) {
	manager := repository.Postgressmanager()
	// Retrieve the user's password from the database
	user, err := repository.GetUserByEmailForOrganization(manager.Db, loginRequest.Email)
	if err != nil {
		log.Print(err.Error())
		return &dto.LoginResponse{
			Code:    401,
			Status:  "failed",
			Message: "User not registered",
		}, err
	}

	// Hash the provided password and compare it with the stored password hash
	err = repository.Checkpassword(loginRequest.Password, user.Password)
	if err != nil {
		log.Print(err.Error())
		return &dto.LoginResponse{
			Code:    401,
			Status:  "failed",
			Message: "Incorrect Password,Please try again",
		}, err
	}

	return &dto.LoginResponse{
		Code:    200,
		Status:  "success",
		Message: "user created successfully",
	}, nil
}
func CreateTenant(tenant dto.CreateTenantRequest) (*dto.CreateTenantResponse, error) {

	manager := repository.Postgressmanager()
	var tenantBody models.Tenant
	tenantBody.TenantName = tenant.TenantName
	tenantBody.AdminEmail = tenant.Email
	tenantBody.SiteURL = tenant.Url
	tenantBody.CreatedAt = time.Now()
	tenantBody.UpdatedAt = tenantBody.CreatedAt
	tenantBody.Status = "active"

	organization, err := repository.GetUserByEmailForOrganization(manager.Db, tenant.CreatedBy)
	if err != nil {
		log.Print(err.Error())
		return &dto.CreateTenantResponse{
			Code:    401,
			Status:  "failed",
			Message: "User not registered",
		}, err
	}

	tenantBody.OrganizationId = int(organization.Id)

	//insert the tenant to database
	err = manager.Insert(&tenantBody).Error

	if err != nil {
		log.Print(err.Error())
		return &dto.CreateTenantResponse{
			Code:    500,
			Status:  "failed",
			Message: "tenant creation failed",
		}, nil
	}

	// Send a welcome email to the user.

	message := fmt.Sprintf("<h1>Welcome to Authnull</h1><p>Hi, %s</p><p>You have been added as an admin to the tenant %s. Please login to the tenant portal to manage the tenant.</p><p>Regards,</p><p>Authnull Team</p>", tenant.Email, tenant.TenantName)
	val := utils.ValidateEmail(tenant.Email, message)
	if !val {
		return &dto.CreateTenantResponse{
			Code:    400,
			Status:  "failed",
			Message: "email sending failed",
		}, nil
	}

	return &dto.CreateTenantResponse{
		Code:    200,
		Status:  "success",
		Message: "tenant is created successfully",
	}, nil
}
