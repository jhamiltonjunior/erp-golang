package mysql

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type Connection struct {
	Dsn string
}

func (m *Connection) GetConnection() (*sql.DB, error) {
	db, err := sql.Open("mysql", m.Dsn)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
