package controllers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ouvermax/db_go_games/models"
)

// GET /fase
// Pega um fase
func PegaFaseJogo(c *gin.Context) {
	var parameId = c.Param("id")

	jogoID, _ := strconv.Atoi(parameId)

	var jogo = models.Jogo{
		IdFase: jogoID,
	}

	var jogos []models.Jogo

	resultes := models.DB.Table("jogo").Find(&jogos).Where(jogo)
	if resultes.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": resultes.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": jogos})
}

func PegaFase(c *gin.Context) {
	var paramId = c.Param("id")

	faseID, _ := strconv.Atoi(paramId)

	var fase = models.Fase{
		IDFase: faseID,
	}

	results := models.DB.Table("fase").First(&fase)
	if results.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": results.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": fase})
}

// POST /fase
// Cria um novo fase
func CriaFase(c *gin.Context) {
	// Validate input
	var input models.FaseInput
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
	fase := models.Fase{
		Nome:        input.Nome,
		DataInicial: time.Now(),
		DataFinal:   dataFinal,
		IdSite:      input.IdSite,
	}
	results := models.DB.Table("fase").Create(&fase)
	if results.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": results.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": fase})
}

// PUT /fase
// Atualiza um fase
func AtualizaFase(c *gin.Context) {
	// Validate input
	var input models.Fase
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var fase models.Fase
	fase.IDFase = input.IDFase

	findResult := models.DB.Table("fase").First(&fase)
	if findResult.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": findResult.Error.Error()})
		return
	}

	if input.Nome != "" {
		fase.Nome = input.Nome
	}

	if !input.DataFinal.IsZero() {
		fase.DataFinal = input.DataFinal
	}

	results := models.DB.Table("fase").Save(&fase)
	if results.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": results.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": fase})
}

// DELETE /fase
// Remover um fase
func RemoverFase(c *gin.Context) {
	// Validate input
	var input models.Fase
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	results := models.DB.Table("fase").Delete(&input)
	if results.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": results.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": true})
}
