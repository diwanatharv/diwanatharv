package service

import (
	"errors"
	"github.com/authnull0/user-service/src/models"
	"github.com/authnull0/user-service/src/repository"
	"log"
)

func SignUp(user models.User) error {
	manager := repository.Postgressmanager()
	isNotUnique, err := repository.IsFieldNotUnique(manager.Db, "email", user.Email)
	if err != nil {
		log.Print(err.Error())
		return err
	}

	// If the email field is not unique, return an error.
	if isNotUnique {
		return errors.New("email address is already in use")
	}

	hashedPassword, err := repository.HashPassword(user.Password)
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
