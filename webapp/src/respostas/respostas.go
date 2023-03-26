package respostas

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"webapp/src/cookies"
)

// Erro representa a resposta de erro da API
type ErroAPI struct {
	Erro string `json:"erro"`
}

// JSON retorna uma resposta em formato JSON para a requisição
func JSON(w http.ResponseWriter, statusCode int, dados interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if statusCode != http.StatusNoContent {
		if erro := json.NewEncoder(w).Encode(dados); erro != nil {
			log.Fatal("ERRO", erro)
		}
	}
}

// TratarStatusCodeDeErro trata as requisições com status code 400 u superior
func TratarStatusCodeDeErro(w http.ResponseWriter, res *http.Response, req *http.Request) {
	var erro ErroAPI
	json.NewDecoder(res.Body).Decode(&erro)

	if strings.Contains(erro.Erro, "expired") {
		cookies.Deletar(w)
		http.Redirect(w, req, "/login", http.StatusFound)
		return
	}

	JSON(w, res.StatusCode, erro)
}
