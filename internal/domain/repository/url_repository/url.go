package url_repository

import (
	"github.com/jhamiltonjunior/cut-url/internal/domain/entities/url"
)

type Repository interface {
	CreateULR(url *url.URL) (int64, error)
	GetAllByUser() ([]*url.URL, error)
	GetByName(description string) ([]url.URL, error)
	UpdateById(id int, u *url.URL) error
	DeleteById(id int) error
}
