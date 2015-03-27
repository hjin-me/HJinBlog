package cp

import "github.com/hjin-me/banana"

func Users(ctx banana.Context) error {
	layout := ThemeLayout{}
	layout.Header.Name = "cp:page/header.html"
	layout.Sidebar.Name = "cp:page/sidebar.html"
	layout.Footer.Name = "cp:page/footer.html"
	layout.Content = ThemeBlock{"cp:page/users.html", 1}
	return ctx.Tpl("cp:page/layout.html", layout)
}
