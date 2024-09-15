package mysql

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jhamiltonjunior/erp-golang/internal/domain/repository"
	"github.com/jhamiltonjunior/erp-golang/internal/usecase/interfaces_usecase"
)

type Connection struct {
	dsn         string
	hashManager interfaces_usecase.Hash
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

func NewMySQLUserRepository(dsn string, hash interfaces_usecase.Hash) repository.User {
	return &Connection{dsn: dsn, hashManager: hash}
}
