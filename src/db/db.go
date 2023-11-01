package db

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func Makegormserver() *gorm.DB {

	env := viper.GetString("env") + "."
	user := viper.GetString(env + "db.user")
	password := viper.GetString(env + "db.password")
	host := viper.GetString(env + "db.host")
	port := viper.GetString(env + "db.port")
	dbname := viper.GetString(env + "db.name")
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable ", host, user, password, dbname, port)
	db, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		log.Print(err.Error())
		log.Fatalln("error in setting up db connection")
	}
	return db
}
