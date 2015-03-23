package actions

import (
	"models"
	"strconv"

	"github.com/hjin-me/banana"
)

func Read(ctx banana.Context) {
	if idStr, ok := ctx.Params()["id"]; ok {
		id, err := strconv.ParseInt(idStr, 10, 32)
		if err != nil {
			panic(err)
		}
		x := models.Read(int(id))
		ctx.Tpl("my/post", x)
	}
}

func Latest(ctx banana.Context) {
	/*
		posts, err := models.ZRange("pubtime", 0, 4)
		if err != nil {
			log.Fatal(err)
		}
		ctx.Tpl("home.html", posts)
	*/
}

type HomeLayout struct {
	Content interface{}
}

func Query(ctx banana.Context) {
	posts := models.Query(0, 10)
	layout := HomeLayout{}
	layout.Content = posts
	ctx.Tpl("my/home", layout)
}
