package fw

import (
	"fmt"
	"log"
	"net/http"
	"regexp"
	"runtime"
	"strings"
	"time"

	"github.com/golang/net/context"
	//  "log"
)

var routeList map[string][]routeInfo

func Init() {
	fmt.Println("route init")
	routeList = make(map[string][]routeInfo)
}

func parseRule(rule string) (*regexp.Regexp, []string, error) {
	nameList := []string{}
	// 提取字符 key
	re, err := regexp.Compile(":([^/]+)")
	if err != nil {
		fmt.Println(err)
		return re, nameList, err
	}
	tmpList := re.FindAllStringSubmatch(rule, -1)
	for _, v := range tmpList {
		fmt.Println(v)
		nameList = append(nameList, v[1])
	}
	fmt.Println(nameList)
	fmt.Println("rule " + rule)
	fmt.Println(tmpList)
	fmt.Println(re.ReplaceAllString(rule, "([^/]+)"))
	// 构造匹配用的正则
	ruleReg := re.ReplaceAllString(rule, "([^/]+)")
	ruleReg = "^" + ruleReg + "$"
	reg, err := regexp.Compile(ruleReg)
	if err != nil {
		return reg, nameList, err
	}
	return reg, nameList, nil
}

func App(port string) {
	mux := &CustomMux{}
	fmt.Print("before bind")
	fmt.Println("bind port")
	err := http.ListenAndServe(":" + port, mux) //设置监听的端口
	if err != nil {
		fmt.Print("error")
	}
	fmt.Println("after bind port")
	fmt.Println("after bind")
	runtime.Gosched()
	return
}

func Put(pattern string, fn ControllerType) {
	add("PUT", pattern, fn)
}

func Get(pattern string, fn ControllerType) {
	add("GET", pattern, fn)
}

func Post(pattern string, fn ControllerType) {
	add("POST", pattern, fn)
}

func Delete(pattern string, fn ControllerType) {
	add("DELETE", pattern, fn)
}

func Option(pattern string, fn ControllerType) {
	add("OPTION", pattern, fn)
}

func All(pattern string, fn ControllerType) {
	add("GET", pattern, fn)
	add("POST", pattern, fn)
	add("DELETE", pattern, fn)
	add("PUT", pattern, fn)
	add("OPTION", pattern, fn)
	add("HEAD", pattern, fn)
}

func File(prefix string, dir string) {
	fsfn := http.StripPrefix(prefix, http.FileServer(http.Dir(dir))).ServeHTTP
	method := "GET"
	_, exist := routeList[method]
	if !exist {
		fmt.Println(method)
		routeList[method] = []routeInfo{}
	}

	fn := func(w http.ResponseWriter, r *http.Request, context Context) {
		fmt.Println("file route")
		fsfn(w, r)
	}

	nameList := []string{}
	// 提取字符 key
	ruleReg := "^" + prefix
	reg, err := regexp.Compile(ruleReg)
	if err != nil {
		return
	}
	rInfo := routeInfo{regex: reg, controller: fn, nameList: nameList}
	routeList[method] = append(routeList[method], rInfo)
}

func add(method, pattern string, fn ControllerType) {

	fmt.Println("start " + method)
	reg, nameList, err := parseRule(pattern)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println(reg)
	rInfo := routeInfo{regex: reg, controller: fn, nameList: nameList}

	_, exist := routeList[method]
	if !exist {
		fmt.Println(method)
		routeList[method] = []routeInfo{}
	}

	routeList[method] = append(routeList[method], rInfo)
}

type routeInfo struct {
	regex      *regexp.Regexp
	controller ControllerType
	nameList   []string
}

type ControllerType func(http.ResponseWriter, *http.Request, Context)

type controllerType func(http.ResponseWriter, *http.Request)

type CustomMux struct {
}

func (p *CustomMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println(routeList)
	fmt.Println("method" + r.Method)
	list, exist := routeList[r.Method]
	if !exist {
		http.NotFound(w, r)
		return
	}

	for _, v := range list {
		fmt.Println("====")
		fmt.Println("url " + r.URL.Path)
		fmt.Println("regexp")
		fmt.Println(v.regex)
		fmt.Println(v.regex.FindStringSubmatch(r.URL.Path))
		res := v.regex.FindStringSubmatch(r.URL.Path)
		var context Context
		context.Req = r
		context.Res = w
		context.Params = make(map[string]string)
		fmt.Println(v.nameList)
		fmt.Println(res)
		fmt.Println("====")
		for k, v := range v.nameList {
			if len(res) > k+1 {
				context.Params[v] = res[k+1]
			} else {
				context.Params[v] = ""
			}
		}

		if len(res) > 0 {
			v.controller(w, r, context)
			break
		}
	}
	// fmt.Fprintf(w, "Method eq "+r.Method)
	return
}
