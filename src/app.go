package main

import (
	post "actions/post"
	"cp"
	"da"
	"log"
	"routes"

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
	routes.Reg("post_page", post.Read)
	routes.Reg("archives_page", post.Query)
	routes.Reg("home_page", post.Latest)
	routes.Reg("category_page", post.Category)

	routes.Reg("login_page", cp.LoginPage)
	routes.Reg("login_post", cp.Login)

	routes.Reg("admin_users_page", cp.Users)
	routes.Reg("admin_user_page", cp.UsersCreatePage)
	routes.Reg("admin_user_post", cp.UsersCreate)

	routes.Reg("admin_posts_page", cp.Posts)
	routes.Reg("admin_post_new_page", cp.NewPost)
	routes.Reg("admin_post_edit_page", cp.Post)
	routes.Reg("admin_post_new_post", cp.SaveNewPost)
	routes.Reg("admin_post_edit_post", cp.SavePost)

	routes.Reg("admin_dashboard_page", cp.DashBoard)
	log.Println("route complete")

	banana.File("/static", cfg.Env.Statics)

	routes.Handle()
	// route init

	<-ctx.Done()
}
