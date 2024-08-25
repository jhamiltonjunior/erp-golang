package repository

import "github.com/jhamiltonjunior/cut-url/internal/domain/entities"

type UserRepository interface {
	Create(user *entities.User) (int, error)
	GetByID(user *entities.User) (*entities.User, error)
	Auth(user entities.User) (*entities.User, error)
	Update(user entities.User) error
	Delete(user entities.User) error
}
