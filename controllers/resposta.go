package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ouvermax/db_go_games/models"
)

// GET /Resposta
// Pega um Resposta
func PegaResposta(c *gin.Context) {
	var paramId = c.Param("id")

	respostaID, _ := strconv.Atoi(paramId)

	var resposta = models.Resposta{
		IDResposta: respostaID,
	}

	results := models.DB.Table("resposta").First(&resposta)
	if results.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": results.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": resposta})
}

// POST /Resposta
// Cria um novo Resposta
func CriaResposta(c *gin.Context) {
	// Validate input
	var input models.RespostaInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create book
	resposta := models.Resposta{
		Resposta:   input.Resposta,
		EhCorreta: input.EhCorreta,
		IdPergunta: input.IdPergunta,
	}
	results := models.DB.Table("resposta").Create(&resposta)
	if results.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": results.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": resposta})
}

// PUT /Resposta
// Atualiza um Resposta
func AtualizaResposta(c *gin.Context) {
	// Validate input
	var input models.Resposta
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var resposta models.Resposta
	resposta.IDResposta = input.IDResposta

	findResult := models.DB.Table("resposta").First(&resposta)
	if findResult.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": findResult.Error.Error()})
		return
	}

	if input.Resposta != "" {
		resposta.Resposta = input.Resposta
	}



	results := models.DB.Table("resposta").Save(&resposta)
	if results.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": results.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": resposta})
}

// DELETE /Resposta
// Remover um Resposta
func RemoverResposta(c *gin.Context) {
	// Validate input
	var input models.Resposta
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	results := models.DB.Table("resposta").Delete(&input)
	if results.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": results.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": true})
}
