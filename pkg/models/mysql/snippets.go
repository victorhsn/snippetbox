package mysql

import (
	"database/sql"

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
	return nil, nil
}

//Latest this method return the 10 most recently Snippets from database
func (m *SnippetModels) Latest() ([]*models.Snippet, error) {
	return nil, nil
}
