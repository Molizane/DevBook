package repositories

import (
	"api/src/models"
	"database/sql"
)

// Publicacoes repressenta um repositório de publicações
type Publicacoes struct {
	db *sql.DB
}

// NovoRepositorioDePublicacoes cria um repositório de publicações
func NovoRepositorioDePublicacoes(db *sql.DB) *Publicacoes {
	return &Publicacoes{db}
}

// Criar insere uma publicação no banco de dados
func (repositorio Publicacoes) Criar(publicacao models.Publicacao) (uint64, error) {
	statement, erro := repositorio.db.Prepare(
		"INSERT INTO publicacoes (titulo, conteudo, autor_id) VALUES (?, ?, ?)",
	)

	if erro != nil {
		return 0, erro
	}

	defer statement.Close()

	resultado, erro := statement.Exec(publicacao.Titulo, publicacao.Conteudo, publicacao.AutorID)

	if erro != nil {
		return 0, erro
	}

	ultimoIDInserido, erro := resultado.LastInsertId()

	if erro != nil {
		return 0, erro
	}

	return uint64(ultimoIDInserido), nil
}

// BuscarPorID traz uma única publicação do banco de dados pelo seu ID
func (repositorio Publicacoes) BuscarPorID(publicacaoID uint64) (models.Publicacao, error) {
	linha, erro := repositorio.db.Query(`
      SELECT p.id, p.titulo, p.conteudo, p.autor_id,
             (SELECT COUNT(*)
              FROM curtidas
              WHERE publicacao_id = p.id) AS curtidas,
             (SELECT COUNT(*)
              FROM descurtidas
              WHERE publicacao_id = p.id) AS descurtidas,
             p.criadaEm, u.nick
      FROM publicacoes p
      INNER JOIN usuarios u
      ON u.id = p.autor_id
      WHERE p.id = ?
     `, publicacaoID)

	if erro != nil {
		return models.Publicacao{}, erro
	}

	defer linha.Close()

	var publicacao models.Publicacao

	if linha.Next() {
		if erro = linha.Scan(
			&publicacao.ID,
			&publicacao.Titulo,
			&publicacao.Conteudo,
			&publicacao.AutorID,
			&publicacao.Curtidas,
			&publicacao.Descurtidas,
			&publicacao.CriadaEm,
			&publicacao.AutorNick,
		); erro != nil {
			return models.Publicacao{}, erro
		}
	}

	return publicacao, nil
}

// Buscar traz AS publicações dos usuários seguidoes e também do próprio usuário que fez a requisição
func (repositorio Publicacoes) Buscar(usuarioID uint64) ([]models.Publicacao, error) {
	linhas, erro := repositorio.db.Query(`
	SELECT DISTINCT p.id, p.titulo, p.conteudo, p.autor_id,
           (SELECT COUNT(*)
            FROM curtidas
            WHERE publicacao_id = p.id) AS curtidas,
           (SELECT COUNT(*)
            FROM descurtidas
            WHERE publicacao_id = p.id) AS descurtidas,
           (SELECT COUNT(*)
            FROM curtidas
            WHERE publicacao_id = p.id
			AND usuario_id = ?) AS curtiu,
           (SELECT COUNT(*)
            FROM descurtidas
            WHERE publicacao_id = p.id
			AND usuario_id = ?) AS descurtiu,
           p.criadaEm, u.nick
    FROM publicacoes p
    INNER JOIN usuarios u
      ON u.id = p.autor_id
    LEFT OUTER JOIN seguidores s
      ON p.autor_id = s.usuario_id
    WHERE u.id = ?
       OR s.seguidor_id = ?
    ORDER BY 1 DESC`, usuarioID, usuarioID, usuarioID, usuarioID)

	if erro != nil {
		return nil, erro
	}

	defer linhas.Close()

	var publicacoes []models.Publicacao

	for linhas.Next() {
		var publicacao models.Publicacao

		if erro = linhas.Scan(
			&publicacao.ID,
			&publicacao.Titulo,
			&publicacao.Conteudo,
			&publicacao.AutorID,
			&publicacao.Curtidas,
			&publicacao.Descurtidas,
			&publicacao.Curtiu,
			&publicacao.Descurtiu,
			&publicacao.CriadaEm,
			&publicacao.AutorNick,
		); erro != nil {
			return nil, erro
		}

		publicacoes = append(publicacoes, publicacao)
	}

	return publicacoes, nil
}

// Atualizar altera os dados de uma publicação no banco de dados
func (repositorio Publicacoes) Atualizar(publicacaoID uint64, publicacao models.Publicacao) error {
	statement, erro := repositorio.db.Prepare("UPDATE publicacoes SET titulo = ?, conteudo = ? WHERE id = ?")

	if erro != nil {
		return erro
	}

	defer statement.Close()

	_, erro = statement.Exec(publicacao.Titulo, publicacao.Conteudo, publicacaoID)
	return erro
}

