package banco

import (
	"api/src/config"
	"database/sql"
)

func Concetar() (*sql.DB, error) {
	db, erro := sql.Open("mysql", config.StringConecaoBanco)

	if erro != nil {
		return nil, erro
	}

	if erro = db.Ping(); erro != nil {
		db.Close()
		return nil, erro
	}

	return db, nil
}
