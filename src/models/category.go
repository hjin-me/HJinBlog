package models

import (
	"da"
	"database/sql"
)

const (
	TABLE_NAME_CATEGORY = "blog_categories"
	UPDATE_CATEGORY     = "INSERT INTO " + TABLE_NAME_CATEGORY + " (name, description) VALUES (?,?) ON DUPLICATE KEY UPDATE description=VALUES(description)"
	QUERY_CATEGORY      = "SELECT name, description FROM " + TABLE_NAME_CATEGORY + " LIMIT ?, ?"
)

type Category struct {
	Name        string
	Description string
}

func (c *Category) Save() error {
	db, err := da.Connect()
	if err != nil {
		return err
	}
	var (
		stmt *sql.Stmt
	)
	stmt, err = db.Prepare(UPDATE_CATEGORY) // ? = placeholder
	if err != nil {
		return err
	}
	defer stmt.Close() // Close the statement when we leave main() / the program terminates

	_, err = stmt.Exec(c.Name, c.Description)
	if err != nil {
		return err
	}
	return nil
}
