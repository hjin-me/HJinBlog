package main

import (
	post "actions/post"
	"db"
	"framework"
	"log"
)

func main() {
	ctx := fw.App()
	log.Println("server started")
	cfg, ok := ctx.Value("cfg").(fw.AppCfg)
	if !ok {
		log.Fatalln("configuration not ok")
	}
	err := db.Connect(cfg.Env.Db)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("database connected")
	// route init
	fw.Get("/post/:id", post.Read)
	fw.File("/statics", cfg.Env.Statics)
	fw.Get("/archives", post.Scan)
	fw.Get("/", post.Latest)

	<-ctx.Done()
}
