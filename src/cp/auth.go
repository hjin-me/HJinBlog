package cp

import (
	"models"
	"net/http"

	"github.com/hjin-me/banana"
)

func Auth(ctx banana.Context) {
	ctx.Req().Cookie("bnuid")
}

func LoginPage(ctx banana.Context) error {
	return ctx.Tpl("cp:page/login", 0)
}

func Login(ctx banana.Context) error {

	r := ctx.Req()
	username, pwd := r.FormValue("username"), r.FormValue("pwd")
	result, err := models.UserCheck(username, pwd)
	if err != nil {
		return err
	}
	if result {
		http.Redirect(ctx.Res(), ctx.Req(), "/cp/dashboard", http.StatusFound)
	} else {
		http.Redirect(ctx.Res(), ctx.Req(), "/login?error", http.StatusFound)
	}

	return nil

}
