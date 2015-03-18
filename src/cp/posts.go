package cp

import "github.com/hjin-me/banana"

func Posts(ctx banana.Context) {
	ctx.Tpl("cp/posts", nil)
}
