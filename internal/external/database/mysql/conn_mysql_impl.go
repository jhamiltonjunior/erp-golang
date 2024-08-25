package mysql

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jhamiltonjunior/cut-url/internal/domain/repository"
	"github.com/jhamiltonjunior/cut-url/internal/domain/repository/url_repository"
)

type Connection struct {
	dsn string
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

func NewMySQLUserRepository(dsn string) repository.User {
	return &Connection{dsn: dsn}
}
