package cp

import "github.com/hjin-me/banana"

func DashBoard(ctx banana.Context) error {
	layout := ThemeLayout{}
	layout.Content = ThemeBlock{"cp:page/starter", 1}
	return ctx.Tpl("cp:page/layout", layout)
}
