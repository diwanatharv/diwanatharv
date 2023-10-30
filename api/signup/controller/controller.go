package controller

import (
	"github.com/gin-gonic/gin"
)

func Makeroutes(g *gin.Engine) {
	g.POST("/signup", signup)
}
