package routes

import (
	ginexamples "github.com/LucasPLopes/gin-examples"
	"github.com/gin-gonic/gin"
)

var router = gin.Default()

func Run() {
	getRoutes()
	router.Run(ginexamples.GetPort())
}

func getRoutes() {
	v1 := router.Group("/v1")
	addPingRoutes(v1)
	addUserRoutes(v1)

	v2 := router.Group("/v2")
	addPingRoutes(v2)
}
