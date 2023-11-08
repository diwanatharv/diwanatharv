package repository

import (
	"fmt"
	"log"
	"time"

	"github.com/authnull0/user-service/src/db"
	"github.com/authnull0/user-service/src/models"
	"github.com/authnull0/user-service/src/models/dto"
	"github.com/authnull0/user-service/utils"
)

type TenantRepository struct{}

func (t *TenantRepository) CreateTenant(tenant dto.CreateTenantRequest) (*dto.CreateTenantResponse, error) {
	manager := Postgressmanager()
	var tenantBody models.Tenant
	tenantBody.TenantName = tenant.TenantName
	tenantBody.AdminEmail = tenant.Email
	tenantBody.SiteURL = tenant.Url
	tenantBody.CreatedAt = time.Now()
	tenantBody.UpdatedAt = tenantBody.CreatedAt
	tenantBody.Status = "active"

	//check if the tenant name is unique
	isNotUnique, err := IsFieldNotUnique(manager.Db, "tenant_name", tenant.TenantName)

	if err != nil {
		log.Print(err.Error())
		return &dto.CreateTenantResponse{
			Code:    500,
			Status:  "failed",
			Message: "ERROR: Failed to check if tenant name is unique",
		}, nil
	}

	// If the tenant name is not unique, return an error.
	if isNotUnique {
		return &dto.CreateTenantResponse{
			Code:    400,
			Status:  "failed",
			Message: "tenant name already exists",
		}, nil
	}

	//check if the email is unique
	isNotUnique, err = IsFieldNotUnique(manager.Db, "admin_email", tenant.Email)
	if err != nil {
		log.Print(err.Error())
		return &dto.CreateTenantResponse{
			Code:    500,
			Status:  "failed",
			Message: "ERROR: Failed to check if email is unique",
		}, nil
	}

	// If the email field is not unique, return an error.
	if isNotUnique {
		return &dto.CreateTenantResponse{
			Code:    400,
			Status:  "failed",
			Message: "email already exists",
		}, nil
	}

	organization, err := GetOrganization(tenant.CreatedBy)
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
func (t *TenantRepository) Gettenant(req dto.GetTenantListRequest) (*dto.GetTenantListResponse, error) {

	var organization models.Organization
	var res []models.Tenant

	db := db.Makegormserver()

	err := db.Where("admin_email = ?", req.Email).First(&organization).Error
	if err != nil {
		log.Print(err.Error())
		return &dto.GetTenantListResponse{
			Code:    500,
			Status:  "failed",
			Message: "Not able to find organization table",
			Data:    res,
		}, nil
	}
	err = db.Where("organization_id = ?", organization.Id).Find(&res).Error

	if err != nil {
		log.Print(err.Error())
		return &dto.GetTenantListResponse{
			Code:    500,
			Status:  "failed",
			Message: "Not able to find tenant table",
			Data:    res,
		}, err
	}
	return &dto.GetTenantListResponse{
		Code:    200,
		Status:  "success",
		Message: "tenant is created successfully",
		Data:    res,
	}, nil
}
