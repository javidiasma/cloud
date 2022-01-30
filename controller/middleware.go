package controller

import (
	"bytes"
	"cloudPart1/utils"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

type Status struct {
	Status string `json:"status"`
}

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := utils.GetToken(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			c.Abort()
		}
		requestBody, err1 := json.Marshal(map[string]string{
			"token": token,
		})
		if err1 != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err1.Error()})
			c.Abort()
		}
		response, err2 := http.Post("http://localhost:8001/validateUser/", "application/json", bytes.NewBuffer(requestBody))
		if err2 != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err2.Error()})
			c.Abort()
		}
		body, err3 := ioutil.ReadAll(response.Body)
		if err3 != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err3.Error()})
			c.Abort()
		}
		var temp Status
		err = json.Unmarshal(body, &temp)
		fmt.Println(body)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			c.Abort()
		}
		if temp.Status != "validated" {
			c.JSON(http.StatusBadRequest, gin.H{"message": temp.Status})
			c.Abort()
		}
		c.Next()
	}
}
