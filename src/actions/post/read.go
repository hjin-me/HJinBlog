package actions

import "framework"

type Post struct {
	Title   string
	Context string
}

func Read(ctx fw.Context) {
	x := Post{"this is title", "this is content"}
	ctx.Tpl("test.tpl", x)
}
