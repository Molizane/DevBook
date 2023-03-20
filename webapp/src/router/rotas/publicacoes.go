package rotas

import (
	"net/http"
	"webapp/src/controllers"
)

var rotasPublicadcoes = []Rota{
	{
		URI: "/publicacoes",
		Metodo: http.MethodPost,
		Funcao: controllers.CriarPublicacao,
		RequerAutenticacao: true,
	},
	{
		URI: "/publicacoes/{publicacaoId}/alternar-curtir",
		Metodo: http.MethodPost,
		Funcao: controllers.AlternarCurtirPublicacao,
		RequerAutenticacao: true,
	},
	{
		URI: "/publicacoes/{publicacaoId}/alternar-descurtir",
		Metodo: http.MethodPost,
		Funcao: controllers.AlternarDescurtirPublicacao,
		RequerAutenticacao: true,
	},
	{
		URI: "/publicacoes/{publicacaoId}/atualizar",
		Metodo: http.MethodGet,
		Funcao: controllers.CarregarPaginaDeAtualizacaoDePublicacao,
		RequerAutenticacao: true,
	},
	{
		URI: "/publicacoes/{publicacaoId}",
		Metodo: http.MethodPut,
		Funcao: controllers.AtualizarPublicacao,
		RequerAutenticacao: true,
	},
	{
		URI: "/publicacoes/{publicacaoId}",
		Metodo: http.MethodDelete,
		Funcao: controllers.ApagarPublicacao,
		RequerAutenticacao: true,
	},
}