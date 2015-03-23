package cp

import (
	"models"

	"github.com/hjin-me/banana"
)

type PostLayout struct {
	Content interface{}
}

func Posts(ctx banana.Context) {
	posts := models.Query(0, 10)
	layout := PostLayout{}
	layout.Content = posts
	ctx.Tpl("cp/posts", layout)
}
