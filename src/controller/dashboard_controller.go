package controller

import (
	"github.com/authnull0/user-service/src/models/dto"
	"github.com/authnull0/user-service/src/service"
	"github.com/gin-gonic/gin"
)

type DashboardController struct{}

var dashboardService service.DashboardService

func (d *DashboardController) GetDashboard(g *gin.Context) {
	var reqbody dto.DashboardRequest

	err := g.Bind(&reqbody)
	if err != nil {
		g.JSON(400, gin.H{"error": "invalid request"})
		return
	}

	resp, err := dashboardService.GetDashboard(reqbody)
	if err != nil {
		g.JSON(500, gin.H{"error": "internal server error"})
		return
	}

	g.JSON(200, resp)

}
