package connection

import (
	"database/sql"
	"github.com/jhamiltonjunior/cut-url/internal/external/mysql_implement"
)

type Connection interface {
	Connection() (*sql.DB, error)
}

func NewConnectionDB(data *mysql_implement.MySQLConnection) Connection {
	return &mysql_implement.MySQLConnection{
		Host:     data.Host,
		User:     data.User,
		Pass:     data.Pass,
		Database: data.Database,
	}
}
