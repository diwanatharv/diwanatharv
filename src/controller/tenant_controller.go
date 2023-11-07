package controller

import (
	"log"
	"net/http"

	"github.com/authnull0/user-service/src/enums"
	"github.com/authnull0/user-service/src/models/dto"
	"github.com/authnull0/user-service/src/service"
	"github.com/gin-gonic/gin"
)

type TenantController struct{}

var tenantService service.TenantService

func (t *TenantController) CreateTenant(g *gin.Context) {
	var reqbody dto.CreateTenantRequest
	err := g.Bind(&reqbody)
	if err != nil {
		log.Print(err.Error())
		g.JSON(http.StatusBadRequest, gin.H{"error": enums.Invalid})
		return
	}
	resp, err := tenantService.CreateTenant(reqbody)

	if err != nil {
		log.Print(err.Error())
		g.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}
	g.JSON(200, resp)

}
func (t *TenantController) GetTenantList(g *gin.Context) {
	var reqbody dto.GetTenantListRequest
	err := g.Bind(&reqbody)
	if err != nil {
		log.Print(err.Error())
		g.JSON(http.StatusBadRequest, gin.H{"error": enums.Invalid})
		return
	}

	resp, err := tenantService.GetTenant(reqbody)
	if err != nil {
		log.Print(err.Error())
		g.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}
	g.JSON(200, resp)
}
