package connection

import (
	"database/sql"
	"github.com/jhamiltonjunior/cut-url/internal/domain/entities/url"
	"github.com/jhamiltonjunior/cut-url/internal/domain/repository/url_repository"
)

type Connection interface {
	GetConnection() (*sql.DB, error)
}

type DBConnection struct {
	connection url_repository.Repository
}

func NewDBConnection(conn MySQLConnection) *DBConnection {
	return &DBConnection{
		connection: &conn,
	}
}

func (db *DBConnection) GetConnection() (*sql.DB, error) {
	return db.connection.GetConnection()
}

func (db *DBConnection) CreateULR(url *url.URL) (int64, error) {
	return db.connection.CreateULR(url)
}
