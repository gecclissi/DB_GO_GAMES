package models

type Tem struct {
	IDAuxTem     int `json:"id_aux_tem" gorm:"primaryKey"`
	IdJogo       int `json:"id_jogo"`
	IdContaPonto int `json:"id_conta_ponto"`
}

//   CREATE TABLE "Resposta"(
//     "id_Resposta" SERIAL PRIMARY KEY,
//     "nome" VARCHAR(255) NOT NULL,
//     "email" VARCHAR(255) NOT NULL,
//     "senha" VARCHAR(255) NOT NULL,
//     "data_aniversario" DATE NOT NULL,
//     "data_cadastro" DATE NOT NULL,
//     "data_ultima" DATE NOT NULL
// );

type TemInput struct {
	IdJogo       int `json:"id_jogo" binding:"required"`
	IdContaPonto int `json:"id_conta_ponto"`
}
