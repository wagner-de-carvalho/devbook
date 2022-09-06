package router

import "github.com/gorilla/mux"

// Gerar retorna um router contendo todas as rotas configuradas
func Gerar() *mux.Router {
	return mux.NewRouter()
}