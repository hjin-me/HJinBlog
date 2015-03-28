package cp

import (
	"models"
	"strconv"

	"github.com/hjin-me/banana"
)

func Post(ctx banana.Context) error {
	var (
		idStr string
		ok    bool
	)
	if idStr, ok = ctx.Params()["id"]; !ok {
		panic("no id")
	}
	id, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		panic(err)
	}
	x := models.ReadRaw(int(id))

	layout := ThemeLayout{}
	layout.Content = ThemeBlock{"cp:page/post", x}
	return ctx.Tpl("cp:page/layout", layout)
}
