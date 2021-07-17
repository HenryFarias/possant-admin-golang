package main

import (
	"github.com/gin-gonic/gin"
	"possant-admin/internal/config"
	userroutes "possant-admin/internal/domain/user/routes"
)

func main() {
	db := config.Database()
	g := gin.Default()
	r := userroutes.Init(g, InjectUser(db))
	_ = r.Run(":8080")
}