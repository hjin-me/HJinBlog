package fw

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"text/template"
)

type Context struct {
	Params map[string]string
	Res    http.ResponseWriter
	Req    *http.Request
}

func (c Context) Output(data interface{}, contentType string) {
	h := c.Res.Header()
	h.Add("content-type", contentType) 

	fmt.Fprintf(c.Res, "%s", data)
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
	texec := template.Must(template.ParseFiles(tpl))
	texec.Execute(c.Res, data)
}
