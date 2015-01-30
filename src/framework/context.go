package fw

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"text/template"
<<<<<<< HEAD

	"github.com/russross/blackfriday"
=======
>>>>>>> FETCH_HEAD
)

type Context struct {
	Params map[string]string
	Res    http.ResponseWriter
	Req    *http.Request
}

<<<<<<< HEAD
func markDowner(args ...interface{}) string {
	s := blackfriday.MarkdownCommon([]byte(fmt.Sprintf("%s", args...)))
	return string(s)
}

func (c Context) Output(data interface{}) {

=======
func (c Context) Output(data interface{}, contentType string) {
	h := c.Res.Header()
	h.Add("content-type", contentType) 

	fmt.Fprintf(c.Res, "%s", data)
>>>>>>> FETCH_HEAD
}

func (c Context) Json(data interface{}) {
	h := c.Res.Header()
	h.Add("content-type", "application/json")

	str, _ := json.Marshal(data)
	fmt.Fprintf(c.Res, "%s", str)
}

func (c Context) Tpl(data interface{}, path string) {
	viewBase := os.Getenv("GOVIEW")
	tpl := viewBase + path
<<<<<<< HEAD
	texec := template.Must(template.ParseFiles(tpl).Func(template.FuncMap{"markDown": markDowner}))
=======
	texec := template.Must(template.ParseFiles(tpl))
>>>>>>> FETCH_HEAD
	texec.Execute(c.Res, data)
}
