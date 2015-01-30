package main

import (
<<<<<<< HEAD
	"actions"
	"framework"
	"os"
)

func main() {
	fw.Init()
	fw.Post("/cp", actions.Create)
	fw.Get("/cp", actions.ReadBatch)
	fw.Get("/cp/:id", actions.Read)
	fw.Put("/cp/:id", actions.Update)
	fw.Delete("/cp/:id", actions.Delete)
	fw.Get("/install", actions.LevelInit)
	fw.Get("/:prefix/:id", actions.PostRead)
	fw.Get("/", actions.PostReadBatch)
	fw.File("/statics", os.Getenv("GOPATH")+"/src/statics")
	// fw.Get("/:abc/:er", sayHello)
	fw.App()
=======
	// "actions"
	"framework"
	"os"
    "log"
    "fmt"
    "gopkg.in/yaml.v2"
    "io/ioutil"
    "path/filepath"
    "net/http"
//     "runtime"
)

type ServerConfig struct {
    Env struct {
        Port string
        Statics string
    }
}

func main() {
    defer func() {
        if r := recover(); r != nil {
            log.Fatalf("Recovered in main %v\n", r)
        }
    }()

    cfg := loadCfg()

	fw.Init()

    if cfg.Env.Statics != "" {
        s, err :=filepath.Abs(cfg.Env.Statics) 
        if err != nil {
            log.Fatal(err)
        }
        cfg.Env.Statics = s
        log.Println("statics dir is ", cfg.Env.Statics)
	    fw.File("/statics", cfg.Env.Statics)
    }
    // route init
    fw.Get("/a", func(res http.ResponseWriter, req *http.Request, context fw.Context) {
    })
	fw.App(cfg.Env.Port)

}

func loadCfg() (cfg ServerConfig) {
    defer func() {
        if r := recover(); r != nil {
            log.Fatal(r)
        }
    }()
    // load cfg
    confDir := "" 
    if len(os.Args) > 1 {
        var err error
        confDir, err = filepath.Abs(os.Args[1])
        if err != nil {
            panic(err)
        }
    } else {
        log.Fatal("you need set config param")
        return
    }

    data, err := ioutil.ReadFile( filepath.Join(confDir, "app.yaml"))
    if err != nil {
        panic(err)
    }
    err = yaml.Unmarshal(data, &cfg)
    if err != nil {
        panic(err)
        log.Fatalf("error: %v", err)
    }
    fmt.Println(cfg)

    // ircase.LoadCfg(confDir)

    return
>>>>>>> FETCH_HEAD
}
