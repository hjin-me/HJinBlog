package cp

import (
	"log"
	"models"
	"strconv"

	"github.com/hjin-me/banana"
)

func Post(ctx banana.Context) {
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
	log.Println(x)

	layout := PostLayout{}
	layout.Content = x
	ctx.Tpl("cp/post", layout)
}
