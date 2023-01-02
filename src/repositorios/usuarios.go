package repositorios

import (
	"api/src/modelos"
	"database/sql"
)

// Representa um repositório de usuarios
type Usuarios struct {
	db *sql.DB
}

// Cria um repositório de usuarios
func NovoRepositorioDeUsuarios(db *sql.DB) *Usuarios {
	return &Usuarios{db}
}

// Insere um usuário no banco de dados
func (u Usuarios) Criar(usuario modelos.Usuario) (uint64, error) {
	return 0, nil
}
