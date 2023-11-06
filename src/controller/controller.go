package controller

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/authnull0/user-service/src/enums"
	"github.com/authnull0/user-service/src/models/dto"
	"github.com/authnull0/user-service/src/service"
	"github.com/authnull0/user-service/src/validation"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func Signup(g *gin.Context) {
	var reqbody dto.OrganizationRequest
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
	var reqbody dto.LoginRequest
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
	var reqbody dto.CreateTenantRequest
	err := g.Bind(&reqbody)
	if err != nil {
		log.Print(err.Error())
		g.JSON(http.StatusBadRequest, gin.H{"error": enums.Invalid})
		return
	}
	resp, err := service.CreateTenant(reqbody)

	if err != nil {
		log.Print(err.Error())
		g.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}
	g.JSON(200, resp)

}
