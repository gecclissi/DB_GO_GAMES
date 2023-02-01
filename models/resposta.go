package models

type Resposta struct {
	IDResposta int    `json:"id_resposta" gorm:"primaryKey"`
	Resposta   string `json:"resposta"`
	EhCorreta  bool   `json:"-"`
	IdPergunta int    `json:"-"`
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

type RespostaInput struct {
	Resposta   string `json:"resposta" binding:"required"`
	EhCorreta  bool   `json:"eh_correta"`
	IdPergunta int    `json:"id_pergunta" binding:"required"`
}
