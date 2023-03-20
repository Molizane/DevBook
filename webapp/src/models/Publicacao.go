package models

import "time"

// Publicacao bla bla bla
type Publicacao struct {
	ID          uint64    `json:"id,omitempty"`
	Titulo      string    `json:"titulo,omitempty"`
	Conteudo    string    `json:"conteudo,omitempty"`
	AutorID     uint64    `json:"autorId,omitempty"`
	AutorNick   string    `json:"autorNick,omitempty"`
	Curtidas    uint64    `json:"curtidas"`
	Descurtidas uint64    `json:"descurtidas"`
	Curtiu      uint8     `json:"curtiu"`
	Descurtiu   uint8     `json:"descurtiu"`
	CriadaEm    time.Time `json:"criadaEm,omitempty"`
}
