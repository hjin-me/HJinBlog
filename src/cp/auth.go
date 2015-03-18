package cp

import "github.com/hjin-me/banana"

func Auth(ctx banana.Context) {
	ctx.Req().Cookie("bnuid")
}
