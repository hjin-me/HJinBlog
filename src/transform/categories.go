package transform

import (
	"da"
	"log"
)

func importCategories() error {
	var (
		readSQL   = "select cid, name, url from content_categories"
		insertSQL = "INSERT INTO blog_categories ( id, name, alias) VALUES (?,?,?)"
	)
	db, err := da.Connect()
	if err != nil {
		return err
	}

	readStmt, err := db.Prepare(readSQL)
	if err != nil {
		return err
	}

	defer readStmt.Close()
	insertStmt, err := db.Prepare(insertSQL)
	if err != nil {
		log.Println("insert")
		return err
	}
	defer insertStmt.Close()
	_, err = db.Exec("delete from blog_categories")
	if err != nil {
		panic(err)
	}
	log.Println("hehe")

	rows, err := readStmt.Query()
	if err != nil {
		return err
	}

	for rows.Next() {
		var (
			id    int
			name  string
			alias string
		)
		err := rows.Scan(&id, &alias, &name)
		log.Println(id, name, alias)
		if err != nil {
			return err
		}
		_, err = insertStmt.Exec(id, name, alias)
		if err != nil {
			return err
		}
	}

	return nil
}
