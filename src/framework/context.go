package fw

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"text/template"

	"golang.org/x/net/context"
)

type FwContext interface {
	context.Context
	Res() http.ResponseWriter
	Req() *http.Request
	Params() map[string]string
}

type httpContext struct {
	context.Context
}

func (c httpContext) Res() http.ResponseWriter {
	res, ok := c.Value("Res").(http.ResponseWriter)
	if !ok {
		panic(errors.New("Res is not Res"))
	}
	return res
}

func (c httpContext) Req() *http.Request {
	req, ok := c.Value("Req").(*http.Request)
	if !ok {
		panic(errors.New("Req is not Req"))
	}
	return req
}

func (c httpContext) Params() map[string]string {
	params, ok := c.Value("Params").(map[string]string)
	if !ok {
		panic(errors.New("Params is not Params"))
	}
	return params
}

func (c httpContext) Output(data interface{}, contentType string) {
	res := c.Res()
	h := res.Header()
	h.Add("content-type", contentType)

	fmt.Fprintf(res, "%s", data)
}

func (c httpContext) Json(data interface{}) {
	res := c.Res()
	h := res.Header()
	h.Add("content-type", "application/json")

	str, _ := json.Marshal(data)
	fmt.Fprintf(res, "%s", str)
}

func (c httpContext) Tpl(data interface{}, path string) {
	res, ok := c.Res()
	if !ok {
		panic(errors.New("Res is not Res"))
		return
	}
	viewBase := os.Getenv("GOVIEW")
	tpl := viewBase + path
	texec := template.Must(template.ParseFiles(tpl))
	texec.Execute(res, data)
}
