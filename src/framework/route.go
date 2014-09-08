package fw

import (
	"fmt"
	"net/http"
	"regexp"
	"runtime"
	"strings"
	//  "log"
)

var routeList map[string][]routeInfo

func Init() {
	fmt.Println("route init")
	routeList = make(map[string][]routeInfo)
}

func parseRule(rule string) *regexp.Regexp {
	re, _ := regexp.Compile(":([^/]*)")
	fmt.Println(re.ReplaceAllString(rule, "([^/]+)"))
	ruleReg := re.ReplaceAllString(rule, "([^/]+)")
	ruleReg = "^" + ruleReg + "$"
	reg, _ := regexp.Compile(ruleReg)
	return reg
}

func App() {
	mux := &CustomMux{}
	fmt.Print("before bind")
	fmt.Println("bind port")
	err := http.ListenAndServe(":8080", mux) //设置监听的端口
	if err != nil {
		fmt.Print("error")
	}
	fmt.Println("after bind port")
	fmt.Println("after bind")
	runtime.Gosched()
	return
}

func Get(pattern string, fn controllerType) {
	Add("GET", pattern, fn)
}

func Post(pattern string, fn controllerType) {
	Add("POST", pattern, fn)
}
func Delete(pattern string, fn controllerType) {
	Add("DELETEl", pattern, fn)
}

func All(pattern string, fn controllerType) {
	Add("GET", pattern, fn)
	Add("POST", pattern, fn)
	Add("DELETE", pattern, fn)
	Add("PUT", pattern, fn)
	Add("OPTION", pattern, fn)
	Add("HEAD", pattern, fn)
}
func Add(method, pattern string, fn controllerType) {

	fmt.Println("start " + method)
	reg := parseRule(pattern)
	fmt.Println(reg)
	rInfo := routeInfo{regex: reg, controller: fn}

	_, exist := routeList[method]
	if !exist {
		routeList[method] = []routeInfo{}
	}

	routeList[method] = append(routeList[method], rInfo)
}

type routeInfo struct {
	regex      *regexp.Regexp
	controller controllerType
}

type controllerType func(http.ResponseWriter, *http.Request, []string)

type CustomMux struct {
}

func (p *CustomMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	list, exist := routeList[r.Method]
	if !exist {
		http.NotFound(w, r)
	}

	for k, v := range list {
		fmt.Println("url " + r.URL.Path)
		fmt.Println("key " + string(k))
		fmt.Println("match " + strings.Join(v.regex.FindStringSubmatch(r.URL.Path), " "))
		res := v.regex.FindStringSubmatch(r.URL.Path)
		var params = []string{}
		if len(res) > 1 {
			params = res[1:]
		}

		if len(res) > 0 {
			v.controller(w, r, params)
			break
		}
	}
	fmt.Fprintf(w, "Method eq "+r.Method)
	return
}
