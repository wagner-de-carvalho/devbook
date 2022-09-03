package rotas

import (
	"api/src/controllers"
	"net/http"
)

var rotalogin = Rota{
	URI: "/login",
	Metodo: http.MethodPost,
	Funcao: controllers.Login,
	RequerAutenticacao: false,
}