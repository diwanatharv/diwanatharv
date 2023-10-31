package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func Makegormserver() *gorm.DB {
	dsn := "host=65.109.82.187 user=kloudone password=authnull@kloudone@2007 dbname=authnull port=2514 "
	db, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		log.Print(err.Error())
		log.Fatalln("error in setting up db connection")
	}
	return db
}
