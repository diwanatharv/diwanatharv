package main

import (
	"log"

	"github.com/authnull0/user-service/src/controller"
	"github.com/authnull0/user-service/src/models"
	"github.com/authnull0/user-service/src/repository"
	"github.com/gin-gonic/gin"
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
	g.POST("/dashboard", dashboardcontroller.GetDashboard)

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
	Makeroutes(router)
	err := router.Run(":" + viper.GetString(env+"server.port"))
	if err != nil {
		log.Fatalln("Unable to start the server", err.Error())
	}
}
