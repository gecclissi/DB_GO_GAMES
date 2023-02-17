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

	claims, err := GetClaimFromHeader(c)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	var resultados []models.RespostaRealizadaRes

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
		nomeDoJogo string
		correta    models.RespostaRealizadaRes
		errada     models.RespostaRealizadaRes
	)

	for _, resultado := range resultados {
		if resultado.EhCorreta {
			correta = resultado
		} else {
			errada = resultado
		}

		if resultado.Nome != "" {
			nomeDoJogo = resultado.Nome
		}
	}

	// resultsRepost := models.DB.Raw(`select count(e.eh_correta) total
	// from resposta_realizada la
	// inner join resposta e
	// on la.id_resposta = e .id_resposta
	// where id_jogador = ?
	// `, claims.IdJogador).Scan(&result)

	// if resultsRepost.Error != nil {
	// 	c.JSON(400, gin.H{"error": resultsRepost.Error.Error()})
	// 	return
	// }

	var totalPergunta int

	resultsRepost := models.DB.Raw(`select count(*) from pergunta where id_jogo = ?`, jogoID).Scan(&totalPergunta)

	if resultsRepost.Error != nil {
		c.JSON(400, gin.H{"error": resultsRepost.Error.Error()})
		return
	}

	totalRespondida := correta.Total + errada.Total
	totalPendente := totalPergunta - totalRespondida

	temProxima := totalPendente != 0

	c.JSON(200, gin.H{"data": models.Respostasjogo{
		Nome:       nomeDoJogo,
		TemProxima: temProxima,
		InfosRespostas: models.InfosRespostas{
			Corretas:        correta.Total,
			Erradas:         errada.Total,
			Pendentes:       totalPendente,
			TotalRespondida: totalRespondida,
			Total:           totalPergunta,
		},
	}})

	// var paramId = c.Param("id")

	// jogoID, _ := strconv.Atoi(paramId)

	// var jogo = models.Jogo{
	// 	IDJogo: jogoID,
	// }

	// results := models.DB.Table("jogo").First(&jogo)
	// if results.Error != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": results.Error.Error()})
	// 	return
	// }
	// c.JSON(http.StatusOK, gin.H{"data": jogo})
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
