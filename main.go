package main

import (
	"cloudPart1/config"
	"cloudPart1/routes"
	"fmt"
)

func main() {
	config.ConnectDatabase()
	r := routes.SetupRoutes()
	err := r.Run("0.0.0.0:8000")

	if err != nil {
		fmt.Println(err.Error())
	}
}
