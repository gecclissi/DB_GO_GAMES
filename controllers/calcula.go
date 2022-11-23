package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ouvermax/db_go_games/models"
)

// GET /Resposta
// Pega um Resposta
func PegaCalcula(c *gin.Context) {
	var paramId = c.Param("id")

	calculaID, _ := strconv.Atoi(paramId)

	var calcula = models.Calcula{
		IDAuxCalcula: calculaID,
	}

	results := models.DB.Table("calcula").First(&calcula)
	if results.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": results.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": calcula})
}

// POST /Resposta
// Cria um novo Resposta
func CriaCalcula(c *gin.Context) {
	// Validate input
	var input models.CalculaInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create book
	calcula := models.Calcula{
		IdRespostaRealizada: input.IdRespostaRealizada,
		IdContaPonto:        input.IdContaPonto,
	}
	results := models.DB.Table("calcula").Create(&calcula)
	if results.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": results.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": calcula})
}

// PUT /Resposta
// Atualiza um Resposta
func AtualizaCalcula(c *gin.Context) {
	// Validate input
	var input models.Calcula
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var calcula models.Calcula
	calcula.IDAuxCalcula = input.IDAuxCalcula

	findResult := models.DB.Table("calcula").First(&calcula)
	if findResult.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": findResult.Error.Error()})
		return
	}

	results := models.DB.Table("calcula").Save(&calcula)
	if results.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": results.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": calcula})
}

// DELETE /Resposta
// Remover um Resposta
func RemoverCalcula(c *gin.Context) {
	// Validate input
	var input models.Resposta
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	results := models.DB.Table("Calcula").Delete(&input)
	if results.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": results.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": true})
}
