package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ouvermax/db_go_games/controllers"
	"github.com/ouvermax/db_go_games/middlewares"
	"github.com/ouvermax/db_go_games/models"
)

func main() {
	r := gin.Default()

	models.ConnectDatabase()
	api := r.Group("/api")
	{

		v1 := api.Group("/v1")
		{
			jogador := v1.Group("/jogador", middlewares.Auth())
			{
				jogador.GET("/:id", controllers.PegaJogador)
				jogador.POST("", controllers.CriaJogador)
				jogador.PUT("/:id", controllers.AtualizaJogador)
				jogador.DELETE("/:id", controllers.RemoverJogador)
			}

			jogo := v1.Group("/jogo", middlewares.Auth())
			{
				jogo.GET("/:id", controllers.PegaJogo)
				jogo.POST("", controllers.CriaJogo)
				jogo.PUT("/:id", controllers.AtualizaJogo)
				jogo.DELETE("/:id", controllers.RemoverJogo)
			}

			site := v1.Group("/site")
			{
				site.GET("/:id", controllers.PegaSite)
				site.POST("", controllers.CriaSite)
				site.PUT("/:id", controllers.AtualizaSite)
				site.DELETE("/:id", controllers.RemoverSite)
			}

			fase := v1.Group("/fase")
			{
				fase.GET("/:id", controllers.PegaFase)
				fase.POST("", controllers.CriaFase)
				fase.PUT("/:id", controllers.AtualizaFase)
				fase.DELETE("/:id", controllers.RemoverFase)
			}

			pergunta := v1.Group("/pergunta")
			{
				pergunta.GET("/:id", controllers.PegaPergunta)
				pergunta.POST("", controllers.CriaPergunta)
				pergunta.PUT("/:id", controllers.AtualizaPergunta)
				pergunta.DELETE("/:id", controllers.RemoverPergunta)
			}

			resposta := v1.Group("/resposta")
			{
				resposta.GET("/:id", controllers.PegaResposta)
				resposta.POST("", controllers.CriaResposta)
				resposta.PUT("/:id", controllers.AtualizaResposta)
				resposta.DELETE("/:id", controllers.RemoverResposta)
			}

			respostarealizada := v1.Group("/respostarealizada")
			{
				respostarealizada.GET("/:id", controllers.PegaRespostaRealizada)
				respostarealizada.POST("", controllers.CriaRespostaRealizada)
				respostarealizada.PUT("/:id", controllers.AtualizaRespostaRealizada)
				respostarealizada.DELETE("/:id", controllers.RemoverRespostaRealizada)
			}

			joga := v1.Group("/joga")
			{
				joga.GET("/:id", controllers.PegaJoga)
				joga.POST("", controllers.CriaJoga)
				joga.PUT("/:id", controllers.AtualizaJoga)
				joga.DELETE("/:id", controllers.RemoverJoga)
			}

			contaponto := v1.Group("/contaponto")
			{
				contaponto.GET("/:id", controllers.PegaContaPonto)
				contaponto.POST("", controllers.CriaContaPonto)
				contaponto.PUT("/:id", controllers.AtualizaContaPonto)
				contaponto.DELETE("/:id", controllers.RemoverContaPonto)
			}

			calcula := v1.Group("/calcula")
			{
				calcula.GET("/:id", controllers.PegaCalcula)
				calcula.POST("", controllers.CriaCalcula)
				calcula.PUT("/:id", controllers.AtualizaCalcula)
				calcula.DELETE("/:id", controllers.RemoverCalcula)
			}

			tem := v1.Group("/tem")
			{
				tem.GET("/:id", controllers.PegaTem)
				tem.POST("", controllers.CriaTem)
				tem.PUT("/:id", controllers.AtualizaTem)
				tem.DELETE("/:id", controllers.RemoverTem)
			}

			user := v1.Group("/users", middlewares.Auth())
			{
				user.GET("/:id", controllers.PegaUser)
				user.POST("", controllers.CriaUser)
				user.PUT("/:id", controllers.AtualizaUser)
				user.DELETE("/:id", controllers.RemoverUser)
			}

			login := v1.Group("/login")
			{
				login.POST("", controllers.Login)
			}
		}
	}

	r.Run() // listen and serve on 0.0.0.0:8080
}
