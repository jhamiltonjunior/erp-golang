package database

import (
	"github.com/jhamiltonjunior/cut-url/internal/domain/repository/url_repository"
	"github.com/jhamiltonjunior/cut-url/internal/external/database/mysql"
)

func NewMySQLURLRepository(dsn string) url_repository.Repository {
	return &mysql.Connection{Dsn: dsn}
}
