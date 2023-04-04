package middlewares

import (
	"api/src/autenticacao"
	"api/src/respostas"
	"log"
	"net/http"
)

// Looger faz o log da requisição
func Logger(proximaFuncao http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf(": %s http://%s%s", r.Method, r.Host, r.RequestURI)
		proximaFuncao(w, r)
	}
}

// Autenticar verifica se o usuário faznedo a requisição está autenticado
func Autenticar(proximaFuncao http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if erro := autenticacao.ValidarToken(r); erro != nil {
			respostas.Erro(w, http.StatusUnauthorized, erro)
			return
		}

		proximaFuncao(w, r)
	}
}
