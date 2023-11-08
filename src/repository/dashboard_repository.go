package repository

import (
	"log"

	"github.com/authnull0/user-service/src/db"
	"github.com/authnull0/user-service/src/models"
	"github.com/authnull0/user-service/src/models/dto"
)

type DashboardRepository struct{}

func (d *DashboardRepository) GetNoOfTenant(reqbody dto.DashboardRequest) (*dto.DashboardResponse, error) {
	var organization models.Organization

	db := db.Makegormserver()

	err := db.Where("admin_email = ?", reqbody.Email).First(&organization).Error

	if err != nil {
		log.Print(err.Error())
		return &dto.DashboardResponse{
			Code:    500,
			Status:  "failed",
			Message: "Not able to find organization table",
			Data:    0,
		}, nil
	}

	var count int64

	err = db.Model(&models.Tenant{}).Where("organization_id = ?", organization.Id).Count(&count).Error

	if err != nil {
		log.Print(err.Error())
		return &dto.DashboardResponse{
			Code:    500,
			Status:  "failed",
			Message: "Not able to find tenant table",
			Data:    0,
		}, nil
	}

	return &dto.DashboardResponse{
		Code:    200,
		Status:  "success",
		Message: "tenant is created successfully",
		Data:    count,
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
		}, nil
	}

	return &dto.GetUserListResponse{
		Code:    200,
		Status:  "success",
		Message: "user is created successfully",
		Data:    res,
	}, nil

}

func (d *DashboardRepository) GetNoOfUser(reqbody dto.DashboardRequest) (*dto.DashboardResponse, error) {
	var organization models.Organization

	db := db.Makegormserver()

	err := db.Where("admin_email = ?", reqbody.Email).First(&organization).Error

	if err != nil {
		log.Print(err.Error())
		return &dto.DashboardResponse{
			Code:    500,
			Status:  "failed",
			Message: "Not able to find organization table",
			Data:    0,
		}, nil
	}

	var count int64

	err = db.Model(&models.User{}).Where("org_id = ? and status = 'active'", organization.Id).Count(&count).Error

	if err != nil {
		log.Print(err.Error())
		return &dto.DashboardResponse{
			Code:    500,
			Status:  "failed",
			Message: "Not able to find user table",
			Data:    0,
		}, nil
	}

	return &dto.DashboardResponse{
		Code:    200,
		Status:  "success",
		Message: "user is created successfully",
		Data:    count,
	}, nil

}

func (d *DashboardRepository) GetNoOfEndpoints(reqbody dto.DashboardRequest) (*dto.DashboardResponse, error) {
	var organization models.Organization

	db := db.Makegormserver()

	err := db.Where("admin_email = ?", reqbody.Email).First(&organization).Error

	if err != nil {
		log.Print(err.Error())
		return &dto.DashboardResponse{
			Code:    500,
			Status:  "failed",
			Message: "Not able to find organization table",
			Data:    0,
		}, nil
	}

	var tenant []models.Tenant

	err = db.Where("organization_id = ?", organization.Id).Find(&tenant).Error

	if err != nil {
		log.Print(err.Error())
		return &dto.DashboardResponse{
			Code:    500,
			Status:  "failed",
			Message: "Not able to find endpoint table",
			Data:    0,
		}, nil
	}
	var count int64
	for _, v := range tenant {
		err = db.Model(&models.EpmMachine{}).Where("domain_id = ?", v.Id).Count(&count).Error
		if err != nil {
			log.Print(err.Error())
			return &dto.DashboardResponse{
				Code:    500,
				Status:  "failed",
				Message: "Not able to find endpoint table",
				Data:    0,
			}, nil

		}

		count = count + count
	}

	return &dto.DashboardResponse{
		Code:    200,
		Status:  "success",
		Message: "user is created successfully",
		Data:    count,
	}, nil

}

func (d *DashboardRepository) GetEndpointList(reqbody dto.GetEndpointListRequest) (*dto.GetEndpointListResponse, error) {
	var organization models.Organization
	var tenant []models.Tenant
	var res []models.EpmMachine

	var res1 []models.EpmMachine

	db := db.Makegormserver()

	err := db.Where("admin_email = ?", reqbody.Email).First(&organization).Error
	if err != nil {
		log.Print(err.Error())
		return &dto.GetEndpointListResponse{
			Code:    500,
			Status:  "failed",
			Message: "Not able to find organization table",
			Data:    nil,
		}, nil
	}

	err = db.Where("organization_id = ?", organization.Id).Find(&tenant).Error

	if err != nil {
		log.Print(err.Error())
		return &dto.GetEndpointListResponse{
			Code:    500,
			Status:  "failed",
			Message: "Not able to find tenant table",
			Data:    nil,
		}, nil
	}

	for _, v := range tenant {
		err = db.Where("domain_id = ?", v.Id).Find(&res).Error
		if err != nil {
			log.Print(err.Error())
			return &dto.GetEndpointListResponse{
				Code:    500,
				Status:  "failed",
				Message: "Not able to find endpoint table",
				Data:    nil,
			}, nil

		}

		res1 = append(res1, res...)

	}

	return &dto.GetEndpointListResponse{
		Code:    200,
		Status:  "success",
		Message: "user is created successfully",
		Data:    res1,
	}, nil

}
