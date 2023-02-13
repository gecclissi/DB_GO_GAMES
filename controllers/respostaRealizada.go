package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ouvermax/db_go_games/models"
)

// GET /Resposta
// Pega um Resposta
func PegaRespostaRealizada(c *gin.Context) {

	claims, err := GetClaimFromHeader(c)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	var resultado []models.RespostaRealizadaResultado

	resultsRepostas :=  models.DB.Raw(`select p.pergunta ,r.resposta ,r.eh_correta from resposta_realizada 
	inner join resposta r 
	on resposta_realizada.id_resposta = r .id_resposta
	inner join pergunta p  
	on r.id_pergunta = p .id_pergunta
	where id_jogador = ?`, claims.IdJogador).Scan(&resultado)
	
	if resultsRepostas.Error != nil {
		c.JSON(400, gin.H{"error": resultsRepostas.Error.Error()})
		return
	}

	c.JSON(200, gin.H{"data": resultado})


	// var paramId = c.Param("id")

	// respostarealizadaID, _ := strconv.Atoi(paramId)

	// var respostarealizada = models.RespostaRealizada{
	// 	IDRespostaRealizada: respostarealizadaID,
	// }

	// results := models.DB.Table("resposta_realizada").First(&respostarealizada)
	// if results.Error != nil {
	// 	c.JSON(400, gin.H{"error": results.Error.Error()})
	// 	return
	// }

	// var respostas []models.Resposta

	// resultsRepostas := models.DB.Table("resposta").Where(map[string]interface{}{"id_resposta": respostarealizada.IDRespostaRealizada}).Find(&respostas)
	// if resultsRepostas.Error != nil {
	// 	c.JSON(400, gin.H{"error": resultsRepostas.Error.Error()})
	// 	return
	// }

	// respostarealizada.Respostas = respostas;

	// c.JSON(200, gin.H{"data": respostarealizada})
}

// POST /Resposta
// Cria um novo Resposta
func CriaRespostaRealizada(c *gin.Context) {
	// Validate input
	var input models.RespostaRealizadaInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	claims, err := GetClaimFromHeader(c)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	respostarealizada := models.RespostaRealizada{
		IdResposta: input.IdResposta,
		IdJogador:  claims.IdJogador,
	}
	results := models.DB.Table("resposta_realizada").Create(&respostarealizada)
	if results.Error != nil {
		c.JSON(400, gin.H{"error": results.Error.Error()})
		return
	}
	var resposta = models.Resposta{
		IDResposta: input.IdResposta,
	}

	result := models.DB.Table("resposta").First(&resposta)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(200, gin.H{"data": gin.H{
		"eh_correto": resposta.EhCorreta,
	}})

}

// PUT /Resposta
// Atualiza um Resposta
func AtualizaRespostaRealizada(c *gin.Context) {
	// Validate input
	var input models.RespostaRealizada
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var respostarealizada models.RespostaRealizada
	respostarealizada.IDRespostaRealizada = input.IDRespostaRealizada

	findResult := models.DB.Table("resposta_realizada").First(&respostarealizada)
	if findResult.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": findResult.Error.Error()})
		return
	}

	results := models.DB.Table("resposta_realizada").Save(&respostarealizada)
	if results.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": results.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": respostarealizada})
}

// DELETE /Resposta
// Remover um Resposta
func RemoverRespostaRealizada(c *gin.Context) {
	// Validate input
	var input models.Resposta
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	results := models.DB.Table("resposta_realizada").Delete(&input)
	if results.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": results.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": true})
}
