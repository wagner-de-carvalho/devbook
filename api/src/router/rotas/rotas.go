package rotas

import "net/http"

// Representa todas as Rotas da API
type Rota struct {
	URI string
	Metodo string
	Funcao func(http.ResponseWriter, *http.Request)
	RequerAutenticacao bool
}