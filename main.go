package main

import (
	"awesomeProject12/api/signup/controller"
	"awesomeProject12/pkg/data_access"
	"awesomeProject12/pkg/models"
	"github.com/gin-gonic/gin"
	"log"
)

func init() {
	Db := data_access.Postgressmanager()
	err := Db.Db.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatalln("Unable to create table")
	}
}

func main() {
	router := gin.Default()
	controller.Makeroutes(router)
	err := router.Run(":8080")
	if err != nil {
		log.Fatalln("Unable to start the server", err.Error())
	}
}
