package models

type ContaPonto struct {
	IDContaPonto int    `json:"id_conta_ponto" gorm:"primaryKey"`
	Pontos   int `json:"pontos"`
	IdJogador     int    `json:"id_Jogador"`
}


type ContaPontoInput struct {
	Pontos int `json:"pontos" binding:"required"`
	IdJogador   int    `json:"id_Jogador" binding:"required"`
}
