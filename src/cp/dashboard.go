package cp

import (
	"net/http"

	"github.com/hjin-me/banana"
)

func DashBoard(ctx banana.Context) error {
	err := Auth(ctx, PrivilegePostRead)
	switch err {
	case ErrNoPermit:
		http.Redirect(ctx.Res(), ctx.Req(), "/login?error&u=/cp/dashboard", http.StatusFound)
		return err
	case ErrNotLogin:
		http.Redirect(ctx.Res(), ctx.Req(), "/login?error&u=/cp/dashboard", http.StatusFound)
		return err
	case nil:
	default:
		return err
	}

	layout := ThemeLayout{}
	layout.Content = ThemeBlock{"cp:page/starter", 1}
	return ctx.Tpl("cp:page/layout", layout)
}
