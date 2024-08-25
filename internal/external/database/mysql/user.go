package mysql

import (
	"errors"
	"github.com/jhamiltonjunior/cut-url/internal/domain/entities"
	"github.com/jhamiltonjunior/cut-url/internal/external/service"
)

func (m *Connection) CreateUser(user entities.User) (int64, error) {
	db, err := m.GetConnection()
	if err != nil {
		return 0, err
	}

	user.Password, err = service.Encrypt(user.Password)
	if err != nil {
		return 0, err
	}

	query := `INSERT INTO users (name, email, password) VALUES (?, ?, ?)`

	result, err := db.Exec(query, user.Name, user.Email, user.Password)
	if err != nil {
		return 0, nil
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, nil
	}

	return id, nil
}

func (m *Connection) GetUserByID(user entities.User) (*entities.User, error) {
	db, err := m.GetConnection()
	if err != nil {
		return nil, err
	}

	query := `
		SELECT 
		    name, email, password, created_at
		FROM 
		    users
		WHERE
		    id = ?
		AND active = 1
	`

	err = db.QueryRow(query, user.ID).Scan(&user.Name, &user.Email, &user.Password, &user.CreateAt)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (m *Connection) Auth(user entities.User) (*entities.User, error) {
	return nil, errors.New("")
}

func (m *Connection) UpdateUser(user entities.User) error {
	return nil
}

func (m *Connection) DeleteUser(user entities.User) error {
	return nil
}
