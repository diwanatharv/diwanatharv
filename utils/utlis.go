package utils

import (
	"log"

	"github.com/spf13/viper"
	"gopkg.in/gomail.v2"
)

func ValidateEmail(email string, message string) bool {
	m := gomail.NewMessage()

	m.SetHeader("From", "support@authnull.com")
	m.SetHeader("To", email)
	m.SetHeader("Subject", "Welcome to Authnull")

	env := viper.GetString("env")

	Host := viper.GetString(env + ".email.host")

	Port := viper.GetInt(env + ".email.port")

	From := viper.GetString(env + ".email.username")

	Credential := viper.GetString(env + ".email.password")

	log.Default().Println(Host, Port, From, Credential)

	m.SetBody("text/html", message)
	d := gomail.NewDialer(Host, Port, From, Credential)
	err := d.DialAndSend(m)
	if err != nil {
		log.Default().Println("Email sending failed!", err)
		return false
	}
	log.Default().Println("Email sent successfully!")
	return true
}
