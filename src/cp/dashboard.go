package cp

import "github.com/hjin-me/banana"

func DashBoard(ctx banana.Context) {
	ctx.Tpl("cp/dashboard", nil)
}
