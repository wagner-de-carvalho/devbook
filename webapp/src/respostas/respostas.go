package respostas

import (
	"encoding/json"
	"log"
	"net/http"
)

// Erro represenat a resposta de erro da API
type ErroAPI struct {
	Erro string `json:"erro"`
}

// JSON retorna resposta em formato JSON para a requisição
func JSON(w http.ResponseWriter, statusCode int, dados interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if erro := json.NewEncoder(w).Encode(dados); erro != nil {
		log.Fatal(erro)
	}
}

// TratarStatusCodeDeErro trata as requisições com statusCode da faixa 400 
func TratarStatusCodeDeErro(w http.ResponseWriter, r *http.Response) {
	var erro ErroAPI
	json.NewDecoder(r.Body).Decode(&erro)
	JSON(w, r.StatusCode, erro)
}
