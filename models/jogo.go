package models

import "time"

type Jogo struct {
	IDJogo      int       `json:"id_jogo" gorm:"primaryKey"`
	Nome        string    `json:"nome"`
	DataInicial time.Time `json:"data_inicial"`
	DataFinal   time.Time `json:"data_final"`
	IdFase      int       `json:"id_fase"`
}

//   CREATE TABLE "jogo"(
//     "id_jogo" SERIAL PRIMARY KEY,
//     "nome" VARCHAR(255) NOT NULL,
//     "email" VARCHAR(255) NOT NULL,
//     "senha" VARCHAR(255) NOT NULL,
//     "data_aniversario" DATE NOT NULL,
//     "data_cadastro" DATE NOT NULL,
//     "data_ultima" DATE NOT NULL
// );

type JogoInput struct {
	Nome      string `json:"nome" binding:"required"`
	DataFinal string `json:"data_final" `
	IdFase    int    `json:"id_fase" binding:"required"`
}
