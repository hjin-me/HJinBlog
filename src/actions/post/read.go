package actions

import "framework"

func Read(ctx fw.Context) {
	ctx.Output("hehe", "text/plain")
}
