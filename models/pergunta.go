package models

type Pergunta struct {
	IDPergunta int    `json:"id_pergunta" gorm:"primaryKey"`
	Pergunta   string `json:"pergunta"`
	IdJogo     int    `json:"id_jogo"`
}

//   CREATE TABLE "Pergunta"(
//     "id_Pergunta" SERIAL PRIMARY KEY,
//     "nome" VARCHAR(255) NOT NULL,
//     "email" VARCHAR(255) NOT NULL,
//     "senha" VARCHAR(255) NOT NULL,
//     "data_aniversario" DATE NOT NULL,
//     "data_cadastro" DATE NOT NULL,
//     "data_ultima" DATE NOT NULL
// );

type PerguntaInput struct {
	Pergunta string `json:"pergunta" binding:"required"`
	IdJogo   int    `json:"id_jogo" binding:"required"`
}
