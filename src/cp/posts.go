package cp

import "github.com/hjin-me/banana"

type PostLayout struct {
	ContentBlock string
	Content      interface{}
}

func Posts(ctx banana.Context) {
	layout := PostLayout{}
	layout.ContentBlock = "cp:page/bootstrap.html"
	layout.Content = 1
	ctx.Tpl("cp:page/layout.html", layout)
}
