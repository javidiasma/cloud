package controller

import (
	"cloudPart1/DAL"
	"fmt"
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

func CreateDB(c *gin.Context) {
	fmt.Println(1)
	_, err := DAL.CsvReader()
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "done"})
		return
	}
}
func GetDB(c *gin.Context) {
	fmt.Println(2)
	data, err := DAL.GetDatabase()
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	} else {
		c.JSON(http.StatusOK, data)
		return
	}
}

func GetByRank(c *gin.Context) {
	fmt.Println(3)
	rank, _ := strconv.Atoi(c.Param("rank"))
	data, err := DAL.GetByRank(rank)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, data)
	return
}

func GetByName(c *gin.Context) {
	name := c.Param("name")
	data, err := DAL.GetByName(name)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, data)
	return
}

func GetByPlatform(c *gin.Context) {
	platform := c.Param("platform")
	num, _ := strconv.Atoi(c.Param("num"))
	data, err := DAL.GetByPlatform(platform, num)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, data)
	return
}

func GetByYear(c *gin.Context) {
	year, _ := strconv.Atoi(c.Param("year"))
	num, _ := strconv.Atoi(c.Param("num"))
	data, err := DAL.GetByYear(year, num)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, data)
	return
}

func GetByGenre(c *gin.Context) {
	genre := c.Param("genre")
	num, _ := strconv.Atoi(c.Param("num"))
	data, err := DAL.GetByGenre(genre, num)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, data)
	return
}

func GetByBestSellers(c *gin.Context) {
	year, _ := strconv.Atoi(c.Param("year"))
	platform := c.Param("platform")
	data, err := DAL.GetByBestSellers(year, platform)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, data)
	return
}

func GetEUMoreThanNA(c *gin.Context) {
	data, err := DAL.GetEUMoreThanNA()
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, data)
	return
}
