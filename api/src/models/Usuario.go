package models

import (
	"api/src/seguranca"
	"errors"
	"strings"
	"time"

	"github.com/badoux/checkmail"
)

// Usuario representa um usuário utilkizando a rede social
type Usuario struct {
	ID                   uint64    `json:"id,omitempty"`
	Nome                 string    `json:"nome,omitempty"`
	Nick                 string    `json:"nick,omitempty"`
	Email                string    `json:"email,omitempty"`
	Senha                string    `json:"senha,omitempty"`
	CriadoEm             time.Time `json:"CriadoEm,omitempty"`
	Bloqueado            uint8     `json:"bloqueado,omitempty"`
	BloqueadoPeloSeguido uint8     `json:"bloqueadoPeloSeguido,omitempty"`
}

// Preparar faz ajustes no registro de usuário
func (usuario *Usuario) Preparar(etapa string) error {
	if erro := usuario.validar(etapa); erro != nil {
		return erro
	}

	return usuario.formatar(etapa)
}

func (usuario *Usuario) validar(etapa string) error {
	if usuario.Nome == "" {
		return errors.New("nome do usuário é obrigatório")
	}

	if usuario.Nick == "" {
		return errors.New("nickname do usuário é obrigatório")
	}

	if usuario.Email == "" {
		return errors.New("e-mail do usuário é obrigatório")
	}

	if erro := checkmail.ValidateFormat(usuario.Email); erro != nil {
		return errors.New("e-mail informado é inválido")
	}

	if etapa == "cadastro" && usuario.Senha == "" {
		return errors.New("senha do usuário é obrigatório")
	}

	return nil
}

func (usuario *Usuario) formatar(etapa string) error {
	usuario.Nome = strings.TrimSpace(usuario.Nome)
	usuario.Nick = strings.TrimSpace(usuario.Nick)
	usuario.Email = strings.TrimSpace(usuario.Email)

	senhaComHash, erro := seguranca.Hash(usuario.Senha)

	if erro != nil {
		return erro
	}

	usuario.Senha = string(senhaComHash)

	return nil
}
