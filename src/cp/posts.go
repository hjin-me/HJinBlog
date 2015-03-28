package cp

import (
	"models"

	"github.com/hjin-me/banana"
)

func Posts(ctx banana.Context) error {
	ps := models.Query(0, 10)
	layout := ThemeLayout{}
	layout.Content = ThemeBlock{"cp:page/posts", ps}
	return ctx.Tpl("cp:page/layout", layout)
}
