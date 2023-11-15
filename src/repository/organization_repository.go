package repository

import (
	"fmt"
	"log"
	"time"

	"github.com/authnull0/user-service/src/db"
	"github.com/authnull0/user-service/src/models"
	"github.com/authnull0/user-service/src/models/dto"
	"github.com/authnull0/user-service/utils"
	"github.com/spf13/viper"
)

type OrganizationRepository struct{}

func (o *OrganizationRepository) SignUp(user dto.OrganizationRequest) (*dto.OrganizationResponse, error) {

	dbname := viper.GetString(viper.GetString("env") + ".db.dbname")

	db := db.GetConnectiontoDatabaseDynamically(dbname)

	// Check if the email is unique
	var u models.User

	err := db.Where("email_address = ?", user.Email).Find(&u).Error
	if err != nil {
		log.Print(err.Error())
		return &dto.OrganizationResponse{
			Code:    500,
			Status:  "failed",
			Message: "ERROR: Failed to check if email is unique",
		}, nil

	}

	// If the email field is not unique, return an error.
	if u.EmailAddress == user.Email {

		log.Default().Println("email on the database", u.EmailAddress)
		log.Default().Println("email on the request", user.Email)
		return &dto.OrganizationResponse{
			Code:    400,
			Status:  "failed",
			Message: "email already exists",
		}, nil
	}

	//check if the organization name is unique

	var org models.Organization

	err = db.Where("organization_name = ?", user.OrgName).Find(&org).Error
	if err != nil {
		log.Print(err.Error())
		return &dto.OrganizationResponse{
			Code:    500,
			Status:  "failed",
			Message: "ERROR: Failed to check if organization name is unique",
		}, nil
	}

	// If the organization name is not unique, return an error.

	if org.OrganizationName == user.OrgName {
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
	organization.DatabaseStatus = "pending"
	organization.SiteURL = user.SiteURL
	organization.AuthenticationMethod = user.AuthenticationMethod

	db.Create(&organization)

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
		UserRoleID:   4,
		OrgID:        int(organization.Id),
		Status:       "pending",
	}

	db.Create(&users)

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

	// Retrieve the user's password from the database
	user, err := GetUserByEmail(loginRequest.Email)
	if err != nil {
		log.Print(err.Error())
		return &dto.LoginResponse{
			Code:    401,
			Status:  "failed",
			Message: "User not registered",
		}, err
	}

	//check if organization is approved

	var org models.Organization

	dbname := viper.GetString(viper.GetString("env") + ".db.dbname")

	db := db.GetConnectiontoDatabaseDynamically(dbname)

	err = db.Where("admin_email = ?", loginRequest.Email).First(&org).Error
	if err != nil {
		log.Print(err.Error())
		return &dto.LoginResponse{
			Code:    500,
			Status:  "failed",
			Message: "Not able to find organization table",
		}, nil
	}

	if org.Status != "approved" {
		return &dto.LoginResponse{
			Code:    401,
			Status:  "failed",
			Message: "Organization is not approved",
		}, nil
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
		Message: "user logged in successfully",
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
	db := db.GetConnectiontoDatabaseDynamically("epm")

	err = db.Model(&models.Organization{}).Where("email = ?", val).Update("status", "verified").Error
	if err != nil {
		log.Print(err.Error())
		return &dto.VerifyEmailResponse{
			Code:    500,
			Status:  "failed",
			Message: "ERROR: Failed to update organization status",
		}, nil
	}

	err = db.Model(&models.User{}).Where("email_address = ?", val).Update("status", "verified").Error
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
		Message: "user verified successfully",
	}, nil
}
func (o *OrganizationRepository) GetOrg(req dto.GetOrgRequest) (*dto.GetOrgResponse, error) {
	var res models.Organization

	db := db.GetConnectiontoDatabaseDynamically("epm")
	err := db.Where("admin_email = ?", req.Email).First(&res).Error
	if err != nil {
		log.Print(err.Error())
		return &dto.GetOrgResponse{
			Code:    500,
			Status:  "failed",
			Message: "Not able to find organization table",
			Data:    res,
		}, err
	}
	return &dto.GetOrgResponse{
		Code:    200,
		Status:  "success",
		Message: "Details of organization ",
		Data:    res,
	}, nil
}

func (o *OrganizationRepository) ValidateEmailAndOrgName(email string, orgname string) (*dto.OrganizationResponse, error) {
	var message string

	var code int
	if email != "" {
		var u models.User
		db := db.GetConnectiontoDatabaseDynamically(viper.GetString(viper.GetString("env") + ".db.dbname"))
		err := db.Where("email_address = ?", email).Find(&u).Error
		if err != nil {
			log.Print(err.Error())
			return &dto.OrganizationResponse{
				Code:    500,
				Status:  "failed",
				Message: "ERROR: Failed to check if email is unique",
			}, nil
		}

		// If the email field is not unique, return an error.

		if u.EmailAddress == email {
			message = "email already exists"

			code = 400

		} else {
			message = "email is does not exist"

			code = 200
		}
	} else if orgname != "" {
		var org models.Organization
		db := db.GetConnectiontoDatabaseDynamically(viper.GetString(viper.GetString("env") + ".db.dbname"))
		err := db.Where("organization_name = ?", orgname).Find(&org).Error
		if err != nil {
			log.Print(err.Error())
			return &dto.OrganizationResponse{
				Code:    500,
				Status:  "failed",
				Message: "ERROR: Failed to check if organization name is unique",
			}, nil
		}

		// If the organization name is not unique, return an error.

		if org.OrganizationName == orgname {
			message = "organization name already exists"

			code = 400

		} else {
			message = "organization name does not exist"

			code = 200
		}
	}

	return &dto.OrganizationResponse{
		Code:    code,
		Status:  "success",
		Message: message,
	}, nil
}

func (o *OrganizationRepository) GetOrgList(req dto.GetOrgListRequest) (*dto.GetOrgListResponse, error) {
	var res []models.Organization

	offset := (req.PageNo - 1) * req.PageSize

	dbname := viper.GetString(viper.GetString("env") + ".db.dbname")

	db := db.GetConnectiontoDatabaseDynamically(dbname)

	err := db.Model(&models.Organization{}).Where("status = 'verified'").Offset(offset).Limit(req.PageSize).Find(&res).Error
	if err != nil {
		log.Print(err.Error())
		return &dto.GetOrgListResponse{
			Code:    500,
			Status:  "failed",
			Message: "Not able to find organization table",
			Data:    res,
		}, err
	}

	return &dto.GetOrgListResponse{
		Code:    200,
		Status:  "success",
		Message: "Details of organization ",
		Data:    res,
	}, nil
}

func (o *OrganizationRepository) ApproveOrg(req dto.ApproveOrgRequest) (*dto.ApproveOrgResponse, error) {
	db := db.GetConnectiontoDatabaseDynamically(viper.GetString(viper.GetString("env") + ".db.dbname"))

	err := db.Model(&models.Organization{}).Where("id = ?", req.OrgId).Update("status", "approved").Error
	if err != nil {
		log.Print(err.Error())
		return &dto.ApproveOrgResponse{
			Code:    500,
			Status:  "failed",
			Message: "ERROR: Failed to update organization status",
		}, nil
	}

	return &dto.ApproveOrgResponse{
		Code:    200,
		Status:  "success",
		Message: "organization approved successfully",
	}, nil
}
