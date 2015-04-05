package user

import (
	"crypto/sha1"
	"da"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
)

const (
	TABLE_NAME_USERS = "blog_users"
	INSERT_USER      = "INSERT INTO " + TABLE_NAME_USERS + "(username, pwd, privilege, create_time) VALUES (?,?,?,?)"
	UPDATE_USER      = "UPDATE " + TABLE_NAME_USERS + " SET content=?, category=?, pubtime=?, title=?, description=?, tags=? WHERE id=?"
	QUERY_USERS      = "SELECT id, username, privilege FROM " + TABLE_NAME_USERS + " ORDER BY id DESC LIMIT ?,?"
	FIND_USER        = "SELECT id, username, privilege FROM " + TABLE_NAME_USERS + " WHERE id = ?"
	FIND_PRIVILEGE   = "SELECT privilege FROM " + TABLE_NAME_USERS + " WHERE username = ?"
	VALIDATE_USER    = "SELECT count(*) FROM " + TABLE_NAME_USERS + " WHERE username=? AND pwd=? LIMIT 1"
)

var (
	salt    = "xxx"
	Expires = time.Second * 86400
)

type User struct {
	Id        int
	Name      string
	Privilege Privilege
}

func Hash(str string) string {
	s := sha1.Sum([]byte(salt + str))
	return hex.EncodeToString(s[:])
}

func Authentication(username string, p int) (bool, error) {
	db, err := da.Connect()
	if err != nil {
		return false, err
	}

	// Prepare statement for reading data
	stmt, err := db.Prepare(FIND_PRIVILEGE)
	if err != nil {
		return false, err
	}
	defer stmt.Close()

	var privilege int
	// todo load salt form conf
	err = stmt.QueryRow(username).Scan(&privilege)
	if err != nil {
		return false, err
	}

	return privilege&p > 0, nil

}

func Add(username, pwd string, privilege int) error {

	db, err := da.Connect()
	if err != nil {
		return err
	}

	// Prepare statement for reading data
	stmt, err := db.Prepare(INSERT_USER)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(username, Hash(pwd), privilege, time.Now().Unix())
	if err != nil {
		return err
	}

	return nil
}

func Check(username, pwd string) (bool, string, error) {

	db, err := da.Connect()
	if err != nil {
		return false, "", err
	}

	// Prepare statement for reading data
	stmt, err := db.Prepare(VALIDATE_USER)
	if err != nil {
		return false, "", err
	}
	defer stmt.Close()

	var count int

	// todo load salt form conf
	err = stmt.QueryRow(username, Hash(pwd)).Scan(&count)
	if err != nil {
		return false, "", err
	}
	log.Println(username, pwd, Hash(pwd), count)
	if count == 0 {
		return false, "", nil
	}

	ts := strconv.FormatInt(time.Now().Add(Expires).Unix(), 10)
	username = base64.StdEncoding.EncodeToString([]byte(username))
	sign := fmt.Sprintf("%s|%s|%s", username, ts, Hash(username+ts+salt))

	return true, sign, nil
}

func DecodeToken(token string) (isLogin bool, username string, err error) {
	var (
		ts,
		sign string
	)
	sarr := strings.Split(token, "|")
	if len(sarr) != 3 {
		log.Println("length not 3")
		isLogin = false
		return
	}
	username, ts, sign = sarr[0], sarr[1], sarr[2]
	tsInt64, err := strconv.ParseInt(ts, 10, 64)
	if err != nil {
		log.Println("parse int failed")
		return
	}

	if time.Now().After(time.Unix(tsInt64, 0)) {
		log.Println("time is past", tsInt64, time.Now().Unix())
		isLogin = false
		return
	}

	if sign != Hash(username+ts+salt) {
		log.Println("hash not ok", sign, Hash(username+ts+salt))
		isLogin = false
		return
	}

	unameByte, err := base64.StdEncoding.DecodeString(username)
	if err != nil {
		log.Println("parse uname failed")
		return
	}
	username = string(unameByte)
	isLogin = true
	return
}

func Query(start, limit int) ([]User, error) {

	var (
		err      error
		userList []User
	)
	db, err := da.Connect()
	if err != nil {
		return userList, err
	}
	// Prepare statement for reading data
	stmt, err := db.Prepare(QUERY_USERS)
	if err != nil {
		return userList, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(start, limit)
	for rows.Next() {
		u := User{}
		err = rows.Scan(&u.Id, &u.Name, &u.Privilege)
		if err != nil {
			panic(err)
		}

		userList = append(userList, u)
	}

	return userList, nil
}

func FindOne(id int) (user User, err error) {
	db, err := da.Connect()
	if err != nil {
		return
	}
	stmt, err := db.Prepare(FIND_USER)
	if err != nil {
		return
	}
	defer stmt.Close()

	row := stmt.QueryRow(id)
	err = row.Scan(&user.Id, &user.Name, &user.Privilege)
	if err != nil {
		return
	}

	return
}
