package mysql_implement

type MySQLConnection struct {
	Host     string
	User     string
	Pass     string
	Database string
}

func (m *MySQLConnection) Connection() {
	
}
