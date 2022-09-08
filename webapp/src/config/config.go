package config

import (
	"log"
	"os"
	"strconv"
)

var (
	// APIURL representa a URL para comunicação com a API
	APIURL = ""
	// Porta onde a aplicação web está rodando
	Porta = 0
	// HashKey é usada para autenticar o cookie
	HashKey []byte
	// BlockKey é usada para criptografar os dados do cookie
	BlockKey []byte
)

// Carregar inicializa as variáveis de ambiente
func Carregar() {
	var erro error

	if erro = godotenv.Load(); erro != nil {
		
	}
}
