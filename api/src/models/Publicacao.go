package models

import (
	"errors"
	"strings"
	"time"
)

// Publicacao representa uma publicação feita por um usuário
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

// Preparar faz ajustes no registro de publicação
func (publicacao *Publicacao) Preparar() error {
	publicacao.formatar()

	if erro := publicacao.validar(); erro != nil {
		return erro
	}

	return nil
}

func (publicacao *Publicacao) validar() error {
	if publicacao.Titulo == "" {
		return errors.New("o título é obrigatório")
	}

	if publicacao.Conteudo == "" {
		return errors.New("o conteúdo é obrigatório")
	}

	return nil
}

func (publicacao *Publicacao) formatar() {
	publicacao.Titulo = strings.TrimSpace(publicacao.Titulo)
	publicacao.Conteudo = strings.TrimSpace(publicacao.Conteudo)
}
