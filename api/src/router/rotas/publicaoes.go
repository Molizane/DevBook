package rotas

import (
	"api/src/controllers"
	"net/http"
)

var rotasPublicacoess = []Rota{
	{
		URI:                "/api/publicacoes",
		Metodo:             http.MethodPost,
		Funcao:             controllers.CriarPublicacao,
		RequerAutenticacao: false,
	},
	{
		URI:                "/api/publicacoes",
		Metodo:             http.MethodGet,
		Funcao:             controllers.BuscarPublicacoes,
		RequerAutenticacao: true,
	},
	{
		URI:                "/api/publicacoes/{publicacaoId}",
		Metodo:             http.MethodGet,
		Funcao:             controllers.BuscarPublicacao,
		RequerAutenticacao: true,
	},
	{
		URI:                "/api/publicacoes/{publicacaoId}",
		Metodo:             http.MethodPut,
		Funcao:             controllers.AtualizarPublicacao,
		RequerAutenticacao: true,
	},
	{
		URI:                "/api/publicacoes/{publicacaoId}",
		Metodo:             http.MethodDelete,
		Funcao:             controllers.DeletarPublicacao,
		RequerAutenticacao: true,
	},
	{
		URI:                "/api/usuarios/{usuarioId}/publicacoes",
		Metodo:             http.MethodGet,
		Funcao:             controllers.BuscarPublicacoesPorUsuario,
		RequerAutenticacao: true,
	},
	{
		URI:                "/api/publicacoes/{publicacaoId}/curtir",
		Metodo:             http.MethodPost,
		Funcao:             controllers.CurtirPublicacao,
		RequerAutenticacao: true,
	},
	{
		URI:                "/api/publicacoes/{publicacaoId}/desfazer-curtir",
		Metodo:             http.MethodPost,
		Funcao:             controllers.DesfazerCurtirPublicacao,
		RequerAutenticacao: true,
	},
	{
		URI:                "/api/publicacoes/{publicacaoId}/descurtir",
		Metodo:             http.MethodPost,
		Funcao:             controllers.DescurtirPublicacao,
		RequerAutenticacao: true,
	},
	{
		URI:                "/api/publicacoes/{publicacaoId}/desfazer-descurtir",
		Metodo:             http.MethodPost,
		Funcao:             controllers.DesfazerDescurtirPublicacao,
		RequerAutenticacao: true,
	},
	{
		URI:                "/api/publicacoes/{publicacaoId}/alternar-curtir",
		Metodo:             http.MethodPost,
		Funcao:             controllers.AlternarCurtirPublicacao,
		RequerAutenticacao: true,
	},
	{
		URI:                "/api/publicacoes/{publicacaoId}/alternar-descurtir",
		Metodo:             http.MethodPost,
		Funcao:             controllers.AlternarDescurtirPublicacao,
		RequerAutenticacao: true,
	},
}
