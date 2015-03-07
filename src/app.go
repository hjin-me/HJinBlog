package main

import (
	post "actions/post"
	"db"
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
	redisConf, ok := cfg.Env.Db["redis"]
	if !ok {
		log.Fatalln("redis conf not exits")
	}

	redisIp, ok := redisConf.(string)
	if !ok {
		log.Fatalln("redis ip error")
	}
	err := db.Connect(redisIp)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("database connected")
	// route init
	banana.Get("/post/:id", post.Read)
	banana.File("/statics", cfg.Env.Statics)
	banana.Get("/archives", post.Scan)
	banana.Get("/", post.Latest)

	<-ctx.Done()
}
