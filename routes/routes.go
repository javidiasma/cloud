package routes

import (
	"cloudPart1/controller"
	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {

	// use the Default router from Gin and add the endpoints by indicating the
	// HTTP Method to use with the corresponding Function
	router := gin.Default()
	router.GET("/byRank/:rank", controller.ByRank)
	router.POST("/createDB/", controller.Database)

	return router
}
