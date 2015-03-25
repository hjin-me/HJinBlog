package cp

import "github.com/hjin-me/banana"

type PostLayout struct {
	ContentBlock string
	Content      interface{}
}

func Posts(ctx banana.Context) {
	layout := ThemeLayout{}
	layout.Header.Name = "cp:page/header.html"
	layout.Sidebar.Name = "cp:page/sidebar.html"
	layout.Footer.Name = "cp:page/footer.html"
	layout.Content = ThemeBlock{"cp:page/bootstrap.html", 1}
	ctx.Tpl("cp:page/layout.html", layout)
}
