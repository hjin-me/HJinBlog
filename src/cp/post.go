package cp

import (
	"models"
	"net/http"
	"strconv"

	"github.com/hjin-me/banana"
)

func Post(ctx banana.Context) error {
	err := Auth(ctx, PrivilegePostRead)
	switch err {
	case ErrNoPermit:
		return err
	case ErrNotLogin:
		http.Redirect(ctx.Res(), ctx.Req(), "/login?error", http.StatusFound)
		return nil
	case nil:
	default:
		return err
	}

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
	post := models.ReadRaw(int(id))
	categories, err := models.QueryCategory()
	if err != nil {
		return err
	}

	layout := ThemeLayout{}
	layout.Content = ThemeBlock{"cp:page/post", struct{ Post, Categories interface{} }{post, categories}}
	return ctx.Tpl("cp:page/layout", layout)
}

func SavePost(ctx banana.Context) error {
	err := Auth(ctx, PrivilegePostWrite)
	switch err {
	case ErrNoPermit:
		return err
	case ErrNotLogin:
		return err
	case nil:
	default:
		return err
	}

	var (
		idStr string
		ok    bool
	)
	if idStr, ok = ctx.Params()["id"]; !ok {
		panic("no id")
	}
	id, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		return err
	}

	r := ctx.Req()
	p := models.New()
	p.Id = int(id)
	p.Title = r.FormValue("title")
	p.Content = r.FormValue("content")
	p.Category = r.FormValue("category")
	p.Description = r.FormValue("description")
	p.Keywords.Parse(r.FormValue("keywords"))
	err = p.Save()
	if err != nil {
		return err
	}

	return ctx.Json(p)

}
