package modelos

import (
	"errors"
	"strings"
	"time"
)

type Usuario struct {
	ID       uint64    `json: "id,omitempty"`
	Nome     string    `json: "nome,omitempty`
	Nick     string    `json: "nick,omitempty`
	Email    string    `json: "email,omitempty"`
	Senha    string    `json: "senha,omitempty"`
	CriadoEm time.Time `json: "CriadoEm,omitempty"`
}

// Chama os métodos para validar e formatar o usuário recebido
func (usuario *Usuario) Preparar(etapa string) error {
	if erro := usuario.validar(etapa); erro != nil {
		return erro
	}

	if erro := usuario.formatar(etapa); erro != nil {
		return erro
	}

	return nil
}

func (usuario *Usuario) validar(etapa string) error {
	if usuario.Nome == "" {
		return errors.New("O nome é obrigatorio e não pode estar em branco")
	}

	if usuario.Nick == "" {
		return errors.New("O nick é obrigatorio e não pode estar em branco")
	}

	if usuario.Email == "" {
		return errors.New("O e-mail é obrigatorio e não pode estar em branco")
	}

	if etapa == "cadastro" && usuario.Senha == "" {
		return errors.New("A senha é obrigatorio e não pode estar em branco")
	}

	return nil
}

func (usuario *Usuario) formatar(etapa string) error {
	usuario.Nome = strings.TrimSpace(usuario.Nome)
	usuario.Nome = strings.TrimSpace(usuario.Nick)
	usuario.Nome = strings.TrimSpace(usuario.Email)

	return nil
}
