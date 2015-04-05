package transform

import (
	"da"
	"log"
)

func importPosts() error {
	var (
		targetTableName = "blog_posts"
		sourceTableName = "content_articles"

		readSQL   = "SELECT articleid, uid, content, cid, dateline, title, description, keywords FROM " + sourceTableName
		insertSQL = "INSERT INTO " + targetTableName + "(id, alias, uid, content, cid, pubtime, title, description, keywords) VALUES (?,?,?,?,?,?,?,?,?)"
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
	db.Exec("delete from " + targetTableName)

	insertStmt, err := db.Prepare(insertSQL)
	if err != nil {
		log.Println("insert")
		return err
	}
	defer insertStmt.Close()

	rows, err := readStmt.Query()
	if err != nil {
		return err
	}

	for rows.Next() {
		var (
			id, uid, category, pubtime        int64
			content, title, description, tags string
		)
		err := rows.Scan(&id, &uid, &content, &category, &pubtime, &title, &description, &tags)
		if err != nil {
			return err
		}
		_, err = insertStmt.Exec(id, id, uid, content, category, pubtime, title, description, tags)
		if err != nil {
			return err
		}
	}

	return nil
}
