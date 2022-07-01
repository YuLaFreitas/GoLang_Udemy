package repositorios

import (
	"api/src/modelos"
	"database/sql"
)

type usuarios struct {
	db *sql.DB
}

func NovoReposotioDeUsuarios(db *sql.DB) *usuarios {
	return &usuarios{db}
}

func (u usuarios) CriarInBD(usuario modelos.Usuario) (uint64, error) {
	return 0, nil
}
