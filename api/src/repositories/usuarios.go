package repositories

import (
	"api/src/models"
	"database/sql"
	"fmt"
)

type usuarios struct {
	db *sql.DB
}

// NovoRepositorioDeUsuarios cria um repositório de usuários
func NovoRepositorioDeUsuarios(db *sql.DB) *usuarios {
	return &usuarios{db}
}

// Criar insere um usuário no banco de dados
func (repositorio usuarios) Criar(usuario models.Usuario) (uint64, error) {
	statement, erro := repositorio.db.Prepare(
		"INSERT INTO usuarios (nome, nick, email, senha) VALUES (?, ?, ?, ?)",
	)

	if erro != nil {
		return 0, erro
	}

	defer statement.Close()

	resultado, erro := statement.Exec(usuario.Nome, usuario.Nick, usuario.Email, usuario.Senha)

	if erro != nil {
		return 0, erro
	}

	ultimoIDInserido, erro := resultado.LastInsertId()

	if erro != nil {
		return 0, erro
	}

	return uint64(ultimoIDInserido), nil
}

// Buscar traz todos os usuários que atendem a um filtro de nome ou nick
func (repositorio usuarios) Buscar(nomeOuNick string) ([]models.Usuario, error) {
	nomeOuNick = fmt.Sprintf("%%%s%%", nomeOuNick)

	linhas, erro := repositorio.db.Query(`
	    SELECT id, nome, nick, email, criadoEm
		FROM usuarios
		WHERE nome LIKE ?
		   OR nick LIKE ?
		ORDER BY nome`,
		nomeOuNick, nomeOuNick,
	)

	if erro != nil {
		return nil, erro
	}

	defer linhas.Close()

	var usuarios []models.Usuario

	for linhas.Next() {
		var usuario models.Usuario

		if erro = linhas.Scan(
			&usuario.ID,
			&usuario.Nome,
			&usuario.Nick,
			&usuario.Email,
			&usuario.CriadoEm,
		); erro != nil {
			return nil, erro
		}

		usuarios = append(usuarios, usuario)
	}

	return usuarios, nil
}

// Buscar traz todos os usuários que atendem a um filtro de nome ou nick
func (repositorio usuarios) BuscarPorID(ID uint64) (models.Usuario, error) {
	linhas, erro := repositorio.db.Query(`
	    SELECT id, nome, nick, email, criadoEm
		FROM usuarios
		WHERE id = ?`,
		ID,
	)

	if erro != nil {
		return models.Usuario{}, erro
	}

	defer linhas.Close()

	var usuario models.Usuario

	if linhas.Next() {
		if erro = linhas.Scan(
			&usuario.ID,
			&usuario.Nome,
			&usuario.Nick,
			&usuario.Email,
			&usuario.CriadoEm,
		); erro != nil {
			return models.Usuario{}, erro
		}
	}

	return usuario, nil
}

// Atualizar altera as informações de um usuário no banco de dados
func (repositorio usuarios) Atualizar(ID uint64, usuario models.Usuario) error {
	statement, erro := repositorio.db.Prepare(`
	  UPDATE usuarios
	  SET nome = ?, nick = ?, email = ?
	  WHERE id = ?`)

	if erro != nil {
		return erro
	}

	defer statement.Close()

	_, erro = statement.Exec(usuario.Nome, usuario.Nick, usuario.Email, ID)

	return erro
}

// Deletar apaga um usuário no banco de dados
func (repositorio usuarios) Deletar(ID uint64) error {
	statement, erro := repositorio.db.Prepare("DELETE FROM usuarios WHERE id = ?")

	if erro != nil {
		return erro
	}

	defer statement.Close()

	_, erro = statement.Exec(ID)

	return erro
}

// BuscarPorEmail busca um usuário por email e retorna o seu id e senha com senha
func (repositorio usuarios) BuscarPorEmail(email string) (models.Usuario, error) {
	linhas, erro := repositorio.db.Query(
		"SELECT id, senha FROM usuarios WHERE email = ?", email)

	if erro != nil {
		return models.Usuario{}, erro
	}

	defer linhas.Close()

	var usuario models.Usuario

	if linhas.Next() {
		if erro = linhas.Scan(&usuario.ID, &usuario.Senha); erro != nil {
			return models.Usuario{}, erro
		}
	}

	return usuario, nil
}

