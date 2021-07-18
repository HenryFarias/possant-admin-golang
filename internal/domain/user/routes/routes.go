package routes

import (
	"github.com/gin-gonic/gin"
	user "possant-admin/internal/domain/user/handler"
)

func Init(auth *gin.RouterGroup, handler user.UserHandler) {
	auth.GET("/user", handler.FindAll)
	auth.POST("/user", handler.Save)
}
