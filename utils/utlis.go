package utils

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
	"gopkg.in/gomail.v2"
)

func ValidateEmail(email string) bool {
	m := gomail.NewMessage()

	m.SetHeader("From", "support@authnull.com")
	m.SetHeader("To", email)
	m.SetHeader("Subject", "Welcome to Authnull")

	env := viper.GetString("env")

	Host := viper.GetString(env + ".email.host")

	Port := viper.GetInt(env + ".email.port")

	From := viper.GetString(env + ".email.username")

	Credential := viper.GetString(env + ".email.password")

	m.SetBody("text/html", fmt.Sprintf("<h1>Welcome to Authnull</h1><p>Hi, %s</p><p>Thank you for signing up with Authnull. We are excited to have you on board with us.</p><p>Regards,</p><p>Authnull Team</p>", email))

	d := gomail.NewDialer(Host, Port, From, Credential)
	err := d.DialAndSend(m)
	if err != nil {
		return false
	}
	log.Default().Println("Email sent successfully!")
	return true
}
