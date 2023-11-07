package db

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Makegormserver() *gorm.DB {

	env := viper.GetString("env") + "."
	user := viper.GetString(env + "db.user")
	password := viper.GetString(env + "db.password")
	host := viper.GetString(env + "db.host")
	port := viper.GetString(env + "db.port")
	dbname := viper.GetString(env + "db.name")
	schema := viper.GetString(env + "db.schema")
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable search_path=%s", host, user, password, dbname, port, schema)
	db, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		log.Print(err.Error())
		log.Fatalln("error in setting up db connection")
	}
	return db
}
