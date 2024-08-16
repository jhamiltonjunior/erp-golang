package mysql_implement

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type MySQLConnection struct {
	Host     string
	User     string
	Pass     string
	Database string
}

func (m *MySQLConnection) Connection() (*sql.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s", m.User, m.Pass, m.Host, m.Database)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	return db, nil
}
