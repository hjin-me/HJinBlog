package actions

import (
	"framework"
	"models"
)

func Read(ctx fw.Context) {
	if id, ok := ctx.Params()["id"]; ok {
		x := models.Read(id)
		ctx.Tpl("post.html", x)
	}
}
