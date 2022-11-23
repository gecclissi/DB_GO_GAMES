package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ouvermax/db_go_games/models"
)

// GET /Resposta
// Pega um Resposta
func PegaTem(c *gin.Context) {
	var paramId = c.Param("id")

	temID, _ := strconv.Atoi(paramId)

	var tem = models.Tem{
		IDAuxTem: temID,
	}

	results := models.DB.Table("tem").First(&tem)
	if results.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": results.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": tem})
}

// POST /Resposta
// Cria um novo Resposta
func CriaTem(c *gin.Context) {
	// Validate input
	var input models.TemInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create book
	tem := models.Tem{
		IdJogo:       input.IdJogo,
		IdContaPonto: input.IdContaPonto,
	}
	results := models.DB.Table("tem").Create(&tem)
	if results.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": results.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": tem})
}

// PUT /Resposta
// Atualiza um Resposta
func AtualizaTem(c *gin.Context) {
	// Validate input
	var input models.Tem
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var tem models.Tem
	tem.IDAuxTem = input.IDAuxTem

	findResult := models.DB.Table("tem").First(&tem)
	if findResult.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": findResult.Error.Error()})
		return
	}

	results := models.DB.Table("tem").Save(&tem)
	if results.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": results.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": tem})
}

// DELETE /Resposta
// Remover um Resposta
func RemoverTem(c *gin.Context) {
	// Validate input
	var input models.Resposta
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	results := models.DB.Table("tem").Delete(&input)
	if results.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": results.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": true})
}
