package fw

import (
	"html/template"
	"log"
)

type templateCache map[string]*template.Template

var tc = templateCache{}
var t5xx *template.Template

func Load5xx() *template.Template {
	if t5xx == nil {
		log.Println("load new 5xx tpl")
		t5xx = template.Must(template.New("5xx").Parse("Template file error"))
	}

	return t5xx
}

func LoadTpl(path string) *template.Template {
	if t, ok := tc[path]; ok {
		return t
	}
	t, err := template.ParseFiles(path)
	if err != nil {
		log.Println("load tpl failed", err)
		return Load5xx()
	}

	tc[path] = t

	return t
}
