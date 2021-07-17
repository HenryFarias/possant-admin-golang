package routes

import (
	"github.com/gin-gonic/gin"
	user "possant-admin/internal/domain/user/handler"
)

func Init(g *gin.Engine, handler user.UserHandler) *gin.Engine {
	g.GET("/user", handler.FindAll)
	g.POST("/user", handler.Save)
	return g
}
