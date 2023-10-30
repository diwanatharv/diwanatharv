package service

import (
	"awesomeProject12/pkg/data_access"
	"awesomeProject12/pkg/models"
	"errors"
	"log"
)

func SignUp(user models.User) error {
	manager := data_access.Postgressmanager()
	isNotUnique, err := data_access.IsFieldNotUnique(manager.Db, "email", user.Email)
	if err != nil {
		log.Print(err.Error())
		return err
	}

	// If the email field is not unique, return an error.
	if isNotUnique {
		return errors.New("email address is already in use")
	}

	hashedPassword, err := data_access.HashPassword(user.Password)
	if err != nil {
		return err
	}

	// Save the user to the database.
	user.Id = 0
	user.Password = hashedPassword
	err = manager.Insert(&user).Error

	if err != nil {
		log.Print(err.Error())
		return err
	}

	return nil
}
