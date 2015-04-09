package transform

import (
	"da"
	"log"
)

func importPosts() error {
	var (
		targetTableName = "blog_posts"
		sourceTableName = "huangjin_blog"

		readSQL   = "SELECT id, alias, content, category, pubtime, title, description, tags FROM " + sourceTableName
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
			id, uid, pubtime, cid                              int64
			alias, category, content, title, description, tags string
		)
		err := rows.Scan(&id, &alias, &content, &category, &pubtime, &title, &description, &tags)
		switch category {
		case "it":
			cid = 1
		case "chat":
			cid = 2
		case "live":
			cid = 3
		default:
			cid = 0
		}
		uid = 1
		if err != nil {
			return err
		}
		_, err = insertStmt.Exec(id, alias, uid, content, cid, pubtime, title, description, tags)
		if err != nil {
			return err
		}
	}

	return nil
}
