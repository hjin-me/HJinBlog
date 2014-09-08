package main

import (
	"fmt"
	//	"log"
	"net/http"
	// "strings"
	"actions"
	"framework"
	"html/template"
	"os"
)

type Article struct {
	UserName string
	List     []int
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()       //解析参数，默认是不会解析的
	fmt.Println(r.Form) //这些信息是输出到服务器端的打印信息
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])

	// t := template.New("fieldname example")
	// t, _ = t.Parse("hello {{.UserName}}!")

	// dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	// if err != nil {
	//	log.Fatal(err)
	// }
	// fmt.Println(dir)

	// tpl := filepath.Join(dir, "./views/page.gtpl")

	viewBase := os.Getenv("GOPATH") + "/src/views"
	tpl := viewBase + "/page.gtpl"
	fmt.Println(tpl)
	texec := template.Must(template.ParseFiles(tpl))
	p := Article{UserName: "Songsong", List: []int{5, 2, 1, 9}}
	texec.Execute(w, p)

	//	for k, v := range r.Form {
	//		fmt.Println("key:", k)
	//		fmt.Println("val:", strings.Join(v, ""))
	//	}
	//	fmt.Fprintf(w, "Hello astaxie!") //这个写入到w的是输出到客户端的
}

func main() {
	fmt.Println("before Get")
	fw.Init()
	fw.Get("/install", actions.Init)
	fw.Get("/:aa", actions.Post)
	// fw.Get("/:abc/:er", sayHello)
	fmt.Println("after Get")
	fw.App()
}
