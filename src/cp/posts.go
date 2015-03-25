package cp

import (
	"models"

	"github.com/hjin-me/banana"
)

func Posts(ctx banana.Context) {
	ps := models.Query(0, 10)
	layout := ThemeLayout{}
	layout.Header.Name = "cp:page/header.html"
	layout.Sidebar.Name = "cp:page/sidebar.html"
	layout.Footer.Name = "cp:page/footer.html"
	layout.Content = ThemeBlock{"cp:page/posts.html", ps}
	ctx.Tpl("cp:page/layout.html", layout)
}
