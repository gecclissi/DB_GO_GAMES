package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ouvermax/db_go_games/models"
)

// GET /Pergunta
// Pega um Pergunta
func PegaPergunta(c *gin.Context) {
	var paramId = c.Param("id")

	perguntaID, _ := strconv.Atoi(paramId)

	var pergunta = models.Pergunta{
		IDPergunta: perguntaID,
	}

	results := models.DB.Table("pergunta").First(&pergunta)
	if results.Error != nil {
		c.JSON(400, gin.H{"error": results.Error.Error()})
		return
	}

	var respostas []models.Resposta

	resultsRepostas := models.DB.Table("resposta").Where(map[string]interface{}{"id_pergunta": pergunta.IDPergunta}).Find(&respostas)
	if resultsRepostas.Error != nil {
		c.JSON(400, gin.H{"error": resultsRepostas.Error.Error()})
		return
	}

	pergunta.Respostas = respostas;

	c.JSON(200, gin.H{"data": pergunta})
}

// POST /Pergunta
// Cria um novo Pergunta
func CriaPergunta(c *gin.Context) {
	// Validate input
	var input models.PerguntaInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// Create book
	pergunta := models.Pergunta{
		Pergunta: input.Pergunta,
		IdJogo:   input.IdJogo,
	}
	results := models.DB.Table("pergunta").Create(&pergunta)
	if results.Error != nil {
		c.JSON(400, gin.H{"error": results.Error.Error()})
		return
	}

	c.JSON(200, gin.H{"data": pergunta})
}

// PUT /Pergunta
// Atualiza um Pergunta
func AtualizaPergunta(c *gin.Context) {
	// Validate input
	var input models.Pergunta
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var pergunta models.Pergunta
	pergunta.IDPergunta = input.IDPergunta

	findResult := models.DB.Table("pergunta").First(&pergunta)
	if findResult.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": findResult.Error.Error()})
		return
	}

	if input.Pergunta != "" {
		pergunta.Pergunta = input.Pergunta
	}

	results := models.DB.Table("pergunta").Save(&pergunta)
	if results.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": results.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": pergunta})
}

// DELETE /Pergunta
// Remover um Pergunta
func RemoverPergunta(c *gin.Context) {
	// Validate input
	var input models.Pergunta
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	results := models.DB.Table("pergunta").Delete(&input)
	if results.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": results.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": true})
}
