package models

// Senha representa o formato da requisição dde alteração se senha
type Senha struct {
	Nova  string `json:"nova"`
	Atual string `json:"atual"`
}
