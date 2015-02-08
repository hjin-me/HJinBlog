package fw

import (
	"html/template"
	"io"
	"io/ioutil"
	"log"
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
	tc := template.New(path)
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