// Seguir permite que um usuário siga outro
func (repositorio usuarios) Seguir(usuarioID, seguidorID uint64) error {
	statement, erro := repositorio.db.Prepare(`
	  INSERT IGNORE INTO seguidores (usuario_id, seguidor_id)
	  VALUES (?, ?)`)

	if erro != nil {
		return erro
	}

	defer statement.Close()

	_, erro = statement.Exec(usuarioID, seguidorID)

	return erro
}

// PararDeSeguir permite que um usuário deixe de siguir outro
func (repositorio usuarios) PararDeSeguir(usuarioID, seguidorID uint64) error {
	statement, erro := repositorio.db.Prepare(`
	  DELETE FROM seguidores
	  WHERE usuario_id = ?
	  AND seguidor_id = ?
	  AND bloqueado = 0`)

	if erro != nil {
		return erro
	}

	defer statement.Close()

	_, erro = statement.Exec(usuarioID, seguidorID)

	return erro
}

// BuscarSeguidores traz todos os seguidores do usuário
func (repositorio usuarios) BuscarSeguidores(usuarioID uint64) ([]models.Usuario, error) {
	linhas, erro := repositorio.db.Query(`
		 SELECT DISTINCT u.id, u.nome, u.nick, u.email, u.criadoEm, s.bloqueado
		 FROM usuarios u
		 INNER JOIN seguidores s
		 ON s.seguidor_id = u.id
		 WHERE s.usuario_id = ?`,
		usuarioID,
	)

	if erro != nil {
		return nil, erro
	}

	defer linhas.Close()

	var usuarios []models.Usuario

	for linhas.Next() {
		var usuario models.Usuario

		if erro = linhas.Scan(
			&usuario.ID,
			&usuario.Nome,
			&usuario.Nick,
			&usuario.Email,
			&usuario.CriadoEm,
			&usuario.Bloqueado,
		); erro != nil {
			return nil, erro
		}

		usuarios = append(usuarios, usuario)
	}

	return usuarios, nil
}

// BuscarSeguindo traz todos os usuários seguidos por um determinado usuário
func (repositorio usuarios) BuscarSeguindo(usuarioID uint64) ([]models.Usuario, error) {
	linhas, erro := repositorio.db.Query(`
		 SELECT DISTINCT u.id, u.nome, u.nick, u.email, u.criadoEm
		 FROM usuarios u
		 INNER JOIN seguidores s
		 ON s.usuario_id = u.id
		 WHERE s.seguidor_id = ?
		 AND s.bloqueado = 0`,
		usuarioID,
	)

	if erro != nil {
		return nil, erro
	}

	defer linhas.Close()

	var usuarios []models.Usuario

	for linhas.Next() {
		var usuario models.Usuario

		if erro = linhas.Scan(
			&usuario.ID,
			&usuario.Nome,
			&usuario.Nick,
			&usuario.Email,
			&usuario.CriadoEm,
		); erro != nil {
			return nil, erro
		}

		usuarios = append(usuarios, usuario)
	}

	return usuarios, nil
}

// BuscarSenha retorna a senha de um usuário pelo ID
func (repositorio usuarios) BuscarSenha(usuarioID uint64) (string, error) {
	linha, erro := repositorio.db.Query("SELECT senha FROM usuarios WHERE id = ?", usuarioID)

	if erro != nil {
		return "", erro
	}

	defer linha.Close()

	var senha string

	if linha.Next() {
		if erro = linha.Scan(&senha); erro != nil {
			return "", erro
		}
	}

	return senha, nil
}

// AtualizarSenha atualiza a senha de um usuário pelo ID
func (repositorio usuarios) AtualizarSenha(usuarioID uint64, senha string) error {
	statement, erro := repositorio.db.Prepare("UPDATE usuarios SET senha = ? WHERE id = ?")

	if erro != nil {
		return erro
	}

	defer statement.Close()

	_, erro = statement.Exec(senha, usuarioID)

	return erro
}

func (repositorio usuarios) Bloquear(usuarioID, seguidorID uint64) error {
	statement, erro := repositorio.db.Prepare(
		`UPDATE seguidores
		 SET bloqueado = CASE bloqueado WHEN 1 THEN 0 ELSE 1 END
		 WHERE usuario_id = ?
		 AND seguidor_id = ?`,
	)

	if erro != nil {
		return erro
	}

	defer statement.Close()

	_, erro = statement.Exec(usuarioID, seguidorID)

	return erro
}
