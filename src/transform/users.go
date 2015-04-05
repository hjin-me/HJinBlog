package transform

import (
	"da"
	"log"
	"models/user"
)

var (
	targetTableName = "blog_users"
	sourceTableName = "content_users"

	readSQL   = "SELECT userid, username, regdateline FROM " + sourceTableName
	insertSQL = "INSERT INTO " + targetTableName + " ( id, username, pwd, privilege, create_time) VALUES (?,?,?,?,?)"
)

func importUsers() error {
	db, err := da.Connect()
	if err != nil {
		return err
	}
	readStmt, err := db.Prepare(readSQL)
	if err != nil {
		log.Println("read")
		return err
	}
	defer readStmt.Close()
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
			uid      int
			username string
			createTS int64
		)
		err := rows.Scan(&uid, &username, &createTS)
		if err != nil {
			return err
		}

		pwd := "xiaoyeye"
		privilege := 0xFFFF

		_, err = insertStmt.Exec(uid, username, user.Hash(pwd), privilege, createTS)
		if err != nil {
			return err
		}
	}
	return nil

}
