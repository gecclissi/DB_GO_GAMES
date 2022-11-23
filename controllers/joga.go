package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ouvermax/db_go_games/models"
)

// GET /Resposta
// Pega um Resposta
func PegaJoga(c *gin.Context) {
	var paramId = c.Param("id")

	jogaID, _ := strconv.Atoi(paramId)

	var joga = models.Joga{
		IDAuxJoga: jogaID,
	}

	results := models.DB.Table("joga").First(&joga)
	if results.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": results.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": joga})
}

// POST /Resposta
// Cria um novo Resposta
func CriaJoga(c *gin.Context) {
	// Validate input
	var input models.JogaInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create book
	joga := models.Joga{
		IdJogo:    input.IdJogo,
		IdJogador: input.IdJogador,
	}
	results := models.DB.Table("joga").Create(&joga)
	if results.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": results.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": joga})
}

// PUT /Resposta
// Atualiza um Resposta
func AtualizaJoga(c *gin.Context) {
	// Validate input
	var input models.Joga
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var joga models.Joga
	joga.IDAuxJoga = input.IDAuxJoga

	findResult := models.DB.Table("joga").First(&joga)
	if findResult.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": findResult.Error.Error()})
		return
	}

	results := models.DB.Table("joga").Save(&joga)
	if results.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": results.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": joga})
}

// DELETE /Resposta
// Remover um Resposta
func RemoverJoga(c *gin.Context) {
	// Validate input
	var input models.Resposta
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	results := models.DB.Table("joga").Delete(&input)
	if results.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": results.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": true})
}
