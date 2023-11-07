package repository

import (
	"log"

	"github.com/authnull0/user-service/src/db"
	"github.com/authnull0/user-service/src/models"
	"github.com/authnull0/user-service/src/models/dto"
)

type DashboardRepository struct{}

func (d *DashboardRepository) GetDashboard(reqbody dto.DashboardRequest) (*dto.DashboardResponse, error) {

	db := db.Makegormserver()

	var organization models.Organization

	err := db.Model(&models.Organization{}).Where("admin_email = ?", reqbody.Email).First(&organization).Error
	if err != nil {
		return &dto.DashboardResponse{
			Code:    500,
			Message: "failed to fetch dashboard data",
			Status:  "failed",
		}, nil
	}

	var tenant []models.Tenant

	err = db.Model(&models.Tenant{}).Where("organization_id = ?", organization.Id).Find(&tenant).Error

	if err != nil {
		log.Default().Print(err.Error())
		return &dto.DashboardResponse{
			Code:    500,
			Message: "failed to fetch dashboard data",
			Status:  "failed",
		}, nil
	}

	var user []models.User

	err = db.Model(&models.User{}).Where("org_id = ?", organization.Id).Find(&user).Error

	if err != nil {
		log.Default().Println(err.Error())
		return &dto.DashboardResponse{
			Code:    500,
			Message: "failed to fetch dashboard data",
			Status:  "failed",
		}, nil
	}

	var Data dto.DashboardData

	Data.TotalUsers = len(user)
	Data.TotalTenants = len(tenant)
	Data.TotalEndpoints = 100

	return &dto.DashboardResponse{
		Code:    200,
		Message: "successfully fetched dashboard data",
		Status:  "success",
		Data:    Data,
	}, nil
}
