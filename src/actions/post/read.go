package actions

import (
	"framework"
	"log"
	"models"
)

func Read(ctx fw.Context) {
	if id, ok := ctx.Params()["id"]; ok {
		x := models.Read(id)
		ctx.Tpl("post.html", x)
	}
}

func Scan(ctx fw.Context) {
	posts, err := models.Scan()
	if err != nil {
		log.Fatal(err)
	}
	ctx.Json(posts)
}
