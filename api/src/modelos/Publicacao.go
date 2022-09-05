package modelos

import "time"

// Publicacao representa uma publicação feita por um usuario
type Publicacao struct {
	ID uint64 `json:"id,omitempty"`
	Titulo string `json:"titulo,omitempty"`
	Conteudo string `json:"conteudo,omitempty"`
	AutorID int64 `json:"autorId,omitempty"`
	Curtidas int64 `json:"curtidas"`
	CriadaEm time.Time `json:"criadaEm,omitempty"`
}