package actions

import (
	"models/post"
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
	x := post.Read(int(id))
	layout := ThemeLayout{}
	layout.Content = ThemeBlock{"my:page/post", x}
	return ctx.Tpl("my:page/layout", layout)
}

func Latest(ctx banana.Context) error {
	ps := post.Query(0, 10)
	layout := ThemeLayout{}
	layout.Content = ThemeBlock{"my:page/home", ps}
	return ctx.Tpl("my:page/layout", layout)
	/*
		posts, err := models.ZRange("pubtime", 0, 4)
		if err != nil {
			log.Fatal(err)
		}
		ctx.Tpl("home", posts)
	*/
}

type HomeLayout struct {
	Content interface{}
}

func Query(ctx banana.Context) error {
	posts := post.Query(0, 10)
	layout := HomeLayout{}
	layout.Content = posts
	return ctx.Tpl("my/home", layout)
}
