package middlewares

import (
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/ouvermax/db_go_games/services"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		const Bearer_schema = "Bearer "
		header := c.GetHeader("Authorization")

		fmt.Printf("HEADER - %s \n", header)

		if header == "" {
			c.AbortWithStatus(401)
		}
		token := header
		if strings.Contains(strings.ToLower(token), strings.ToLower(Bearer_schema)) {
			token = header[len(Bearer_schema):]
		}

		fmt.Printf("TOKEN - %s \n", token)

		if !services.NewJWTService().ValidateToken(token) {
			c.AbortWithStatus(401)
		}
	}
}
