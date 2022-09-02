package controllers

import "net/http"

// Cria um usuário no banco de dados
func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Criando usuário"))
}

// Busca usuários no banco de dados
func BuscarUsuarios(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscando usuários"))
}

// Busca um usuário no banco de dados
func BuscarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscando usuários"))
}

// Busca um usuário no banco de dados
func AtualizarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Atualizando usuário"))
}

// Exclui um usuário no banco de dados
func ExcluirUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Excluindo usuário"))
}
