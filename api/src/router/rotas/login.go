package rotas

import (
	"api/src/controllers"
	"net/http"
)

var rotaLogin = Rota{
	URI:    "/api/login",
	Metodo: http.MethodPost,
	Funcao: controllers.Login,
	RequerAutenticacao: false,
}
