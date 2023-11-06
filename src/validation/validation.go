package validation

import (
	"errors"
	"regexp"

	"github.com/authnull0/user-service/src/models/dto"
)

func validateEmail(email string) bool {
	re := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+-/=?^_`{|}~-]+@[a-zA-Z0-9-]+(?:\\.[a-zA-Z0-9-]+)*$")
	return re.MatchString(email)
}

func validateSiteURL(siteURL string) bool {
	re := regexp.MustCompile("https?://(?:[a-zA-Z0-9-]+\\.)+[a-zA-Z]{2,6}(?:/[a-zA-Z0-9-._~/?%&=]*)?")
	return re.MatchString(siteURL)
}

func validatePassword(password string) bool {
	if len(password) < 8 {
		return false
	}

	if !regexp.MustCompile("[a-z]").MatchString(password) {
		return false
	}

	if !regexp.MustCompile("[A-Z]").MatchString(password) {
		return false
	}

	if !regexp.MustCompile("[0-9]").MatchString(password) {
		return false
	}

	return true
}

func Validate(u dto.UserRequest) error {
	// Email validation
	if !validateEmail(u.Email) {
		return errors.New("invalid email address")
	}

	// // Site URL validation
	// if !validateSiteURL(u.Url) {
	// 	return errors.New("invalid site URL")
	// }

	// Password validation
	if len(u.Password) < 8 {
		return errors.New("password must be at least 8 characters long")
	}
	if !validatePassword(u.Password) {
		return errors.New("password must have atleat 1 uppercase lowercase and number & special charachter")
	}

	return nil
}
