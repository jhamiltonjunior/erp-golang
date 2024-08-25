package usecase

import (
	"github.com/jhamiltonjunior/cut-url/internal/domain/entities"
	"github.com/jhamiltonjunior/cut-url/internal/domain/repository"
	"github.com/jhamiltonjunior/cut-url/internal/usecase/interfaces_usecase"
)

type User struct {
	repo         repository.User
	tokenManager interfaces_usecase.Token
}

func NewUserUseCase(repo repository.User) *User {
	return &User{
		repo: repo,
	}
}

func (u *User) Create(user entities.User) (string, error) {
	id, err := u.repo.CreateUser(user)
	if err != nil {
		return "", nil
	}

	roles := []string{
		"user",
	}
	permissions := map[string]string{
		"dashboard": "read",
	}

	token, err := u.tokenManager.GenerateToken(id, roles, permissions)
	if err != nil {
		return "", err
	}

	return token, nil
}
