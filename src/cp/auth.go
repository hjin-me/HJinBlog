package cp

import "github.com/hjin-me/banana"

func Auth(ctx banana.Context) {
	ctx.Req().Cookie("bnuid")
}

func Login(ctx banana.Context) error {
	return ctx.Tpl("cp:page/login", 0)
}
