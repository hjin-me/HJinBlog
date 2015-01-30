package actions

import (
	"bytes"
	"database/sql"
	"fmt"
	"framework"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	//"os"

	"code.google.com/p/go-uuid/uuid"
	_ "github.com/mattn/go-sqlite3"
)

func Init(w http.ResponseWriter, r *http.Request, context fw.Context) {

	fmt.Println(context.Params)
	fmt.Println("======")
	u := uuid.New()
	fmt.Println(u)
	userFile := os.Getenv("GOPATH") + "/src/install/createdb.sql"
	by, err := ioutil.ReadFile(userFile)
	if err != nil {
		log.Fatal(err)
		return
	}
	var buffer bytes.Buffer
	buffer.Write(by)
	createTableSql := buffer.String()
	dbp := os.Getenv("GOPATH") + "/src/blog.db"
	db, err := sql.Open("sqlite3", dbp)
	if err != nil {
		log.Fatal(err)
		return
	}

	_, e1 := db.Exec(createTableSql)
	// db, err := sql.Open("mysql", "root:hj111111@tcp(10.211.55.8:3306)/hjinblog")
	if e1 != nil {
		log.Fatal(e1)
		return
	}
	io.WriteString(w, "install success")

}
