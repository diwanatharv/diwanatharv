package repository

import "github.com/authnull0/user-service/src/models/dto"

type DashboardRepository struct{}

func (d *DashboardRepository) GetDashboard(reqbody dto.DashboardRequest) (*dto.DashboardResponse, error) {

	var resp dto.DashboardResponse

	resp.Message = "successfully fetched dashboard data"
	resp.Code = 200
	resp.Status = "success"

	resp.Data.TotalUsers = 100
	resp.Data.TotalTenants = 100
	resp.Data.TotalEndpoints = 100

	return &resp, nil
}
