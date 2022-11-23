package controllers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ouvermax/db_go_games/models"
)

// GET /jogo
// Pega um jogo
func PegaJogo(c *gin.Context) {
	var paramId = c.Param("id")

	jogoID, _ := strconv.Atoi(paramId)

	var jogo = models.Jogo{
		IDJogo: jogoID,
	}

	results := models.DB.Table("jogo").First(&jogo)
	if results.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": results.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": jogo})
}

// POST /jogo
// Cria um novo jogo
func CriaJogo(c *gin.Context) {
	// Validate input
	var input models.JogoInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	dataFinal, err := time.Parse("2006-01-02", input.DataFinal)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Create book
	jogo := models.Jogo{
		Nome:        input.Nome,
		DataInicial: time.Now(),
		DataFinal:   dataFinal,
		IdFase:      input.IdFase,
	}
	results := models.DB.Table("jogo").Create(&jogo)
	if results.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": results.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": jogo})
}

// PUT /jogo
// Atualiza um jogo
func AtualizaJogo(c *gin.Context) {
	// Validate input
	var input models.Jogo
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var jogo models.Jogo
	jogo.IDJogo = input.IDJogo

	findResult := models.DB.Table("jogo").First(&jogo)
	if findResult.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": findResult.Error.Error()})
		return
	}

	if input.Nome != "" {
		jogo.Nome = input.Nome
	}

	if !input.DataFinal.IsZero() {
		jogo.DataFinal = input.DataFinal
	}

	results := models.DB.Table("jogo").Save(&jogo)
	if results.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": results.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": jogo})
}

// DELETE /jogo
// Remover um jogo
func RemoverJogo(c *gin.Context) {
	// Validate input
	var input models.Jogo
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	results := models.DB.Table("jogo").Delete(&input)
	if results.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": results.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": true})
}
