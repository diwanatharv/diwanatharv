package main

import (
	"github.com/authnull0/user-service/src/controller"
	"github.com/authnull0/user-service/src/models"
	"github.com/authnull0/user-service/src/repository"
	"github.com/gin-gonic/gin"
	"log"
)

func init() {
	Db := repository.Postgressmanager()
	err := Db.Db.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatalln("Unable to create table")
	}
}
func Makeroutes(g *gin.Engine) {
	g.POST("/signup", controller.Signup)
}
func main() {
	router := gin.Default()
	Makeroutes(router)
	err := router.Run(":8080")
	if err != nil {
		log.Fatalln("Unable to start the server", err.Error())
	}
}
