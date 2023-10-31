package controller

import (
	"encoding/json"
	"github.com/authnull0/user-service/src/enums"
	"github.com/authnull0/user-service/src/models"
	"github.com/authnull0/user-service/src/service"
	"github.com/authnull0/user-service/src/validation"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"log"
	"net/http"
)

func Signup(g *gin.Context) {
	var reqbody models.User
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
	err = service.SignUp(reqbody)
	if err != nil {
		log.Print(err.Error())
		g.JSON(http.StatusInternalServerError, gin.H{"error": enums.ServerIssue})
		return
	}
	g.Status(http.StatusFound)
	g.Redirect(http.StatusFound, "/login")
}
