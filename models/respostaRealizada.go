package models

type RespostaRealizada struct {
	IDRespostaRealizada int        `json:"id_resposta_realizada" gorm:"primaryKey"`
	IdResposta          int        `json:"id_resposta"`
	IdJogador           int        `json:"id_jogador"`
	Pergunta            []Pergunta `json:"pergunta" gorm:"-"`
	Respostas           []Resposta `json:"respostas" gorm:"-"`
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

type RespostaRealizadaInput struct {
	IdResposta int `json:"id_resposta" binding:"required"`
	IdJogador  int `json:"id_jogador"`
}

type RespostaRealizadaResultado struct {
	Pergunta  string `json:"pergunta"`
	Resposta  string `json:"resposta"`
	EhCorreta bool   `json:"eh_correta"`
}
