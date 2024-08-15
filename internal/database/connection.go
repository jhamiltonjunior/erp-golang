package connection

import "github.com/jhamiltonjunior/cut-url/internal/database/mysql_implement"

type Connection interface {
	Connection()
}

func NewConnectionDB(data *mysql_implement.MySQLConnection) Connection {

	return &mysql_implement.MySQLConnection{
		Host:     data.Host,
		User:     data.User,
		Pass:     data.Pass,
		Database: data.Database,
	}
}
