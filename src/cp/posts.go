package cp

import (
	"models"
	"net/http"

	"github.com/hjin-me/banana"
)

func Posts(ctx banana.Context) error {

	err := Auth(ctx, PrivilegePostRead)
	switch err {
	case ErrNoPermit:
		return err
	case ErrNotLogin:
		http.Redirect(ctx.Res(), ctx.Req(), "/login?error&u=/cp/posts", http.StatusFound)
		return nil
	case nil:
	default:
		return err
	}

	ps := models.Query(0, 10)
	layout := ThemeLayout{}
	layout.Content = ThemeBlock{"cp:page/posts", ps}
	return ctx.Tpl("cp:page/layout", layout)
}
