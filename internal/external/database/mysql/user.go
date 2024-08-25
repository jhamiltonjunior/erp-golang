package mysql

import (
	"database/sql"
	"errors"
	"github.com/jhamiltonjunior/cut-url/internal/domain/entities"
)

func (m *Connection) CreateUser(user entities.User) (int64, error) {
	db, err := m.GetConnection()
	if err != nil {
		return 0, err
	}
	defer func(db *sql.DB) {
		_ = db.Close()
	}(db)

	user.Password, err = m.hashManager.Encrypt(user.Password)
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
	defer func(db *sql.DB) {
		_ = db.Close()
	}(db)

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
	db, err := m.GetConnection()
	if err != nil {
		return nil, err
	}
	defer func(db *sql.DB) {
		_ = db.Close()
	}(db)

	query := `
		SELECT
			id, password
		FROM
			users
		WHERE
		    email = ?
	`
	var hash string

	err = db.QueryRow(query, user.Email).Scan(&user.ID, &hash)
	if err != nil {
		return nil, err
	}

	if !m.hashManager.Compare(user.Password, hash) {
		return nil, errors.New("incorrect password")
	}

	return &user, nil
}

func (m *Connection) UpdateUser(user entities.User) error {
	db, err := m.GetConnection()
	if err != nil {
		return err
	}
	defer func(db *sql.DB) {
		_ = db.Close()
	}(db)

	query := `
		UPDATE 
		    users 
		SET
		    name = ?,
		    password = ?
		WHERE
		    id = ?
	`

	newHash, err := m.hashManager.Encrypt(user.Password)
	if err != nil {
		return err
	}

	err = db.QueryRow(query, user.Name, newHash, user.ID).Scan()
	if err != nil {
		return err
	}

	return nil
}

func (m *Connection) DeleteUser(user entities.User) error {
	db, err := m.GetConnection()
	if err != nil {
		return err
	}
	defer func(db *sql.DB) {
		_ = db.Close()
	}(db)

	query := `
		UPDATE 
		    users 
		SET
		    active = 0
		WHERE
		    id = ?
	`

	err = db.QueryRow(query, user.ID).Scan()
	if err != nil {
		return err
	}

	return nil
}
