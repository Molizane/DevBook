package banco

import (
	"api/src/config"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

// Conectar abre a conexão como banco de dados e a retorna
func Conectar() (*sql.DB, error) {
	db, erro := sql.Open(config.Servidor, config.StringConexaoBanco)

	if erro != nil {
		return nil, erro
	}

	if erro = db.Ping(); erro != nil {
		db.Close()
		return nil, erro
	}

	return db, nil
}
