package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/ouvermax/db_go_games/models"
	"github.com/ouvermax/db_go_games/services"
)

func Login(c *gin.Context) {
	db := models.DB

	var p models.Login
	err := c.ShouldBindJSON(&p)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot bind JSON: " + err.Error(),
		})
		return
	}

	var user models.Users
	dbError := db.Where("email = ?", p.Email).First(&user).Error
	if dbError != nil {
		c.JSON(400, gin.H{
			"error": "cannot find user: ",
		})
		return
	}

	token, err := services.NewJWTService().GenerateToken(user.Userid)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"token": token,
	})
}
