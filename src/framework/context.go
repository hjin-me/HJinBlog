package fw

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"text/template"

	"github.com/russross/blackfriday"
)

type Context struct {
	Params map[string]string
	Res    http.ResponseWriter
	Req    *http.Request
}

func markDowner(args ...interface{}) string {
	s := blackfriday.MarkdownCommon([]byte(fmt.Sprintf("%s", args...)))
	return string(s)
}

func (c Context) Output(data interface{}) {

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
	texec := template.Must(template.ParseFiles(tpl).Func(template.FuncMap{"markDown": markDowner}))
	texec.Execute(c.Res, data)
}
