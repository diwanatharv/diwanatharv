package main

import (
	"log"

	"github.com/authnull0/user-service/src/controller"
	"github.com/authnull0/user-service/src/models"
	"github.com/authnull0/user-service/src/repository"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var env string

func init() {
	loadConfig()
	Db := repository.Postgressmanager()
	err := Db.Db.AutoMigrate(&models.Organization{})
	if err != nil {
		log.Fatalln("Unable to create table")
	}

	err = Db.Db.AutoMigrate(&models.Tenant{})
	if err != nil {
		log.Fatalln("Unable to create table")
	}

}
func Makeroutes(g *gin.Engine) {
	var orgcontroller controller.OrganizationController
	var tenantcontroller controller.TenantController
	var dashboardcontroller controller.DashboardController
	g.POST("/orgsignup", orgcontroller.SignUp)
	g.POST("/orglogin", orgcontroller.Login)
	g.POST("/orgsignupverify", orgcontroller.SignUpVerify)
	g.POST("/createtenant", tenantcontroller.CreateTenant)
	g.GET("/tenantlist", tenantcontroller.GetTenantList)
	g.POST("/dashboardnooftenant", dashboardcontroller.GetNoOfTenant)
	g.POST("/dashboardnoofuser", dashboardcontroller.GetNoOfUser)
	g.POST("/dashboardnoofendpoints", dashboardcontroller.GetNoOfEndpoints)
	g.POST("/userlist", dashboardcontroller.GetUserList)
	g.POST("/endpointlist", dashboardcontroller.GetEndpointList)
}
func loadConfig() {
	viper.SetConfigName("user-service")
	viper.AddConfigPath("..")
	viper.AddConfigPath("conf")
	viper.AddConfigPath(".")
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}
	env = viper.GetString("env") + "."

}
func main() {
	router := gin.Default()
	router.Use(CORSMiddleware())
	Makeroutes(router)
	err := router.Run(":" + viper.GetString(env+"server.port"))
	if err != nil {
		log.Fatalln("Unable to start the server", err.Error())
	}
}

func CORSMiddleware() gin.HandlerFunc {
	logrus.Info("Middleware:CORSMiddleware")
	return func(c *gin.Context) {

		allowOrigin := viper.GetString(env + "cors.allowOrigin")
		logrus.Info("Middleware:CORSMiddleware: allowOrigin: ", allowOrigin)

		c.Writer.Header().Set("Access-Control-Allow-Origin", allowOrigin)
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With, X-Authorization, X-Requesturl, withCredentials")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
