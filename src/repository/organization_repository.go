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

type OrganizationRepository struct{}

func (o *OrganizationRepository) SignUp(user dto.OrganizationRequest) (*dto.OrganizationResponse, error) {
	manager := Postgressmanager()
	isNotUnique, err := IsFieldNotUnique(manager.Db, "email", user.Email)
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

	//check if the organization name is unique
	isNotUnique, err = IsFieldNotUnique(manager.Db, "organization_name", user.OrgName)
	if err != nil {
		log.Print(err.Error())
		return &dto.OrganizationResponse{
			Code:    500,
			Status:  "failed",
			Message: "ERROR: Failed to check if organization name is unique",
		}, nil
	}

	// If the organization name is not unique, return an error.
	if isNotUnique {
		return &dto.OrganizationResponse{
			Code:    400,
			Status:  "failed",
			Message: "organization name already exists",
		}, nil
	}

	hashedPassword, err := GenerateFromPassword(user.Password)
	if err != nil {
		return &dto.OrganizationResponse{
			Code:    500,
			Status:  "failed",
			Message: "ERROR: Failed to hash password",
		}, nil
	}

	// Save the user to the database.

	user.Password = hashedPassword

	var organization models.Organization

	organization.AdminEmail = user.Email
	organization.OrganizationName = user.OrgName
	organization.CreatedAt = time.Now()
	organization.UpdatedAt = organization.CreatedAt
	organization.Status = "pending"
	organization.SiteURL = user.SiteURL
	organization.AuthenticationMethod = user.AuthenticationMethod

	err = manager.Insert(&organization).Error

	if err != nil {
		log.Print(err.Error())
		return &dto.OrganizationResponse{
			Code:    500,
			Status:  "failed",
			Message: "organization creation failed",
		}, nil
	}

	var users = models.User{
		FirstName:    user.FirstName,
		LastName:     user.LastName,
		EmailAddress: user.Email,
		Password:     user.Password,
		UserRoleID:   1,
		OrgID:        int(organization.Id),
		Status:       "pending",
	}

	err = manager.Insert(&users).Error

	if err != nil {
		log.Print(err.Error())
		return &dto.OrganizationResponse{
			Code:    500,
			Status:  "failed",
			Message: "user creation failed",
		}, nil
	}

	// Create a new token for the newly registered user
	token, err := CreateToken(user.Email)
	if err != nil {
		log.Print(err.Error())
		return &dto.OrganizationResponse{
			Code:    500,
			Status:  "failed",
			Message: "ERROR: Failed to create token",
		}, nil

	}

	url := fmt.Sprintf("http://authnull.com/verify?token=%s", token)

	// Send a welcome email to the user.

	message := fmt.Sprintf("<h1>Welcome to Authnull</h1><p>Hi, %s</p><p>Thank you for signing up with Authnull. We are excited to have you on board with us.</p><p>Please click on the link below to verify your email address.</p><p><a href=\"%s\">Verify Email</a></p><p>Regards,</p><p>Authnull Team</p>", user.FirstName, url)
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

func (o *OrganizationRepository) Login(loginRequest dto.LoginRequest) (*dto.LoginResponse, error) {
	manager := Postgressmanager()
	// Retrieve the user's password from the database
	user, err := GetUserByEmail(manager.Db, loginRequest.Email)
	if err != nil {
		log.Print(err.Error())
		return &dto.LoginResponse{
			Code:    401,
			Status:  "failed",
			Message: "User not registered",
		}, err
	}

	// Hash the provided password and compare it with the stored password hash
	match, err := ComparePasswordAndHash(loginRequest.Password, user.Password)
	if err != nil || !match {
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

func (o *OrganizationRepository) SignUpVerify(token string) (*dto.VerifyEmailResponse, error) {

	// Verify the token
	val, err := VerifyToken(token)
	if err != nil {
		log.Print(err.Error())
		return &dto.VerifyEmailResponse{
			Code:    401,
			Status:  "failed",
			Message: "Invalid token",
		}, err
	}

	//update the user status to active
	db := db.Makegormserver()

	err = db.Model(&models.Organization{}).Where("email = ?", val).Update("status", "active").Error
	if err != nil {
		log.Print(err.Error())
		return &dto.VerifyEmailResponse{
			Code:    500,
			Status:  "failed",
			Message: "ERROR: Failed to update organization status",
		}, nil
	}

	err = db.Model(&models.User{}).Where("email_address = ?", val).Update("status", "active").Error
	if err != nil {
		log.Print(err.Error())
		return &dto.VerifyEmailResponse{
			Code:    500,
			Status:  "failed",
			Message: "ERROR: Failed to update user status",
		}, nil
	}

	return &dto.VerifyEmailResponse{
		Code:    200,
		Status:  "success",
		Message: "user created successfully",
	}, nil
}
