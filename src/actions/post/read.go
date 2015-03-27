package actions

import (
	"models"
	"strconv"

	"github.com/hjin-me/banana"
)

func Read(ctx banana.Context) error {
	var (
		idStr string
		ok    bool
	)
	if idStr, ok = ctx.Params()["id"]; !ok {
		panic("no id")
	}
	id, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		panic(err)
	}
	x := models.Read(int(id))
	layout := ThemeLayout{}
	layout.Content = ThemeBlock{"my:page/post.html", x}
	return ctx.Tpl("my:page/layout.html", layout)
}

func Latest(ctx banana.Context) error {
	ps := models.Query(0, 10)
	layout := ThemeLayout{}
	layout.Content = ThemeBlock{"my:page/home.html", ps}
	return ctx.Tpl("my:page/layout.html", layout)
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

func Query(ctx banana.Context) error {
	posts := models.Query(0, 10)
	layout := HomeLayout{}
	layout.Content = posts
	return ctx.Tpl("my/home", layout)
}
