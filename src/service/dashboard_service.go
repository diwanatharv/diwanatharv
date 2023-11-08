package service

import (
	"log"

	"github.com/authnull0/user-service/src/models/dto"
	"github.com/authnull0/user-service/src/repository"
)

type DashboardService struct{}

var dashboardRepository repository.DashboardRepository

func (d *DashboardService) GetUserList(reqbody dto.GetUserListRequest) (*dto.GetUserListResponse, error) {
	resp, err := dashboardRepository.GetUserList(reqbody)
	if err != nil {
		log.Print(err.Error())
		return resp, err
	}

	return resp, nil
}

func (d *DashboardService) GetNoOfTenant(reqbody dto.DashboardRequest) (*dto.DashboardResponse, error) {
	resp, err := dashboardRepository.GetNoOfTenant(reqbody)
	if err != nil {
		log.Print(err.Error())
		return resp, err
	}

	return resp, nil
}

func (d *DashboardService) GetNoOfUser(reqbody dto.DashboardRequest) (*dto.DashboardResponse, error) {
	resp, err := dashboardRepository.GetNoOfUser(reqbody)
	if err != nil {
		log.Print(err.Error())
		return resp, err
	}

	return resp, nil
}

func (d *DashboardService) GetNoOfEndpoints(reqbody dto.DashboardRequest) (*dto.DashboardResponse, error) {
	resp, err := dashboardRepository.GetNoOfEndpoints(reqbody)
	if err != nil {
		log.Print(err.Error())
		return resp, err
	}

	return resp, nil
}

func (d *DashboardService) GetEndpointList(reqbody dto.GetEndpointListRequest) (*dto.GetEndpointListResponse, error) {
	resp, err := dashboardRepository.GetEndpointList(reqbody)
	if err != nil {
		log.Print(err.Error())
		return resp, err
	}

	return resp, nil
}
