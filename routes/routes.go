package routes

import (
	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {

	// use the Default router from Gin and add the endpoints by indicating the
	// HTTP Method to use with the corresponding Function
	router := gin.Default()
	//router.POST("/signup/admin/", controllers.SignUpAdmin)
	return router
}
