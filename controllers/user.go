package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ouvermax/db_go_games/models"
)

// GET /User
// Pega um User
func PegaUser(c *gin.Context) {
	var paramId = c.Param("id")

	userID, _ := strconv.Atoi(paramId)

	var user = models.Users{
		Userid: userID,
	}

	results := models.DB.Table("users").First(&user)
	if results.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": results.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": user})
}

// POST /User
// Cria um novo User
func CriaUser(c *gin.Context) {
	// Validate input
	var input models.UsersInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create book
	user := models.Users{
		Email: input.Email,
	}
	results := models.DB.Table("users").Create(&user)
	if results.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": results.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": user})
}

// PUT /User
// Atualiza um User
func AtualizaUser(c *gin.Context) {
	// Validate input
	var input models.Users
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user models.Users
	user.Userid = input.Userid

	findResult := models.DB.Table("users").First(&user)
	if findResult.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": findResult.Error.Error()})
		return
	}

	results := models.DB.Table("users").Save(&user)
	if results.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": results.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": user})
}

// DELETE /User
// Remover um User
func RemoverUser(c *gin.Context) {
	// Validate input
	var input models.Users
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	results := models.DB.Table("users").Delete(&input)
	if results.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": results.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": true})
}
