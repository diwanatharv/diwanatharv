package controller

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/authnull0/user-service/src/models"

	"github.com/authnull0/user-service/src/enums"
	"github.com/authnull0/user-service/src/models/dto"
	"github.com/authnull0/user-service/src/service"
	"github.com/authnull0/user-service/src/validation"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func Signup(g *gin.Context) {
	var reqbody dto.UserRequest
	err := json.NewDecoder(g.Request.Body).Decode(&reqbody)
	if err != nil {
		log.Print(err.Error())
		g.JSON(http.StatusBadRequest, gin.H{"error": enums.Invalid})
		return
	}
	v := validator.New()
	err = v.Struct(&reqbody)
	if err != nil {
		log.Print(err.Error())
		g.JSON(http.StatusBadRequest, gin.H{"error": "validation failed" + err.Error()})
		return
	}
	err = validation.Validate(reqbody)
	if err != nil {
		log.Print(err.Error())
		g.JSON(http.StatusBadRequest, gin.H{"error": "validation failed" + err.Error()})
		return
	}
	resp, err := service.SignUp(reqbody)
	if err != nil {
		log.Print(err.Error())
		g.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	g.JSON(http.StatusOK, resp)
}
func Login(g *gin.Context) {
	var reqbody models.LoginCredentials
	err := g.Bind(&reqbody)
	if err != nil {
		log.Print(err.Error())
		g.JSON(http.StatusBadRequest, gin.H{"error": enums.Invalid})
		return
	}
	v := validator.New()
	err = v.Struct(&reqbody)
	if err != nil {
		log.Print(err.Error())
		g.JSON(http.StatusBadRequest, gin.H{"error": "validation failed" + err.Error()})
		return
	}
	resp, err := service.Login(reqbody)
	if err != nil {
		log.Print(err.Error())
		g.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	g.JSON(http.StatusOK, resp)
}
func CreateTenant(g *gin.Context) {
	var reqbody models.Tenant
	err := g.Bind(&reqbody)
	if err != nil {
		log.Print(err.Error())
		g.JSON(http.StatusBadRequest, gin.H{"error": enums.Invalid})
		return
	}
	err = service.CreateTenant(reqbody)

	if err != nil {
		log.Print(err.Error())
		g.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}
	g.JSON(200, gin.H{
		"code":    "200",
		"status":  "success",
		"message": "tenant is created successfully",
	})

}
