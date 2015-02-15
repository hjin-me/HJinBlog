package fw

import (
	"flag"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"time"

	"runtime"

	"golang.org/x/net/context"
	"gopkg.in/yaml.v2"
)

var routeList map[string][]routeInfo

type AppCfg struct {
	Env struct {
		ConfRoot string
		Port     string
		Level    string
		Tpl      string
		Db       string
	}
}

func loadCfg(filename string) (cfg AppCfg) {
	filename, err := filepath.Abs(filename)
	if err != nil {
		log.Fatalln("config file path error", err)
	}
	f, err := os.Open(filename)
	if err != nil {
		log.Fatalln("open config file failed", err)
	}
	defer f.Close()

	cfg.Env.ConfRoot = filepath.Dir(filename)

	bf, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatalln("read config file failed", err)
	}
	err = yaml.Unmarshal(bf, &cfg)
	if err != nil {
		log.Fatalln("load config fail", err)
	}

	cfg.Env.Tpl = filepath.Join(cfg.Env.ConfRoot, cfg.Env.Tpl)
	if err != nil {
		log.Fatalln("conf tpl is illegal", err)
	}

	fi, err := os.Lstat(cfg.Env.Tpl)
	if err != nil {
		log.Fatalln(err)
	}
	if !fi.IsDir() {
		log.Fatalf("%s: should be directory\n", cfg.Env.Tpl)
	}
	return

}

func parseRule(rule string) (*regexp.Regexp, []string, error) {
	nameList := []string{}
	// 提取字符 key
	re, err := regexp.Compile(":([^/]+)")
	if err != nil {
		log.Panic(err)
		return re, nameList, err
	}
	tmpList := re.FindAllStringSubmatch(rule, -1)
	for _, v := range tmpList {
		// log.Println(v)
		nameList = append(nameList, v[1])
	}
	////log.Println(nameList)
	////log.Println("rule " + rule)
	////log.Println(tmpList)
	////log.Println(re.ReplaceAllString(rule, "([^/]+)"))
	// 构造匹配用的正则
	ruleReg := re.ReplaceAllString(rule, "([^/]+)")
	ruleReg = "^" + ruleReg + "$"
	reg, err := regexp.Compile(ruleReg)
	if err != nil {
		return reg, nameList, err
	}
	return reg, nameList, nil
}

func initial() *MuxContext {
	runtime.GOMAXPROCS(runtime.NumCPU())
	routeList = make(map[string][]routeInfo)

	var confFilename string
	flag.StringVar(&confFilename, "c", "./app.yaml", "server configuration")
	flag.Parse()

	cfg := loadCfg(confFilename)

	return &MuxContext{context.WithValue(context.Background(), "cfg", cfg)}
}

func App() context.Context {
	ctx := initial()

	go func() {
		cfg, ok := ctx.Value("cfg").(AppCfg)
		if !ok {
			log.Print("configuration not ok")
			return
		}
		err := http.ListenAndServe(":"+cfg.Env.Port, ctx) //设置监听的端口
		if err != nil {
			log.Print(err)
		}
	}()
	return ctx
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
		routeList[method] = []routeInfo{}
	}

	fn := func(ctx Context) {
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

	reg, nameList, err := parseRule(pattern)
	if err != nil {
		log.Fatal(err)
		return
	}
	rInfo := routeInfo{regex: reg, controller: fn, nameList: nameList}

	_, exist := routeList[method]
	if !exist {
		routeList[method] = []routeInfo{}
	}

	routeList[method] = append(routeList[method], rInfo)
}

type routeInfo struct {
	regex      *regexp.Regexp
	controller ControllerType
	nameList   []string
}

type ControllerType func(ctx Context)

type controllerType func(http.ResponseWriter, *http.Request)

type MuxContext struct {
	context.Context
}

func (p *MuxContext) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	list, exist := routeList[r.Method]
	if !exist {
		http.NotFound(w, r)
		return
	}

	var ctx context.Context
	ctx, cancel := context.WithTimeout(p, 5*time.Second)
	defer cancel()

	for _, v := range list {
		res := v.regex.FindStringSubmatch(r.URL.Path)

		params := make(map[string]string)
		for k, v := range v.nameList {
			if len(res) > k+1 {
				params[v] = res[k+1]
			} else {
				params[v] = ""
			}
		}
		if len(res) > 0 {
			go func() {
				v.controller(WithHttp(ctx, w, r, params))
				cancel()
			}()
			break
		}
	}
	<-ctx.Done()
	return
}
