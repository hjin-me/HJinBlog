package actions

import (
	"framework"
	"models"
)

func Read(ctx fw.Context) {
	x := models.Read("761")
	ctx.Tpl("post.html", x)
}
