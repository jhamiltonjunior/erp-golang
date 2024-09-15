package repository

import "github.com/jhamiltonjunior/erp-golang/internal/domain/entities"

type User interface {
	CreateUser(user entities.User) (entities.UserID, error)
	GetUserByID(id entities.UserID) (*entities.User, error)
	Auth(user entities.User) (*entities.User, error)
	UpdateUser(user entities.User) error
	DeleteUser(user entities.UserID) error
}