// Deletar exclui os dados de uma publicação no banco de dados
func (repositorio Publicacoes) Deletar(publicacaoID uint64) error {
	statement, erro := repositorio.db.Prepare("DELETE FROM publicacoes WHERE id = ?")

	if erro != nil {
		return erro
	}

	defer statement.Close()

	_, erro = statement.Exec(publicacaoID)
	return erro
}

func (repositorio Publicacoes) BuscarPorUsuario(usuarioID uint64) ([]models.Publicacao, error) {
	linhas, erro := repositorio.db.Query(`
      SELECT p.id, p.titulo, p.conteudo, p.autor_id,
             (SELECT COUNT(*)
              FROM curtidas
              WHERE publicacao_id = p.id) AS curtidas,
             (SELECT COUNT(*)
              FROM descurtidas
              WHERE publicacao_id = p.id) AS descurtidas,
             p.criadaEm, u.nick
      FROM publicacoes p
      INNER JOIN usuarios u
      ON u.id = p.autor_id
      WHERE p.autor_id = ?
	  ORDER BY 1 DESC
     `, usuarioID)

	if erro != nil {
		return nil, erro
	}

	defer linhas.Close()

	var publicacoes []models.Publicacao

	for linhas.Next() {
		var publicacao models.Publicacao

		if erro = linhas.Scan(
			&publicacao.ID,
			&publicacao.Titulo,
			&publicacao.Conteudo,
			&publicacao.AutorID,
			&publicacao.Curtidas,
			&publicacao.Descurtidas,
			&publicacao.CriadaEm,
			&publicacao.AutorNick,
		); erro != nil {
			return nil, erro
		}

		publicacoes = append(publicacoes, publicacao)
	}

	return publicacoes, nil
}

func (repositorio Publicacoes) buscarEngajamento(publicacaoID, usuarioID uint64) (models.Publicacao, error) {
	linhas, erro := repositorio.db.Query(`
	SELECT p.id,
           (SELECT COUNT(*)
            FROM curtidas
            WHERE publicacao_id = p.id) AS curtidas,
           (SELECT COUNT(*)
            FROM descurtidas
            WHERE publicacao_id = p.id) AS descurtidas,
           (SELECT COUNT(*)
            FROM curtidas
            WHERE publicacao_id = p.id
			AND usuario_id = ?) AS curtiu,
           (SELECT COUNT(*)
            FROM descurtidas
            WHERE publicacao_id = p.id
			AND usuario_id = ?) AS descurtiu
    FROM publicacoes p
    WHERE p.id = ?`, usuarioID, usuarioID, publicacaoID)

	if erro != nil {
		return models.Publicacao{}, erro
	}

	defer linhas.Close()

	var publicacao models.Publicacao

	if linhas.Next() {
		if erro = linhas.Scan(
			&publicacao.ID,
			&publicacao.Curtidas,
			&publicacao.Descurtidas,
			&publicacao.Curtiu,
			&publicacao.Descurtiu,
		); erro != nil {
			return models.Publicacao{}, erro
		}
	} else {
		publicacao = models.Publicacao{}
	}

	return publicacao, nil
}

// Curtir adiciona uma curtida na publicação
func (repositorio Publicacoes) Curtir(publicacaoID, usuarioID uint64) (models.Publicacao, error) {
	apagarOposto, erro := repositorio.db.Prepare("DELETE FROM descurtidas WHERE publicacao_id = ? AND usuario_id = ?")

	if erro != nil {
		return models.Publicacao{}, erro
	}

	defer apagarOposto.Close()

	_, erro = apagarOposto.Exec(publicacaoID, usuarioID)

	if erro != nil {
		return models.Publicacao{}, erro
	}

	statement, erro := repositorio.db.Prepare("INSERT IGNORE INTO curtidas (publicacao_id, usuario_id) VALUES (?, ?)")

	if erro != nil {
		return models.Publicacao{}, erro
	}

	defer statement.Close()

	_, erro = statement.Exec(publicacaoID, usuarioID)

	if erro != nil {
		return models.Publicacao{}, erro
	}

	publicacao, erro := repositorio.buscarEngajamento(publicacaoID, usuarioID)

	return publicacao, erro
}

// Descurtir retira uma curtida na publicação
func (repositorio Publicacoes) DesfazerCurtir(publicacaoID, usuarioID uint64) (models.Publicacao, error) {
	statement, erro := repositorio.db.Prepare("DELETE FROM curtidas WHERE publicacao_id = ? AND usuario_id = ?")

	if erro != nil {
		return models.Publicacao{}, erro
	}

	defer statement.Close()

	_, erro = statement.Exec(publicacaoID, usuarioID)

	if erro != nil {
		return models.Publicacao{}, erro
	}

	publicacao, erro := repositorio.buscarEngajamento(publicacaoID, usuarioID)

	return publicacao, erro
}

