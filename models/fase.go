package models

import (
	"time"
)

type Fase struct {
	IDFase      int       `json:"id_fase" gorm:"primaryKey"`
	Nome        string    `json:"nome"`
	DataInicial time.Time `json:"data_inicial"`
	DataFinal   time.Time `json:"data_final"`
	IdSite      int       `json:"id_site"`
}

//   CREATE TABLE "fase"(
//     "id_fase" SERIAL PRIMARY KEY,
//     "nome" VARCHAR(255) NOT NULL,
//     "email" VARCHAR(255) NOT NULL,
//     "senha" VARCHAR(255) NOT NULL,
//     "data_aniversario" DATE NOT NULL,
//     "data_cadastro" DATE NOT NULL,
//     "data_ultima" DATE NOT NULL
// );

type FaseInput struct {
	Nome      string `json:"nome" binding:"required"`
	DataFinal string `json:"data_final" `
	IdSite    int    `json:"id_site" binding:"required"`
}
