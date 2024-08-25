package interfaces_usecase

import "github.com/jhamiltonjunior/cut-url/internal/domain/entities"

type Hash interface {
	Encrypt(pass string) (string, error)
	Compare(pass, hash string) bool
}

type Token interface {
	GenerateToken(user entities.UserID, roles []string, permissions map[string]string) (string, error)
	ValidateToken()
}
