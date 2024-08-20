package database

import (
	"database/sql"
	"fmt"
	"github.com/jhamiltonjunior/cut-url/internal/domain/entities/url"
	"github.com/jhamiltonjunior/cut-url/internal/domain/repository/url_repository"

	_ "github.com/go-sql-driver/mysql"
)

type MySQLConnection struct {
	dsn string
}

func NewMySQLURLRepository(dsn string) url_repository.Repository {
	return &MySQLConnection{dsn: dsn}
}

func (m *MySQLConnection) GetConnection() (*sql.DB, error) {
	db, err := sql.Open("mysql", m.dsn)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
func (m *MySQLConnection) CreateULR(url *url.URL) (int64, error) {
	connection, err := m.GetConnection()
	if err != nil {
		panic(err)
		return 0, err
	}

	fmt.Println("cheguei mysql")
	query := fmt.Sprintf("INSERT INTO urls (original, destination, user_id) VALUES ('%s', '%s', '%d')", url.OriginalURL, url.DestinationURL, url.UserID)

	exec, err := connection.Exec(query)
	if err != nil {
		panic(err)
		return 0, err
	}

	urlId, err := exec.LastInsertId()
	if err != nil {
		return 0, err
	}

	return urlId, nil
}

func (m *MySQLConnection) GetAll() ([]*url.URL, error) {
	connection, err := m.GetConnection()
	if err != nil {
		panic(err)
		return nil, err
	}

	var urlSlice []*url.URL
	query := "SELECT id, user_id, original, destination FROM urls"

	rows, err := connection.Query(query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var id, userId int
		var original, destination string

		err := rows.Scan(&id, &userId, &original, &destination)
		if err != nil {
			return nil, err
		}

		urlSlice = append(urlSlice, &url.URL{
			Id:             id,
			UserID:         userId,
			OriginalURL:    original,
			DestinationURL: destination,
		})
	}

	return urlSlice, nil
}
