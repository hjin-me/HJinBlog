package cp

import (
	"models"
	"net/http"

	"github.com/hjin-me/banana"
)

func Posts(ctx banana.Context) error {

	can, err := Auth(ctx, PrivilegePostRead)
	if err != nil {
		return err
	}
	if !can {
		http.Redirect(ctx.Res(), ctx.Req(), "/login?error", http.StatusFound)
		return nil
	}

	ps := models.Query(0, 10)
	layout := ThemeLayout{}
	layout.Content = ThemeBlock{"cp:page/posts", ps}
	return ctx.Tpl("cp:page/layout", layout)
}
