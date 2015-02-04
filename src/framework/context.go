package fw

import (
	"encoding/json"
	"fmt"
	"log"
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
	Output(interface{}, string)
	Json(interface{})
}

type httpContext struct {
	context.Context
	w   http.ResponseWriter
	r   *http.Request
	p   map[string]string
	tpl string
}

func WithHttp(parent context.Context, w http.ResponseWriter, r *http.Request, p map[string]string, tpl string) FwContext {
	return &httpContext{parent, w, r, p, tpl}
}

func (c *httpContext) Res() http.ResponseWriter {
	return c.w
}

func (c *httpContext) Req() *http.Request {
	return c.r
}

func (c *httpContext) Params() map[string]string {
	return c.p
}

func (c *httpContext) Output(data interface{}, contentType string) {
	res := c.Res()

	select {
	case <-c.Done():
		log.Println("request timeout", c.Err())
	default:
		h := res.Header()
		h.Add("content-type", contentType)
		fmt.Fprintf(res, "%s", data)
	}
}

func (c *httpContext) Json(data interface{}) {
	res := c.Res()
	select {
	case <-c.Done():
		log.Println("request timeout", c.Err())
	default:
		h := res.Header()
		h.Add("content-type", "application/json")

		str, _ := json.Marshal(data)
		fmt.Fprintf(res, "%s", str)
	}
}

func (c *httpContext) Tpl(data interface{}, path string) {
	res := c.Res()
	viewBase := os.Getenv("GOVIEW")
	tpl := viewBase + path
	texec := template.Must(template.ParseFiles(tpl))
	select {
	case <-c.Done():
		log.Println("request timeout", c.Err())
	default:
		texec.Execute(res, data)

	}
}
