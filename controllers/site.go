package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ouvermax/db_go_games/models"
)

// GET /site
// Pega um site
func PegaSite(c *gin.Context) {
	var paramId = c.Param("id")

	siteID, _ := strconv.Atoi(paramId)

	var site = models.Site{
		IDSite: siteID,
	}

	results := models.DB.Table("site").First(&site)
	if results.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": results.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": site})
}

// POST /site
// Cria um novo site
func CriaSite(c *gin.Context) {
	// Validate input
	var input models.SiteInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create book
	site := models.Site{
		Nome:     input.Nome,
		Logotipo: input.Logotipo,
	}
	results := models.DB.Table("site").Create(&site)
	if results.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": results.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": site})
}

// PUT /site
// Atualiza um site
func AtualizaSite(c *gin.Context) {
	// Validate input
	var input models.Site
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var site models.Site
	site.IDSite = input.IDSite

	findResult := models.DB.Table("site").First(&site)
	if findResult.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": findResult.Error.Error()})
		return
	}

	if input.Nome != "" {
		site.Nome = input.Nome
	}

	results := models.DB.Table("site").Save(&site)
	if results.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": results.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": site})
}

// DELETE /site
// Remover um site
func RemoverSite(c *gin.Context) {
	// Validate input
	var input models.Site
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	results := models.DB.Table("site").Delete(&input)
	if results.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": results.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": true})
}
