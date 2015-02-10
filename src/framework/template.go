package fw

import (
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"log"

	"github.com/russross/blackfriday"
)

var t5xx *template.Template

func Load5xx() *template.Template {
	if t5xx == nil {
		log.Println("load new 5xx tpl")
		t5xx = template.Must(template.New("5xx").Parse("Template file error"))
	}

	return t5xx
}

func LoadTpl(path string) *template.Template {
	var err error
	funcMaps := template.FuncMap{
		"md": markDowner,
	}
	tc := template.New(path).Funcs(funcMaps)
	b, err := ioutil.ReadFile(path)
	if err != nil {
		log.Println("load tpl failed", err)
		return Load5xx()
	}
	s := string(b)

	tc, err = tc.Parse(s)
	if err != nil {
		log.Println("load tpl failed", err)
		return Load5xx()
	}
	return tc
}

func Render(w io.Writer, path string, data interface{}) {
	t := LoadTpl(path)
	t.ExecuteTemplate(w, path, data)
}

func markDowner(args ...interface{}) template.HTML {
	s := blackfriday.MarkdownCommon([]byte(fmt.Sprintf("%s", args...)))
	return template.HTML(s)
}
