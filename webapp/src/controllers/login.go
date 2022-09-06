package controllers

import (
	"fmt"
	"net/http"
)

// CarregarTelaDeLogin carrega a tela de login
func CarregarTelaDeLogin(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
	w.Write([]byte("Rota login"))
}	