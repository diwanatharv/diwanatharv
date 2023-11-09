package controller

import (
	"github.com/authnull0/user-service/src/models/dto"
	"github.com/authnull0/user-service/src/service"
	"github.com/gin-gonic/gin"
)

type DashboardController struct{}

var dashboardService service.DashboardService

func (d *DashboardController) GetUserList(g *gin.Context) {
	var reqbody dto.GetUserListRequest

	err := g.Bind(&reqbody)
	if err != nil {
		g.JSON(400, gin.H{"error": "invalid request"})
		return
	}

	resp, err := dashboardService.GetUserList(reqbody)
	if err != nil {
		g.JSON(500, gin.H{"error": "internal server error"})
		return
	}

	g.JSON(200, resp)

}

func (d *DashboardController) GetNoOfTenant(g *gin.Context) {
	var reqbody dto.DashboardRequest

	err := g.Bind(&reqbody)
	if err != nil {
		g.JSON(400, gin.H{"error": "invalid request"})
		return
	}

	resp, err := dashboardService.GetNoOfTenant(reqbody)
	if err != nil {
		g.JSON(500, gin.H{"error": "internal server error"})
		return
	}

	g.JSON(200, resp)

}

func (d *DashboardController) GetNoOfUser(g *gin.Context) {
	var reqbody dto.DashboardRequest

	err := g.Bind(&reqbody)
	if err != nil {
		g.JSON(400, gin.H{"error": "invalid request"})
		return
	}

	resp, err := dashboardService.GetNoOfUser(reqbody)
	if err != nil {
		g.JSON(500, gin.H{"error": "internal server error"})
		return
	}

	g.JSON(200, resp)

}

func (d *DashboardController) GetNoOfEndpoints(g *gin.Context) {
	var reqbody dto.DashboardRequest

	err := g.Bind(&reqbody)
	if err != nil {
		g.JSON(400, gin.H{"error": "invalid request"})
		return
	}

	resp, err := dashboardService.GetNoOfEndpoints(reqbody)
	if err != nil {
		g.JSON(500, gin.H{"error": "internal server error"})
		return
	}

	g.JSON(200, resp)

}

func (d *DashboardController) GetEndpointList(g *gin.Context) {
	var reqbody dto.GetEndpointListRequest

	err := g.Bind(&reqbody)
	if err != nil {
		g.JSON(400, gin.H{"error": "invalid request"})
		return
	}

	resp, err := dashboardService.GetEndpointList(reqbody)
	if err != nil {
		g.JSON(500, gin.H{"error": "enter a valid email"})
		return
	}

	g.JSON(200, resp)

}
