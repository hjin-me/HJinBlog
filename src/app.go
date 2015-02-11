package main

import (
	post "actions/post"
	"db"
	"framework"
	"log"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			log.Fatalf("Recovered in main %v\n", r)
		}
	}()

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

	<-ctx.Done()
}
