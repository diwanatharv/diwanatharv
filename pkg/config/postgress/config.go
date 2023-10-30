package postgress

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func Makegormserver() *gorm.DB {
	dsn := "host=localhost user=postgres password=atharv12345 dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		log.Print(err.Error())
		log.Fatalln("error in setting up db connection")
	}
	return db
}
