package models

type Joga struct {
	IDAuxJoga    int `json:"id_aux_joga" gorm:"primaryKey"`
	IdJogo    int `json:"id_jogo"`
	IdJogador int `json:"id_jogador"`
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

type JogaInput struct {
	IdJogo    int `json:"id_jogo" binding:"required"`
	IdJogador int `json:"id_jogador" binding:"required"`
}
