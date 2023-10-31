package service

import (
	"errors"
	"github.com/authnull0/user-service/src/db"
	"github.com/authnull0/user-service/src/models"
	"log"
)

func SignUp(user models.User) error {
	manager := db.Postgressmanager()
	isNotUnique, err := db.IsFieldNotUnique(manager.Db, "email", user.Email)
	if err != nil {
		log.Print(err.Error())
		return err
	}

	// If the email field is not unique, return an error.
	if isNotUnique {
		return errors.New("email address is already in use")
	}

	hashedPassword, err := db.HashPassword(user.Password)
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
