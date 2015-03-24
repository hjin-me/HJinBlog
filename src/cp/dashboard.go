package cp

import "github.com/hjin-me/banana"

func DashBoard(ctx banana.Context) {
	layout := PostLayout{}
	layout.ContentBlock = "cp:page/starter.html"
	ctx.Tpl("cp:page/layout.html", layout)
}
