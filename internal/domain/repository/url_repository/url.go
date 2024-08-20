package url_repository

import (
	"github.com/jhamiltonjunior/cut-url/internal/domain/entities/url"
)

type Repository interface {
	CreateULR(url *url.URL) (int64, error)
	GetAll() ([]*url.URL, error)
}
