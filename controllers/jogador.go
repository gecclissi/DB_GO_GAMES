package controllers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ouvermax/db_go_games/models"
)

// GET /jogador
// Pega um jogador

func PegaJogador(c *gin.Context) {
	var paramId = c.Param("id")

	jogadorID, _ := strconv.Atoi(paramId)

	var jogador = models.Jogador{
		IDJogador: jogadorID,
	}

	results := models.DB.Table("jogador").First(&jogador)
	if results.Error != nil {
		c.JSON(500, gin.H{"error": results.Error.Error()})
		return
	}
	c.JSON(200, gin.H{"data": jogador})
}

func PegaJogadorEmail(c *gin.Context) {
	var input models.JogadorInput
	if err := c.ShouldBindJSON(&input); err != nil {
		// c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.JSON(400, gin.H{
			"error":    "cannot bind JSON: " + err.Error(),
			"time":     time.Now(),
			"mensagen": "Fail",
		})
		return
	}
	jogador := models.Jogador{
		Email: input.Email,
	}

	results := models.DB.Table("jogador").Where(map[string]string{"email": input.Email}).First(&jogador)
	if results.Error != nil {
		//c.JSON(http.StatusBadRequest, gin.H{"error": results.Error.Error()})
		c.JSON(500, gin.H{
			"error":    results.Error.Error(),
			"mensagen": "Fail",
			"time":     time.Now(),
		})
		return
	}

	//c.JSON(http.StatusOK, gin.H{"data": jogador})
	c.JSON(200, gin.H{
		"mensagen":     "success",
		"DataCadastro": jogador.DataCadastro,
		"id":           jogador.IDJogador,
		"email":        jogador.Email,
	})
}

// POST /jogador
// Cria um novo Jogador
func CriaJogador(c *gin.Context) {
	// Validate input
	var input models.JogadorInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.JSON(400, gin.H{
			"error":    "cannot bind JSON: " + err.Error(),
			"time":     time.Now(),
			"mensagen": "Fail",
		})
		return
	}

	//dataAniver, err := time.Parse("2006-01-02", input.DataAniversario)
	//if err != nil {
	//	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	//	c.JSON(401, gin.H{
	//		"error":    "cannot find user: ",
	//		"time":     time.Now(),
	//		"mensagen": "Fail",
	//	})
	//	return
	//}
	// Create book
	//Nome:            input.Nome,
	//Senha:           input.Senha,
	//DataAniversario: dataAniver,
	jogador := models.Jogador{
		Email:        input.Email,
		DataCadastro: time.Now(),
		DataUltima:   time.Now(),
	}
	results := models.DB.Table("jogador").Create(&jogador)
	if results.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": results.Error.Error()})
		c.JSON(500, gin.H{
			"mensagen": "Fail",
			"time":     time.Now(),
		})
		return
	}

	//c.JSON(http.StatusOK, gin.H{"data": jogador})

	c.JSON(201, gin.H{
		"mensagen":     "success",
		"DataCadastro": jogador.DataCadastro,
		"id":           jogador.IDJogador,
	})
}

// PUT /jogador
// Atualiza um Jogador
func AtualizaJogador(c *gin.Context) {
	// Validate input
	var input models.Jogador
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.JSON(400, gin.H{
			"error":    "cannot bind JSON: " + err.Error(),
			"time":     time.Now(),
			"mensagen": "Fail",
		})
		return
	}

	var jogador models.Jogador
	jogador.IDJogador = input.IDJogador

	findResult := models.DB.Table("jogador").First(&jogador)
	if findResult.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": findResult.Error.Error()})
		c.JSON(401, gin.H{
			"error":    "cannot find user: ",
			"time":     time.Now(),
			"mensagen": "Fail",
		})
		return
	}

	if input.Nome != "" {
		jogador.Nome = input.Nome
	}

	if input.Senha != "" {
		jogador.Senha = input.Senha
	}

	if !input.DataAniversario.IsZero() {
		jogador.DataAniversario = input.DataAniversario
	}

	jogador.DataUltima = time.Now()

	results := models.DB.Table("jogador").Save(&jogador)
	if results.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": results.Error.Error()})
		c.JSON(500, gin.H{
			"error":    results.Error.Error(),
			"mensagen": "Fail",
			"time":     time.Now(),
		})
		return
	}

	//c.JSON(http.StatusOK, gin.H{"data": jogador})
	c.JSON(201, gin.H{
		"mensagen":     "success",
		"DataCadastro": jogador.DataCadastro,
		"id":           jogador.IDJogador,
	})
}

// DELETE /jogador
// Remover um Jogador
func RemoverJogador(c *gin.Context) {
	// Validate input
	var input models.Jogador
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	results := models.DB.Table("jogador").Delete(&input)
	if results.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": results.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": true})

}
