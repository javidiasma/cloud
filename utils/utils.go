package utils

import (
	"errors"
	"github.com/gin-gonic/gin"
	"strings"
)

func GetToken(ctx *gin.Context) (string, error) {
	authHeader := ctx.Request.Header.Get("Authorization")
	tokenParts := strings.Split(authHeader, " ")
	if len(tokenParts) != 2 {
		return "", errors.New("unauthorized1")
	}
	return tokenParts[1], nil
}
