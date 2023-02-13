package controllers

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/ouvermax/db_go_games/services"
)

func GetClaimFromHeader(c *gin.Context) (*services.Claim, error) {
	token := c.GetHeader("Authorization")

	token = strings.ReplaceAll(token, "Bearer ", "")

	claims, err := services.NewJWTService().GetClaimFromToken(token)
	if err != nil {
		return nil, err
	}

	return &claims, nil
}
