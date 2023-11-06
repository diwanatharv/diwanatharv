package service

import (
	"log"

	"github.com/authnull0/user-service/src/models"
	"github.com/authnull0/user-service/src/models/dto"
	"github.com/authnull0/user-service/src/repository"
	"github.com/authnull0/user-service/utils"
)

func SignUp(user dto.UserRequest) (*dto.UserResponse, error) {
	manager := repository.Postgressmanager()
	isNotUnique, err := repository.IsFieldNotUnique(manager.Db, "email", user.Email)
	if err != nil {
		log.Print(err.Error())
		return &dto.UserResponse{
			Code:    500,
			Status:  "failed",
			Message: "ERROR: Failed to check if email is unique",
		}, nil
	}

	// If the email field is not unique, return an error.
	if isNotUnique {
		return &dto.UserResponse{
			Code:    400,
			Status:  "failed",
			Message: "email already exists",
		}, nil
	}

	hashedPassword, err := repository.HashPassword(user.Password)
	if err != nil {
		return &dto.UserResponse{
			Code:    500,
			Status:  "failed",
			Message: "ERROR: Failed to hash password",
		}, nil
	}

	// Save the user to the database.

	user.Password = hashedPassword

	var users models.User

	users.FirstName = user.FirstName
	users.LastName = user.LastName
	users.Email = user.Email
	users.OrgName = user.OrgName
	users.Password = user.Password

	err = manager.Insert(&users).Error

	if err != nil {
		log.Print(err.Error())
		return &dto.UserResponse{
			Code:    500,
			Status:  "failed",
			Message: "user creation failed",
		}, nil
	}

	// Send a welcome email to the user.

	val := utils.ValidateEmail(user.Email)

	if !val {
		return &dto.UserResponse{
			Code:    400,
			Status:  "failed",
			Message: "email sending failed",
		}, nil
	}

	return &dto.UserResponse{
		Code:    200,
		Status:  "success",
		Message: "user created successfully",
	}, nil
}
