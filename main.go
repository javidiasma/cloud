package main

import (
	"cloudPart1/config"
	"cloudPart1/routes"
)

func main() {
	config.ConnectDatabase()
	r := routes.SetupRoutes()
	r.Run("0.0.0.0:8000")
}
