package routes

import (
	"cloudPart1/controller"
	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {

	router := gin.Default()

	router.GET("/createDB/", controller.CreateDB)
	//"/byRank/:rank"
	router.GET("/getDB/", controller.GetDB)
	router.GET("/byRank/:rank", controller.GetByRank)
	router.GET("/byName/:name", controller.GetByName)
	router.GET("/byPlatform/:platform/:num", controller.GetByPlatform)
	router.GET("/byYear/:year/:num", controller.GetByYear)
	router.GET("/byGenre/:genre/:num", controller.GetByGenre)
	router.GET("/by5BestSellers/:year/:platform", controller.GetByBestSellers)
	router.GET("/byEUMoreThanNA/", controller.GetEUMoreThanNA)
	return router
}
