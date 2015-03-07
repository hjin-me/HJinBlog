package actions

import (
	"log"
	"models"

	"github.com/hjin-me/banana"
)

func Read(ctx banana.Context) {
	if id, ok := ctx.Params()["id"]; ok {
		x := models.Read(id)
		ctx.Tpl("post.html", x)
	}
}

func Latest(ctx banana.Context) {
	posts, err := models.ZRange("pubtime", 0, 4)
	if err != nil {
		log.Fatal(err)
	}
	ctx.Tpl("home.html", posts)
}
