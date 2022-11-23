package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ouvermax/db_go_games/models"
)

// GET /ContaPonto
// Pega um ContaPonto
func PegaContaPonto(c *gin.Context) {
	var paramId = c.Param("id")

	contapontoID, _ := strconv.Atoi(paramId)

	var contaponto = models.ContaPonto{
		IDContaPonto: contapontoID,
	}

	results := models.DB.Table("conta_ponto").First(&contaponto)
	if results.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": results.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": contaponto})
}

// POST /ContaPonto
// Cria um novo ContaPonto
func CriaContaPonto(c *gin.Context) {
	// Validate input
	var input models.ContaPontoInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create book
	contaponto := models.ContaPonto{
		Pontos:    input.Pontos,
		IdJogador: input.IdJogador,
	}
	results := models.DB.Table("conta_ponto").Create(&contaponto)
	if results.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": results.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": contaponto})
}

// PUT /ContaPonto
// Atualiza um ContaPonto
func AtualizaContaPonto(c *gin.Context) {
	// Validate input
	var input models.ContaPonto
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var contaponto models.ContaPonto
	contaponto.IDContaPonto = input.IDContaPonto

	findResult := models.DB.Table("conta_ponto").First(&contaponto)
	if findResult.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": findResult.Error.Error()})
		return
	}



	results := models.DB.Table("conta_ponto").Save(&contaponto)
	if results.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": results.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": contaponto})
}

// DELETE /ContaPonto
// Remover um ContaPonto
func RemoverContaPonto(c *gin.Context) {
	// Validate input
	var input models.ContaPonto
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	results := models.DB.Table("conta_ponto").Delete(&input)
	if results.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": results.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": true})
}
