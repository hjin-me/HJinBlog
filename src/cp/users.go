package cp

import "github.com/hjin-me/banana"

func Users(ctx banana.Context) error {
	layout := ThemeLayout{}
	layout.Content = ThemeBlock{"cp:page/users", 1}
	return ctx.Tpl("cp:page/layout", layout)
}
