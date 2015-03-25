package cp

import "github.com/hjin-me/banana"

func DashBoard(ctx banana.Context) {
	layout := ThemeLayout{}
	layout.Header.Name = "cp:page/header.html"
	layout.Sidebar.Name = "cp:page/sidebar.html"
	layout.Footer.Name = "cp:page/footer.html"
	layout.Content = ThemeBlock{"cp:page/starter.html", 1}
	ctx.Tpl("cp:page/layout.html", layout)
}
