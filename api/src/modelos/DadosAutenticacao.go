package modelos

// DadosAutenticacao contém token e id do usuário
type DadosAutenticacao struct {
	ID string `json:"id"`
	Token string `json:"token"`
}