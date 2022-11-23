-- create database games_show

CREATE TABLE "site" (
    "id_site" SERIAL PRIMARY KEY,
    "nome" VARCHAR(255) NOT NULL,
    "logotipo" VARCHAR(255) NOT NULL
);
-- ALTER TABLE `site` ADD PRIMARY KEY `site_id_site_primary`(`id_site`);
CREATE TABLE "fase"(
    "id_fase" SERIAL PRIMARY KEY,
    "nome" VARCHAR(255) NOT NULL,
    "data_inicial" DATE NOT NULL,
    "data_final" DATE NOT NULL,
    "id_site" INT NOT NULL
);
-- ALTER TABLE "fase" ADD PRIMARY KEY "fase_id_fase_primary"("id_fase");
-- ALTER table "fase" ADD UNIQUE "fase_id_site_unique"("id_site");
CREATE TABLE "jogo"(
    "id_jogo" SERIAL PRIMARY KEY,
    "nome" VARCHAR(255) NOT NULL,
    "data_inicial" DATE NOT NULL,
    "data_final" DATE NOT NULL,
    "id_fase" INT NOT NULL
);
-- ALTER TABLE "jogo" ADD PRIMARY KEY "jogo_id_jogo_primary"("id_jogo");
-- ALTER TABLE "jogo" ADD UNIQUE "jogo_id_fase_unique"("id_fase");
CREATE TABLE "joga"(
	"id_aux_joga"	SERIAL PRIMARY KEY,
    "id_jogo" INT NOT NULL,
    "id_jogador" INT NOT NULL
);
-- ALTER TABLE "joga" ADD UNIQUE "joga_id_jogo_unique"("id_jogo");
-- ALTER TABLE "joga" ADD UNIQUE "joga_id_jogador_unique"("id_jogador");
CREATE TABLE "jogador"(
    "id_jogador" SERIAL PRIMARY KEY,
    "nome" VARCHAR(255) NOT NULL,
    "email" VARCHAR(255) NOT NULL,
    "senha" VARCHAR(255) NOT NULL,
    "data_aniversario" DATE NOT NULL,
    "data_cadastro" DATE NOT NULL,
    "data_ultima" DATE NOT NULL
);
-- ALTER table "jogador" ADD PRIMARY KEY "jogador_id_jogador_primary"("id_jogador");
CREATE TABLE "resposta_realizada"(
    "id_resposta_realizada" SERIAL PRIMARY KEY,
    "id_resposta" INT NOT NULL,
    "id_jogador" INT NOT NULL
);
-- ALTER table "resposta_realizada" ADD PRIMARY KEY "resposta_realizada_id_resposta_realizada_primary"("id_resposta_realizada");
-- ALTER TABLE "resposta_realizada" ADD UNIQUE "resposta_realizada_id_resposta_unique"("id_resposta");
-- ALTER TABLE "resposta_realizada" ADD UNIQUE "resposta_realizada_id_jogador_unique"("id_jogador");
CREATE TABLE "resposta"(
    "id_resposta" SERIAL PRIMARY KEY,
    "resposta" VARCHAR(255) NOT NULL,,
    "eh_correta" BOOLEAN NOT NULL,
    "id_pergunta" INT NOT NULL
);
-- ALTER table "resposta" ADD PRIMARY KEY "resposta_id_resposta_primary"("id_resposta");
-- ALTER table "resposta" ADD UNIQUE "resposta_id_pergunta_unique"("id_pergunta");
CREATE TABLE "calcula"(
	"id_aux_calcula"SERIAL PRIMARY KEY,
    "id_resposta_realizada" INT not NULL,
    "id_conta_ponto" INT NOT NULL
);
-- ALTER table "calcula" ADD UNIQUE "calcula_id_resposta_realizada_unique"("id_resposta_realizada");
-- ALTER TABLE "calcula" ADD UNIQUE "calcula_id_conta_ponto_unique"("id_conta_ponto");
CREATE TABLE "conta_ponto"(
    "id_conta_ponto" SERIAL PRIMARY KEY,
    "pontos" INT NOT NULL,
    "id_jogador" INT NOT NULL
);
-- ALTER table "conta_ponto" ADD PRIMARY KEY "conta_ponto_id_conta_ponto_primary"("id_conta_ponto");
-- ALTER TABLE "conta_ponto" ADD UNIQUE "conta_ponto_id_jogador_unique"("id_jogador");
CREATE TABLE "tem"(
    "id_aux_tem" SERIAL PRIMARY KEY,
    "id_jogo" INT NOT NULL,
    "id_conta_ponto" INT NOT NULL
);
-- ALTER TABLE "tem" ADD UNIQUE "tem_id_jogo_unique"("id_jogo");
-- ALTER TABLE "tem" ADD UNIQUE "tem_id_conta_ponto_unique"("id_conta_ponto");
CREATE TABLE "pergunta"(
    "id_pergunta" SERIAL PRIMARY KEY,
    "pergunta" VARCHAR(255) NOT NULL,
    "id_jogo" INT NOT NULL
);
-- ALTER table "pergunta" ADD PRIMARY KEY "pergunta_id_pergunta_primary"("id_pergunta");
-- ALTER TABLE "pergunta" ADD UNIQUE "pergunta_id_jogo_unique"("id_jogo");
ALTER TABLE "fase" ADD CONSTRAINT "fase_id_site_foreign" FOREIGN KEY("id_site") REFERENCES "site"("id_site");
ALTER TABLE "jogo" ADD CONSTRAINT "jogo_id_fase_foreign" FOREIGN KEY("id_fase") REFERENCES "fase"("id_fase");
ALTER TABLE "joga" ADD CONSTRAINT "joga_id_jogo_foreign" FOREIGN KEY("id_jogo") REFERENCES "jogo"("id_jogo");
ALTER TABLE "joga" ADD CONSTRAINT "joga_id_jogador_foreign" FOREIGN KEY("id_jogador") REFERENCES "jogador"("id_jogador");
ALTER TABLE "resposta_realizada" ADD CONSTRAINT "resposta_realizada_id_resposta_foreign" FOREIGN KEY("id_resposta") REFERENCES "resposta"("id_resposta");
ALTER TABLE "resposta_realizada" ADD CONSTRAINT "resposta_realizada_id_jogador_foreign" FOREIGN KEY("id_jogador") REFERENCES "jogador"("id_jogador");
ALTER TABLE "resposta" ADD CONSTRAINT "resposta_id_pergunta_foreign" FOREIGN KEY("id_pergunta") REFERENCES "pergunta"("id_pergunta");
ALTER TABLE "calcula" ADD CONSTRAINT "calcula_id_resposta_realizada_foreign" FOREIGN KEY("id_resposta_realizada") REFERENCES "resposta_realizada"("id_resposta_realizada");
ALTER TABLE "calcula" ADD CONSTRAINT "calcula_id_conta_ponto_foreign" FOREIGN KEY("id_conta_ponto") REFERENCES "conta_ponto"("id_conta_ponto");
ALTER TABLE "conta_ponto" ADD CONSTRAINT "conta_ponto_id_jogador_foreign" FOREIGN KEY("id_jogador") REFERENCES "jogador"("id_jogador");
ALTER TABLE "tem" ADD CONSTRAINT "tem_id_jogo_foreign" FOREIGN KEY("id_jogo") REFERENCES "jogo"("id_jogo");
ALTER TABLE "tem" ADD CONSTRAINT "tem_id_conta_ponto_foreign" FOREIGN KEY("id_conta_ponto") REFERENCES "conta_ponto"("id_conta_ponto");
ALTER TABLE "pergunta" ADD CONSTRAINT "pergunta_id_jogo_foreign" FOREIGN KEY("id_jogo") REFERENCES "jogo"("id_jogo");