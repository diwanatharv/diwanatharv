package service

import (
	"log"

	"github.com/authnull0/user-service/src/models/dto"
	"github.com/authnull0/user-service/src/repository"
)

type DashboardService struct{}

var dashboardRepository repository.DashboardRepository

func (d *DashboardService) GetDashboard(reqbody dto.DashboardRequest) (*dto.DashboardResponse, error) {
	resp, err := dashboardRepository.GetDashboard(reqbody)
	if err != nil {
		log.Print(err.Error())
		return resp, err
	}

	return resp, nil
}
