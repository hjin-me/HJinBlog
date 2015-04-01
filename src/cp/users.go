package cp

import (
	"net/http"

	"github.com/hjin-me/banana"
)

func Users(ctx banana.Context) error {
	err := Auth(ctx, PrivilegeUserRead)
	switch err {
	case ErrNoPermit:
		return err
	case ErrNotLogin:
		http.Redirect(ctx.Res(), ctx.Req(), "/login?error&u=/cp/users", http.StatusFound)
		return nil
	case nil:
	default:
		return err
	}
	layout := ThemeLayout{}
	layout.Content = ThemeBlock{"cp:page/users", 1}
	return ctx.Tpl("cp:page/layout", layout)
}
