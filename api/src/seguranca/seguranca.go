package seguranca

import "golang.org/x/crypto/bcrypt"

// Hash recebe uma senha e adiciona um hash
func Hash(senha string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(senha), bcrypt.DefaultCost)
}

// VerificarSenha verifica se senhas são iguais
func VerificarSenha(senhaComHash string, senhaString string) error {
	return bcrypt.CompareHashAndPassword([]byte(senhaComHash), []byte(senhaString))
}