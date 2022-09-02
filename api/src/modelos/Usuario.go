package modelos

import (
	"errors"
	"strings"
	"time"
)

// Representa um usuário usando rede social
type Usuario struct {
	ID       uint64    `json:"id,omitempty"`
	Nome     string    `json:"nome,omitempty"`
	Nick     string    `json:"nick,omitempty"`
	Email    string    `json:"email,omitempty"`
	Senha    string    `json:"senha,omitempty"`
	CriadoEm time.Time `json:"criadoEm,omitempty"`
}

// Preparar valida os campos de usuário
func (usuario *Usuario) Preparar() error {
	usuario.formatar()
	if erro := usuario.validar(); erro != nil {
		return erro
	}
	return nil
}

// Validar valida os campos do usuário
func (usuario *Usuario) validar() error {
	if usuario.Nome == "" {
		return errors.New("O nome é obrigatório e não pode estar em branco")
	}

	if usuario.Email == "" {
		return errors.New("O e-mail é obrigatório e não pode estar em branco")
	}

	if usuario.Senha == "" {
		return errors.New("A senha é obrigatória e não pode estar em branco")
	}

	if usuario.Nick == "" {
		return errors.New("O apelido é obrigatório e não pode estar em branco")
	}

	return nil
}

func (usuario *Usuario) formatar() {
	usuario.Email = strings.TrimSpace(usuario.Email)
	usuario.Nick = strings.TrimSpace(usuario.Nick)
	usuario.Nome = strings.TrimSpace(usuario.Nome)
}
