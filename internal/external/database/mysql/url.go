package mysql

import (
	"fmt"
	"github.com/jhamiltonjunior/cut-url/internal/domain/entities/url"
)

func (m *Connection) CreateULR(url *url.URL) (int64, error) {
	connection, err := m.GetConnection()
	if err != nil {
		panic(err)
		return 0, err
	}

	query := fmt.Sprintf(
		"INSERT INTO urls (original, destination, user_id, description) VALUES ('%s', '%s', '%d', '%s')",
		url.OriginalURL, url.DestinationURL, url.UserID, url.Description)

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

func (m *Connection) GetAllByUser(id int) ([]*url.URL, error) {
	connection, err := m.GetConnection()
	if err != nil {
		panic(err)
		return nil, err
	}

	var urlSlice []*url.URL
	query := `
		SELECT 
    		id, user_id, original, destination
		FROM
		    urls 
		WHERE user_id = ?
		AND active = 1
`

	rows, err := connection.Query(query, id)
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

func (m *Connection) GetByName(description string) ([]url.URL, error) {
	db, err := m.GetConnection()
	if err != nil {
		return nil, err
	}

	var u []url.URL

	query := `
			SELECT 
			    id, original, destination, description 
			FROM 
			    urls 
			WHERE 
			    description LIKE ?
				AND active = 1
			LIMIT 20
			`

	rows, err := db.Query(query, "%"+description+"%")
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var id int
		var original, destination, desc string

		err := rows.Scan(&id, &original, &destination, &desc)
		if err != nil {
			return nil, err
		}

		u = append(u, url.URL{
			Id:          id,
			OriginalURL: original,
			Description: desc,
		})
	}

	return u, nil
}

func (m *Connection) UpdateById(u *url.URL) error {
	db, err := m.GetConnection()
	if err != nil {
		return err
	}

	query := `UPDATE urls SET destination = ?, description = ? WHERE id = ?`

	if _, err = db.Exec(query, u.DestinationURL, u.Description, u.Id); err != nil {
		return err
	}

	return nil
}

func (m *Connection) DeleteById(id int) error {
	db, err := m.GetConnection()
	if err != nil {
		return err
	}

	query := `UPDATE urls SET active = 0 WHERE id = ?`

	if _, err = db.Exec(query, id); err != nil {
		return err
	}

	return nil
}