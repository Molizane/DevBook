package rotas

import (
	"api/src/controllers"
	"net/http"
)

var rotasUsuarios = []Rota{
	{
		URI:                "/api/usuarios",
		Metodo:             http.MethodPost,
		Funcao:             controllers.CriarUsuario,
		RequerAutenticacao: false,
	},
	{
		URI:                "/api/usuarios",
		Metodo:             http.MethodGet,
		Funcao:             controllers.BuscarUsuarios,
		RequerAutenticacao: true,
	},
	{
		URI:                "/api/usuarios/{usuarioId}",
		Metodo:             http.MethodGet,
		Funcao:             controllers.BuscarUsuarioPorId,
		RequerAutenticacao: true,
	},
	{
		URI:                "/api/usuarios/{usuarioId}",
		Metodo:             http.MethodPut,
		Funcao:             controllers.AtualizarUsuario,
		RequerAutenticacao: true,
	},
	{
		URI:                "/api/usuarios/{usuarioId}",
		Metodo:             http.MethodDelete,
		Funcao:             controllers.DeletarUsuario,
		RequerAutenticacao: true,
	},
	{
		URI:                "/api/usuarios/{usuarioId}/seguir",
		Metodo:             http.MethodPost,
		Funcao:             controllers.SeguirUsuario,
		RequerAutenticacao: true,
	},
	{
		URI:                "/api/usuarios/{usuarioId}/parar-de-seguir",
		Metodo:             http.MethodPost,
		Funcao:             controllers.ParaDeSeguirUsuario,
		RequerAutenticacao: true,
	},
	{
		URI:                "/api/usuarios/{usuarioId}/seguidores",
		Metodo:             http.MethodGet,
		Funcao:             controllers.BuscarSeguidores,
		RequerAutenticacao: true,
	},
	{
		URI:                "/api/usuarios/{usuarioId}/seguindo",
		Metodo:             http.MethodGet,
		Funcao:             controllers.BuscarSeguindo,
		RequerAutenticacao: true,
	},
	{
		URI:                "/api/usuarios/{usuarioId}/atualizar-senha",
		Metodo:             http.MethodPost,
		Funcao:             controllers.AtualizarSenha,
		RequerAutenticacao: true,
	},
	{
		URI:                "/api/usuarios/{usuarioId}/bloquear",
		Metodo:             http.MethodPost,
		Funcao:             controllers.BloquearSeguidor,
		RequerAutenticacao: true,
	},
	{
		URI:                "/api/usuarios/{usuarioId}/desbloquear",
		Metodo:             http.MethodPost,
		Funcao:             controllers.DesbloquearSeguidor,
		RequerAutenticacao: true,
	},
}
