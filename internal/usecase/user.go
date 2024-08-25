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

func (u *User) GetByID(id entities.UserID) (*entities.User, error) {

	// checar token permissoes etc

	var user *entities.User

	user, err := u.repo.GetUserByID(id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *User) UpdateByID(user entities.User) error {
	// checar token permissoes etc

	err := u.repo.UpdateUser(user)
	if err != nil {
		return err
	}

	return nil
}

func (u *User) Auth(user *entities.User) (string, error) {
	// para um sistema maior as permissoes deveriam vim do banco de dados

	user, err := u.repo.Auth(*user)
	if err != nil {
		return "", nil
	}

	roles := []string{
		"user",
	}
	permissions := map[string]string{
		"dashboard": "read",
	}

	token, err := u.tokenManager.GenerateToken(entities.UserID(user.ID), roles, permissions)
	if err != nil {
		return "", nil
	}

	return token, nil
}

func (u *User) Delete(id entities.UserID) error {
	// usuarios devem deletar somente a si mesmos a menos que sejam admin

	err := u.repo.DeleteUser(id)
	if err != nil {
		return err
	}

	return nil
}
