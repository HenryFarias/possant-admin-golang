package main

import (
	"github.com/gin-gonic/gin"
	"possant-admin/internal/auth"
	"possant-admin/internal/config"
	userroutes "possant-admin/internal/domain/user/routes"
)

func main() {
	db := config.Database()
	g := gin.Default()
	group := auth.Init(g)
	userroutes.Init(group, InjectUser(db))
	_ = g.Run(":8080")
}
