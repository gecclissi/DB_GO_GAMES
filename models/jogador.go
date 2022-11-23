package models

import "time"

type Jogador struct {
	IDJogador       int       `json:"id_jogador" gorm:"primaryKey"`
	Nome            string    `json:"nome"`
	Email           string    `json:"email"`
	Senha           string    `json:"senha"`
	DataAniversario time.Time `json:"data_aniversario"`
	DataCadastro    time.Time `json:"data_cadastro"`
	DataUltima      time.Time `json:"data_ultima"`
}

//   CREATE TABLE "jogador"(
//     "id_jogador" SERIAL PRIMARY KEY,
//     "nome" VARCHAR(255) NOT NULL,
//     "email" VARCHAR(255) NOT NULL,
//     "senha" VARCHAR(255) NOT NULL,
//     "data_aniversario" DATE NOT NULL,
//     "data_cadastro" DATE NOT NULL,
//     "data_ultima" DATE NOT NULL
// );

type JogadorInput struct {
	Nome            string `json:"nome" binding:"required"`
	Email           string `json:"email" binding:"required"`
	Senha           string `json:"senha" binding:"required"`
	DataAniversario string `json:"data_aniversario" `
}
