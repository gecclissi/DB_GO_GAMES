package models

type Site struct {
	IDSite   int    `json:"id_site" gorm:"primaryKey"`
	Nome     string `json:"nome"`
	Logotipo string    `json:"logotipo"`
}

//   CREATE TABLE "site"(
//     "id_site" SERIAL PRIMARY KEY,
//     "nome" VARCHAR(255) NOT NULL,
//     "email" VARCHAR(255) NOT NULL,
//     "senha" VARCHAR(255) NOT NULL,
//     "data_aniversario" DATE NOT NULL,
//     "data_cadastro" DATE NOT NULL,
//     "data_ultima" DATE NOT NULL
// );

type SiteInput struct {
	Nome     string `json:"nome" binding:"required"`
	Logotipo string    `json:"logotipo" binding:"required"`
}
