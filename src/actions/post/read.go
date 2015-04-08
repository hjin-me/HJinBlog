package actions

import (
	"errors"
	"models/category"
	"models/post"
	"strconv"
	"theme"

	"github.com/hjin-me/banana"
)

func Read(ctx banana.Context) error {
	var (
		idStr string
		ok    bool
	)
	if idStr, ok = ctx.Params()["id"]; !ok {
		return errors.New("no id")
	}
	id, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		return err
	}
	x := post.Read(int(id))
	if yearStr, ok := ctx.Params()["year"]; ok {
		t, err := strconv.ParseInt(yearStr, 10, 32)
		if err != nil {
			return err
		}
		if int(t) != x.PubTime.Year() {
			return errors.New("404")
		}
	}
	if monthStr, ok := ctx.Params()["month"]; ok {
		t, err := strconv.ParseInt(monthStr, 10, 32)
		if err != nil {
			return err
		}
		if int(t) != int(x.PubTime.Month()) {
			return errors.New("404")
		}
	}
	if dayStr, ok := ctx.Params()["day"]; ok {
		t, err := strconv.ParseInt(dayStr, 10, 32)
		if err != nil {
			return err
		}
		if int(t) != x.PubTime.Day() {
			return errors.New("404")
		}
	}

	layout := ThemeLayout{}
	layout.Content = ThemeBlock{theme.UI("post"), x}
	return ctx.Tpl(theme.UI("layout"), layout)
}

func Latest(ctx banana.Context) error {
	ps := post.Query(0, 5)
	layout := ThemeLayout{}
	layout.Content = ThemeBlock{theme.UI("home"), ps}
	return ctx.Tpl(theme.UI("layout"), layout)
	/*
		posts, err := models.ZRange("pubtime", 0, 4)
		if err != nil {
			log.Fatal(err)
		}
		ctx.Tpl("home", posts)
	*/
}

func Category(ctx banana.Context) error {
	c, ok := ctx.Params()["category"]
	if !ok {
		return errors.New("no category")

	}
	ps, err := post.QueryByCategory(c, 0, 10)
	if err != nil {
		return err
	}
	var ca = category.Category{}
	if len(ps) > 0 {
		ca = ps[0].Category
	}

	layout := ThemeLayout{}
	layout.Content = ThemeBlock{theme.UI("category"), struct {
		List  []post.Post
		Bread category.Category
	}{ps, ca}}
	return ctx.Tpl(theme.UI("layout"), layout)
}

type HomeLayout struct {
	Content interface{}
}

func Query(ctx banana.Context) error {
	posts := post.Query(0, 10)
	layout := HomeLayout{}
	layout.Content = posts
	return ctx.Tpl(theme.UI("home"), layout)
}
