package controller

import (
	"cloudPart1/DAL"
	"strconv"

	"github.com/gin-gonic/gin"
	"net/http"
)

func ByRank(c *gin.Context) {

	rank, _ := strconv.Atoi(c.Param("rank"))

	data := DAL.ReturnByRank(rank)
	if len(data) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "error exits"})
		return
	} else {
		c.JSON(http.StatusOK, data)
		return
	}
}

func Database(c *gin.Context) {
	_, err := DAL.CsvReader()
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "done"})
		return
	}
}
