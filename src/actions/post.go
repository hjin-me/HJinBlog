package actions

import (
	"fmt"
	"framework"
	"log"
	"models"
	"net/http"

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
	Tags        []string
	Keywords    []string
}

func PostReadBatch(w http.ResponseWriter, r *http.Request, context fw.Context) {
	fmt.Println(context.Params)

	article, err := models.ReadBatch()
	if err != nil && err != models.ErrNotFound {
		log.Fatal(err)
	}

	context.Json(article)
}

func PostRead(w http.ResponseWriter, r *http.Request, context fw.Context) {
	fmt.Println(context.Params)

	id := context.Params["prefix"] + "-" + context.Params["id"]
	article, err := models.Read(id)
	if err != nil && err != models.ErrNotFound {
		log.Fatal(err)
	}
	fmt.Println(article.Content)
	context.Tpl(article, "/post.html")

	// context.Json(article)
}

//// Post(w http.ResponseWriter, r *http.Request, context fw.Context) {
////r.ParseForm()       //解析参数，默认是不会解析的
////fmt.Println(r.Form) //这些信息是输出到服务器端的打印信息
////fmt.Println("path", r.URL.Path)
////fmt.Println("scheme", r.URL.Scheme)

////fmt.Println(context.Params)

////aid := context.Params["aa"]

////var dbp string
////dbp = os.Getenv("GOPATH") + "/src/blog.db"
////fmt.Println(dbp)

////db, err := sql.Open("sqlite3", dbp)
////if err != nil {
////	log.Fatal(err)
////	return
////}

////stmt, err := db.Prepare("SELECT * FROM huangjin_blog WHERE alias = ? LIMIT 0, 1")
////if err != nil {
////	log.Fatal(err)
////	return
////}

////row := stmt.QueryRow(aid)
////var A Article
////err = row.Scan(&A.Id, &A.Alias, &A.Content, &A.Category, &A.Pubtime, &A.Title, &A.Description, &A.Tags, &A.Keywords)
////if err != nil {
////	if err == sql.ErrNoRows {
////		fmt.Println("no result")
////		http.NotFound(w, r)
////		return
////	}
////	log.Fatal(err)
////	return
////}
////viewBase := os.Getenv("GOVIEW")
////tpl := viewBase + "page.gtpl"
////texec := template.Must(template.ParseFiles(tpl))
////texec.Execute(w, A)
