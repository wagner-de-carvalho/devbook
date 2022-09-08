package main

import (
	"fmt"
	"log"
	"net/http"
	"webapp/src/router"
	"webapp/src/utils"
)

func main() {
	config.Carregar()
	utils.CarregarTemplates()
	r := router.Gerar()

	fmt.Printf("Rodando webapp na porta %s", config.Porta)
	fmt.Println("Rodando webapp na porta 3000!")
	log.Fatal(http.ListenAndServe(":3000", r))
}
