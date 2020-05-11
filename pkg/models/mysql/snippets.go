package mysql

import (
	"database/sql"
	"errors"

	"github.com/victorhsn/snippetbox/pkg/models"
)

//SnippetModels to create a sql.DB connection pool
type SnippetModels struct {
	DB *sql.DB
}

//Insert method to save a Snippet on database
func (m *SnippetModels) Insert(title, content, expires string) (int, error) {

	stmt := `INSERT INTO snippets (title, content, created, expires) 
			 VALUES(?,?, UTC_TIMESTAMP(), DATE_ADD(UTC_TIMESTAMP(), INTERVAL ? DAY))`

	result, err := m.DB.Exec(stmt, title, content, expires)
	if err != nil {
		return 0, nil
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

//Get method to get a Snippet from database
func (m *SnippetModels) Get(id int) (*models.Snippet, error) {

	s := &models.Snippet{}
	stmt := `SELECT id, title, content, created, expires FROM snippets WHERE expires > UTC_TIMESTAMP() AND id = ?`
	err := m.DB.QueryRow(stmt, id).Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, models.ErrNoRecord
		} else {
			return nil, err
		}
	}

	return s, nil
}

//Latest this method return the 10 most recently Snippets from database
func (m *SnippetModels) Latest() ([]*models.Snippet, error) {
	return nil, nil
}
