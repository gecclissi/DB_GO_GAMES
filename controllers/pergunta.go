package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ouvermax/db_go_games/database"
	"github.com/ouvermax/db_go_games/models"
)

// GET /Pergunta
// Pega um Pergunta

func total(perguntaRespondidas []database.PerguntaRespondida) (corretas, incorretas int) {
	for _, perguntaRespondida := range perguntaRespondidas {
		if perguntaRespondida.EhCorreta {
			corretas = corretas + 1
		} else {
			incorretas = incorretas + 1
		}
	}
	return
}

func PegaProxima2(c *gin.Context) {
	jogoID, _ := strconv.Atoi(c.Param("id"))
	claims, err := GetClaimFromHeader(c)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	perguntaRespondidas, err := database.PegarPerguntaRespondidas(jogoID, claims.IdJogador)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	corretas, incorretas := total(perguntaRespondidas)
	
	var totalPergunta int
	

	totalRespondida := corretas + incorretas
	totalPendente := totalPergunta - totalRespondida

	temProxima := totalPendente != 0

	fmt.Println(temProxima)	
}

func PegaProxima(c *gin.Context) {
	var paramId = c.Param("id")

	jogoID, _ := strconv.Atoi(paramId)

	claims, err := GetClaimFromHeader(c)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	var resultados []models.RespostaRealizadaRes

	//select pega nome he correta total

	resultsRepostas := models.DB.Raw(`
	select count(r.eh_correta) as total,r.eh_correta, j.nome
	from resposta_realizada ra
	inner join resposta r
	on ra.id_resposta = r .id_resposta
	inner join pergunta p
	on r.id_pergunta = p .id_pergunta
	inner join jogo j
	on p.id_jogo = j .id_jogo
	where id_jogador = ? and j.id_jogo = ?
	group by r.eh_correta,j.nome ,j.id_jogo
	`, claims.IdJogador, jogoID).Scan(&resultados)

	if resultsRepostas.Error != nil {
		c.JSON(400, gin.H{"error": resultsRepostas.Error.Error()})
		return
	}

	var (
		correta models.RespostaRealizadaRes
		errada  models.RespostaRealizadaRes
	)

	for _, resultado := range resultados {
		if resultado.EhCorreta {
			correta = resultado
		} else {
			errada = resultado
		}

	}

	var totalPergunta int

	resultsRepost := models.DB.Raw(`select count(*) from pergunta where id_jogo = ?`, jogoID).Scan(&totalPergunta)

	if resultsRepost.Error != nil {
		c.JSON(400, gin.H{"error": resultsRepost.Error.Error()})
		return
	}

	totalRespondida := correta.Total + errada.Total
	totalPendente := totalPergunta - totalRespondida

	temProxima := totalPendente != 0

	if temProxima {

		parameId := totalRespondida + 1

		perguntaID := parameId

		var pergunta = models.Pergunta{
			IDPergunta: perguntaID,
		}

		results := models.DB.Table("pergunta").First(&pergunta)
		if results.Error != nil {
			c.JSON(400, gin.H{"error": results.Error.Error()})
			return
		}

		var respostas []models.Resposta

		resultsRpostas := models.DB.Table("resposta").Where(map[string]interface{}{"id_pergunta": pergunta.IDPergunta}).Find(&respostas)
		if resultsRpostas.Error != nil {
			c.JSON(400, gin.H{"error": resultsRpostas.Error.Error()})
			return
		}

		pergunta.Respostas = respostas

		c.JSON(200, gin.H{"data": pergunta})
	} else {
		c.JSON(200, gin.H{"data": models.Respostasjogo{
			TemProxima: temProxima,
			InfosRespostas: models.InfosRespostas{
				Corretas:        correta.Total,
				Erradas:         errada.Total,
				Pendentes:       totalPendente,
				TotalRespondida: totalRespondida,
				Total:           totalPergunta,
			},
		}})
	}
}

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

	pergunta.Respostas = respostas

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
