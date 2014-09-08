package actions

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"text/template"

	_ "github.com/mattn/go-sqlite3"
)

// INSERT INTO `huangjin_blog` (`id`, `alias`, `content`, `category`, `pubtime`, `title`, `description`, `tags`, `keywords`) VALUES
type Article struct {
	Id          string
	Alias       string
	Content     string
	Category    string
	Pubtime     int32
	Title       string
	Description string
	Tags        string
	Keywords    string
}

func Post(w http.ResponseWriter, r *http.Request, params []string) {
	r.ParseForm()       //解析参数，默认是不会解析的
	fmt.Println(r.Form) //这些信息是输出到服务器端的打印信息
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)

	fmt.Println(params)

	aid := params[0]

	var dbp string
	dbp = os.Getenv("GOPATH") + "/src/blog.db"
	fmt.Println(dbp)

	db, err := sql.Open("sqlite3", dbp)
	if err != nil {
		log.Fatal(err)
		return
	}

	stmt, err := db.Prepare("SELECT * FROM huangjin_blog WHERE alias = ? LIMIT 0, 1")
	if err != nil {
		log.Fatal(err)
		return
	}

	row := stmt.QueryRow(aid)
	var A Article
	err = row.Scan(&A.Id, &A.Alias, &A.Content, &A.Category, &A.Pubtime, &A.Title, &A.Description, &A.Tags, &A.Keywords)
	if err != nil {
		log.Fatal(err)
		return
	}
	viewBase := os.Getenv("GOVIEW")
	tpl := viewBase + "page.gtpl"
	texec := template.Must(template.ParseFiles(tpl))
	texec.Execute(w, A)
}
