package rotas

import (
	"github.com/gorilla/mux"
	"net/http"
)

// Rota representa todas as rotas da aplicação web
type Rota struct {
	URI                string
	Metodo             string
	Funcao             func(http.ResponseWriter, *http.Request)
	RequerAutenticacao bool
}

// Configurar insere todas as rotas no router
func Configurar(router *mux.Router) *mux.Router {
	rotas := rotasLogin
	for _, rota := range rotas {
		router.HandleFunc(rota.URI, rota.Funcao).Methods(rota.Metodo)
	}

	return router
}
