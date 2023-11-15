package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/authnull0/user-service/src/controller"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var env string

func setupRoutes(engine *gin.Engine) *gin.Engine {
	var orgcontroller controller.OrganizationController
	var tenantcontroller controller.TenantController
	var dashboardcontroller controller.DashboardController
	engine.POST("/orgsignup", orgcontroller.SignUp)
	engine.POST("/orglogin", orgcontroller.Login)
	engine.GET("/orgsignupverify", orgcontroller.SignUpVerify)
	engine.POST("/orglist", orgcontroller.GetOrgList)
	engine.POST("/approveorg", orgcontroller.ApproveOrg)
	engine.POST("/createtenant", tenantcontroller.CreateTenant)
	engine.POST("/tenantlist", tenantcontroller.GetTenantList)
	engine.POST("/dashboardnooftenant", dashboardcontroller.GetNoOfTenant)
	engine.POST("/dashboardnoofuser", dashboardcontroller.GetNoOfUser)
	engine.POST("/dashboardnoofendpoints", dashboardcontroller.GetNoOfEndpoints)
	engine.POST("/userlist", dashboardcontroller.GetUserList)
	engine.POST("/endpointlist", dashboardcontroller.GetEndpointList)
	engine.POST("/validateemailandorgname", orgcontroller.ValidateEmailAndOrgName)
	engine.POST("/orgdetail", orgcontroller.GetOrg)

	return engine
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

func start() {
	loadConfig()
	gin.DisableConsoleColor()

	r := gin.Default()
	r.Use(CORSMiddleware())

	setupRoutes(r)
	// setupDatabase()
	startServer(r)

}

// startServer - Start server
func startServer(r *gin.Engine) {
	srv := &http.Server{
		Addr:    ":" + viper.GetString(env+"server.port"),
		Handler: r,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Printf("Listen: %s\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Default().Printf("Shutting down server...\n")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Default().Printf("Server forced to shutdown: %s\n", err)
	}

	log.Default().Printf("Server exiting\n")
}

func main() {
	start()
}
