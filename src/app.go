package main

import (
	post "actions/post"
	"cp"
	"da"
	"log"

	"github.com/hjin-me/banana"
)

func main() {
	ctx := banana.App()
	log.Println("server started")
	cfg, ok := ctx.Value("cfg").(banana.AppCfg)
	if !ok {
		log.Fatalln("configuration not ok")
	}
	log.Println(cfg)
	dsnRaw, ok := cfg.Env.Db["mysql"]
	if !ok {
		log.Fatalln("mysql conf not exits")
	}

	dsn, ok := dsnRaw.(string)
	if !ok {
		log.Fatalln("mysql dsn not ok")
	}
	err := da.Create(dsn)
	if err != nil {
		log.Fatalln(err)
	}
	defer da.Close()
	log.Println("database connected")
	// route init
	banana.Get("/post/:id", post.Read)
	banana.File("/static", cfg.Env.Statics)
	banana.Get("/archives", post.Query)
	banana.Get("/cp/dashboard", cp.DashBoard)
	banana.Get("/cp/users", cp.DashBoard)
	banana.Get("/cp/posts", cp.Posts)
	// banana.Get("/", post.Latest)

	<-ctx.Done()
}
