package service

import (
	"log"

	"github.com/authnull0/user-service/src/models/dto"
	"github.com/authnull0/user-service/src/repository"
)

type OrganizationService struct{}

var orgRepository repository.OrganizationRepository

func (o *OrganizationService) SignUp(user dto.OrganizationRequest) (*dto.OrganizationResponse, error) {
	resp, err := orgRepository.SignUp(user)
	if err != nil {
		log.Print(err.Error())
		return resp, err
	}

	return resp, nil

}
func (o *OrganizationService) Login(loginRequest dto.LoginRequest) (*dto.LoginResponse, error) {
	resp, err := orgRepository.Login(loginRequest)
	if err != nil {
		log.Print(err.Error())
		return resp, err
	}

	return resp, nil
}

func (o *OrganizationService) SignUpVerify(token string) (*dto.VerifyEmailResponse, error) {
	resp, err := orgRepository.SignUpVerify(token)
	if err != nil {
		log.Print(err.Error())
		return resp, err
	}

	return resp, nil
}

func (o *OrganizationService) ValidateEmailAndOrgName(email string, orgname string) (*dto.OrganizationResponse, error) {
	resp, err := orgRepository.ValidateEmailAndOrgName(email, orgname)
	if err != nil {
		log.Print(err.Error())
		return resp, err
	}

	return resp, nil
}
