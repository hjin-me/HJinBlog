package fw

import (
	"log"
	"net/http"
	"regexp"
	"runtime"
	"time"

	"golang.org/x/net/context"

	//  "log"
)

var routeList map[string][]routeInfo

func Init() {
	log.Println("route init")
	routeList = make(map[string][]routeInfo)
}

func parseRule(rule string) (*regexp.Regexp, []string, error) {
	nameList := []string{}
	// 提取字符 key
	re, err := regexp.Compile(":([^/]+)")
	if err != nil {
		log.Println(err)
		return re, nameList, err
	}
	tmpList := re.FindAllStringSubmatch(rule, -1)
	for _, v := range tmpList {
		log.Println(v)
		nameList = append(nameList, v[1])
	}
	log.Println(nameList)
	log.Println("rule " + rule)
	log.Println(tmpList)
	log.Println(re.ReplaceAllString(rule, "([^/]+)"))
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
	log.Print("before bind")
	log.Println("bind port")
	err := http.ListenAndServe(":"+port, mux) //设置监听的端口
	if err != nil {
		log.Print("error")
	}
	log.Println("after bind port")
	log.Println("after bind")
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
		log.Println(method)
		routeList[method] = []routeInfo{}
	}

	fn := func(ctx FwContext) {
		log.Println("file route")
		w := ctx.Res()
		r := ctx.Req()

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

	log.Println("start " + method)
	reg, nameList, err := parseRule(pattern)
	if err != nil {
		log.Fatal(err)
		return
	}
	log.Println(reg)
	rInfo := routeInfo{regex: reg, controller: fn, nameList: nameList}

	_, exist := routeList[method]
	if !exist {
		log.Println(method)
		routeList[method] = []routeInfo{}
	}

	routeList[method] = append(routeList[method], rInfo)
}

type routeInfo struct {
	regex      *regexp.Regexp
	controller ControllerType
	nameList   []string
}

type ControllerType func(ctx FwContext)

type controllerType func(http.ResponseWriter, *http.Request)

type CustomMux struct {
}

func (p *CustomMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println(routeList)
	log.Println("method" + r.Method)
	list, exist := routeList[r.Method]
	if !exist {
		http.NotFound(w, r)
		return
	}

	var ctx context.Context
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	for _, v := range list {
		log.Println("====")
		log.Println("url " + r.URL.Path)
		log.Println("regexp")
		log.Println(v.regex)
		log.Println(v.regex.FindStringSubmatch(r.URL.Path))
		res := v.regex.FindStringSubmatch(r.URL.Path)
		ctx = context.WithValue(ctx, "Req", r)
		ctx = context.WithValue(ctx, "Res", w)

		params := make(map[string]string)
		log.Println(v.nameList)
		log.Println(res)
		log.Println("====")
		for k, v := range v.nameList {
			if len(res) > k+1 {
				params[v] = res[k+1]
			} else {
				params[v] = ""
			}
		}
		ctx = context.WithValue(ctx, "Params", params)
		httpCtx, ok := ctx.(FwContext)
		if !ok {
			log.Fatal("convert Context to FwContext failed")
		}

		if len(res) > 0 {
			v.controller(httpCtx)
			break
		}
	}
	// log.Fprintf(w, "Method eq "+r.Method)
	return
}
