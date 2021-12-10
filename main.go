package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.Run("0.0.0.0:8000")
	fmt.Println(1)
}
