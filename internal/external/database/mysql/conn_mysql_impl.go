package mysql

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jhamiltonjunior/cut-url/internal/domain/repository"
	"github.com/jhamiltonjunior/cut-url/internal/domain/repository/url_repository"
	"github.com/jhamiltonjunior/cut-url/internal/usecase"
)

type Connection struct {
	dsn         string
	hashManager usecase.Hash
}

func (m *Connection) GetConnection() (*sql.DB, error) {
	db, err := sql.Open("mysql", m.dsn)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

func NewMySQLURLRepository(dsn string) url_repository.Repository {
	return &Connection{dsn: dsn}
}

func NewMySQLUserRepository(dsn string, hash *usecase.Hash) repository.User {
	return &Connection{dsn: dsn, hashManager: *hash}
}
