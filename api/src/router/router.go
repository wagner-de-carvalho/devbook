package router

import "github.com/gorilla/mux"

// Rertorna Router
func Gerar() *mux.Router {
	return mux.NewRouter()
}