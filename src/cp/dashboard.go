package cp

import (
	"net/http"

	"github.com/hjin-me/banana"
)

func DashBoard(ctx banana.Context) error {
	can, err := Auth(ctx, PrivilegePostRead)
	if err != nil {
		return err
	}
	if !can {
		http.Redirect(ctx.Res(), ctx.Req(), "/login?error", http.StatusFound)
		return nil
	}
	layout := ThemeLayout{}
	layout.Content = ThemeBlock{"cp:page/starter", 1}
	return ctx.Tpl("cp:page/layout", layout)
}
