package cp

import (
	"net/http"

	"github.com/hjin-me/banana"
)

func Users(ctx banana.Context) error {
	can, err := Auth(ctx, PrivilegeUserRead)
	if err != nil {
		return err
	}
	if !can {
		http.Redirect(ctx.Res(), ctx.Req(), "/login?error", http.StatusFound)
		return nil
	}
	layout := ThemeLayout{}
	layout.Content = ThemeBlock{"cp:page/users", 1}
	return ctx.Tpl("cp:page/layout", layout)
}
