package url_repository

import (
	"database/sql"
	"github.com/jhamiltonjunior/cut-url/internal/domain/entities/url"
)

type Repository interface {
	GetConnection() (*sql.DB, error)
	CreateULR(url *url.URL) (int64, error)
}
