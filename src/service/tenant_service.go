package service

import (
	"log"

	"github.com/authnull0/user-service/src/models/dto"
	"github.com/authnull0/user-service/src/repository"
)

type TenantService struct{}

var tenantRepository repository.TenantRepository

func (t *TenantService) CreateTenant(tenant dto.CreateTenantRequest) (*dto.CreateTenantResponse, error) {

	resp, err := tenantRepository.CreateTenant(tenant)
	if err != nil {
		log.Print(err.Error())
		return resp, err
	}

	return resp, nil
}
func (t *TenantService) GetTenant(tenant dto.GetTenantListRequest) (*dto.GetTenantListResponse, error) {
	resp, err := tenantRepository.Gettenant(tenant)
	if err != nil {
		log.Print(err.Error())
		return resp, err
	}
	return resp, nil
}
