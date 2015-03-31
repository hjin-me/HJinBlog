package models

import (
	"da"
	"log"
)

const (
	TABLE_NAME_USERS = "blog_users"
	INSERT_USER      = "INSERT INTO " + TABLE_NAME_POSTS + "(content, category, pubtime, title, description, tags) VALUES (?,?,?,?,?,?)"
	UPDATE_USER      = "UPDATE " + TABLE_NAME_POSTS + " SET content=?, category=?, pubtime=?, title=?, description=?, tags=? WHERE id=?"
	QUERY_USERS      = "SELECT id, content, category, pubtime, title, description, tags FROM " + TABLE_NAME_POSTS + " ORDER BY pubtime DESC LIMIT ?,?"
	FIND_USER        = "SELECT id, , category, pubtime, title, description, tags FROM " + TABLE_NAME_POSTS + " WHERE id = ?"
	VALIDATE_USER    = "SELECT count(*) FROM " + TABLE_NAME_USERS + " WHERE username=? AND pwd=? LIMIT 1"
)

func UserCheck(username, pwd string) (bool, error) {

	db, err := da.Connect()
	if err != nil {
		return false, err
	}
	// Prepare statement for reading data
	stmt, err := db.Prepare(VALIDATE_USER)
	if err != nil {
		return false, err
	}
	defer stmt.Close()

	var count int

	err = stmt.QueryRow(username, pwd).Scan(&count)
	if err != nil {
		return false, err
	}
	log.Println(username, pwd, count)
	if count == 0 {
		return false, nil
	}
	return true, nil

}
