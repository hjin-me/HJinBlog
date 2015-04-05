package category

import (
	"da"
	"database/sql"
)

const (
	TABLE_NAME_CATEGORY = "blog_categories"
	UPDATE_CATEGORY     = "INSERT INTO " + TABLE_NAME_CATEGORY + " (name, alias) VALUES (?,?) ON DUPLICATE KEY UPDATE alias=VALUES(alias)"
	QUERY_CATEGORY      = "SELECT name, alias FROM " + TABLE_NAME_CATEGORY + " LIMIT ?, ?"
)

type Category struct {
	Id    int
	Name  string
	Alias string
}

func Query() ([]Category, error) {
	db, err := da.Connect()
	if err != nil {
		return nil, err
	}
	var (
		stmt *sql.Stmt
	)
	stmt, err = db.Prepare(QUERY_CATEGORY) // ? = placeholder
	if err != nil {
		return nil, err
	}
	defer stmt.Close() // Close the statement when we leave main() / the program terminates

	var (
		name, alias string
		cs          []Category
	)
	rows, err := stmt.Query(0, 10)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		err = rows.Scan(&name, &alias)
		if err != nil {
			panic(err)
		}
		var (
			c Category
		)
		c.Name = name
		c.Alias = alias

		cs = append(cs, c)
	}
	return cs, nil

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

	_, err = stmt.Exec(c.Name, c.Alias)
	if err != nil {
		return err
	}
	return nil
}
