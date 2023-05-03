package database

import (
	"github.com/ouvermax/db_go_games/models"
)

type PerguntaRespondida struct {
	EhCorreta  bool   `json:"eh_correta"`
	IDPergunta int    `json:"id_pergunta"`
	Pergunta   string `json:"pergunta"`
}

func PegarPerguntaRespondidas(IdJogo, IdJogador int) ([]PerguntaRespondida, error) {
	var resultados []PerguntaRespondida

	resultsRepostas := models.DB.Raw(`
	select r.eh_correta , p.id_pergunta ,p.pergunta 
	from resposta_realizada ra
	inner join resposta r
	on ra.id_resposta = r .id_resposta
	inner join pergunta p
	on r.id_pergunta = p .id_pergunta
	where id_jogador = ? and p.id_jogo = ?
	`, IdJogador, IdJogo).Scan(&resultados)

	if resultsRepostas.Error != nil {
		return nil, resultsRepostas.Error
	}
	return resultados, nil
}

func PegarTotalPerguntasJogo(IdJogo int) (totalPergunta int, err error) {

	resultsRepost := models.DB.Raw(`select count(*) from pergunta where id_jogo = ?`, IdJogo).Scan(&totalPergunta)
	err = resultsRepost.Error

	return
}
