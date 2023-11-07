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

func (d *DashboardRepository) GetUserList(reqbody dto.GetUserListRequest) (*dto.GetUserListResponse, error) {
	var organization models.Organization
	var res []models.User

	db := db.Makegormserver()

	err := db.Where("admin_email = ?", reqbody.Email).First(&organization).Error
	if err != nil {
		log.Print(err.Error())
		return &dto.GetUserListResponse{
			Code:    500,
			Status:  "failed",
			Message: "Not able to find organization table",
			Data:    nil,
		}, nil
	}

	err = db.Where("org_id = ? and status = 'active'", organization.Id).Find(&res).Error

	if err != nil {
		log.Print(err.Error())
		return &dto.GetUserListResponse{
			Code:    500,
			Status:  "failed",
			Message: "Not able to find user table",
			Data:    nil,
		}, err
	}

	return &dto.GetUserListResponse{
		Code:    200,
		Status:  "success",
		Message: "user is created successfully",
		Data:    res,
	}, nil

}