// Curtir adiciona uma curtida na publicação
func (repositorio Publicacoes) Descurtir(publicacaoID, usuarioID uint64) (models.Publicacao, error) {
	apagarOposto, erro := repositorio.db.Prepare("DELETE FROM curtidas WHERE publicacao_id = ? AND usuario_id = ?")

	if erro != nil {
		return models.Publicacao{}, erro
	}

	defer apagarOposto.Close()

	_, erro = apagarOposto.Exec(publicacaoID, usuarioID)

	if erro != nil {
		return models.Publicacao{}, erro
	}

	statement, erro := repositorio.db.Prepare("INSERT IGNORE INTO descurtidas (publicacao_id, usuario_id) VALUES (?, ?)")

	if erro != nil {
		return models.Publicacao{}, erro
	}

	defer statement.Close()

	_, erro = statement.Exec(publicacaoID, usuarioID)

	if erro != nil {
		return models.Publicacao{}, erro
	}

	publicacao, erro := repositorio.buscarEngajamento(publicacaoID, usuarioID)

	return publicacao, erro
}

// Descurtir retira uma curtida na publicação
func (repositorio Publicacoes) DesfazerDescurtir(publicacaoID, usuarioID uint64) (models.Publicacao, error) {
	statement, erro := repositorio.db.Prepare("DELETE FROM descurtidas WHERE publicacao_id = ? AND usuario_id = ?")

	if erro != nil {
		return models.Publicacao{}, erro
	}

	defer statement.Close()

	_, erro = statement.Exec(publicacaoID, usuarioID)

	if erro != nil {
		return models.Publicacao{}, erro
	}

	publicacao, erro := repositorio.buscarEngajamento(publicacaoID, usuarioID)

	return publicacao, erro
}

// AlternarCurtir alterna os curtidas entre curtiu e desfez na publicação
func (repositorio Publicacoes) AlternarCurtir(publicacaoID, usuarioID uint64) (models.Publicacao, error) {
	linha, erro := repositorio.db.Query(`
      SELECT COUNT(*)
      FROM curtidas
      WHERE publicacao_id = ?
      AND usuario_id = ?
     `, publicacaoID, usuarioID)

	if erro != nil {
		return models.Publicacao{}, erro
	}

	defer linha.Close()

	var qt uint64 = 0

	if linha.Next() {
		if erro = linha.Scan(
			&qt,
		); erro != nil {
			return models.Publicacao{}, erro
		}
	}

	if qt == 0 {
		apagaOposto, erro := repositorio.db.Prepare("DELETE FROM descurtidas WHERE publicacao_id = ? AND usuario_id = ?")

		if erro != nil {
			return models.Publicacao{}, erro
		}

		defer apagaOposto.Close()

		_, erro = apagaOposto.Exec(publicacaoID, usuarioID)

		if erro != nil {
			return models.Publicacao{}, erro
		}
	}

	var statement *sql.Stmt

	if qt == 0 {
		statement, erro = repositorio.db.Prepare("INSERT INTO curtidas (publicacao_id, usuario_id) VALUES (?, ?)")
	} else {
		statement, erro = repositorio.db.Prepare("DELETE FROM curtidas WHERE publicacao_id = ? AND usuario_id = ?")
	}

	if erro != nil {
		return models.Publicacao{}, erro
	}

	defer statement.Close()

	_, erro = statement.Exec(publicacaoID, usuarioID)

	if erro != nil {
		return models.Publicacao{}, erro
	}

	publicacao, erro := repositorio.buscarEngajamento(publicacaoID, usuarioID)

	return publicacao, erro
}

// AlternarDescurtir alterna as descurtidas entre descurtiu e desfez na publicação
func (repositorio Publicacoes) AlternarDescurtir(publicacaoID, usuarioID uint64) (models.Publicacao, error) {
	linha, erro := repositorio.db.Query(`
      SELECT COUNT(*)
      FROM descurtidas
      WHERE publicacao_id = ?
      AND usuario_id = ?
     `, publicacaoID, usuarioID)

	if erro != nil {
		return models.Publicacao{}, erro
	}

	defer linha.Close()

	var qt uint64 = 0

	if linha.Next() {
		if erro = linha.Scan(
			&qt,
		); erro != nil {
			return models.Publicacao{}, erro
		}
	}

	if qt == 0 {
		apagaOposto, erro := repositorio.db.Prepare("DELETE FROM curtidas WHERE publicacao_id = ? AND usuario_id = ?")

		if erro != nil {
			return models.Publicacao{}, erro
		}

		defer apagaOposto.Close()

		_, erro = apagaOposto.Exec(publicacaoID, usuarioID)

		if erro != nil {
			return models.Publicacao{}, erro
		}
	}

	var statement *sql.Stmt

	if qt == 0 {
		statement, erro = repositorio.db.Prepare("INSERT INTO descurtidas (publicacao_id, usuario_id) VALUES (?, ?)")
	} else {
		statement, erro = repositorio.db.Prepare("DELETE FROM descurtidas WHERE publicacao_id = ? AND usuario_id = ?")
	}

	if erro != nil {
		return models.Publicacao{}, erro
	}

	defer statement.Close()

	_, erro = statement.Exec(publicacaoID, usuarioID)

	if erro != nil {
		return models.Publicacao{}, erro
	}

	publicacao, erro := repositorio.buscarEngajamento(publicacaoID, usuarioID)

	return publicacao, erro
}
