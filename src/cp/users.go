package cp

import (
	"models/user"
	"net/http"
	"strconv"

	"github.com/hjin-me/banana"
)

func Users(ctx banana.Context) error {
	err := Auth(ctx, PrivilegeUserRead)
	switch err {
	case ErrNoPermit:
		return err
	case ErrNotLogin:
		http.Redirect(ctx.Res(), ctx.Req(), "/login?error&u=/cp/users", http.StatusFound)
		return nil
	case nil:
	default:
		return err
	}

	users, err := user.Query(0, 10)
	if err != nil {
		return err
	}
	layout := ThemeLayout{}
	layout.Content = ThemeBlock{"cp:page/users", struct{ List interface{} }{users}}
	return ctx.Tpl("cp:page/layout", layout)
}

func UsersCreatePage(ctx banana.Context) error {
	err := Auth(ctx, PrivilegeUserRead)
	switch err {
	case ErrNoPermit:
		return err
	case ErrNotLogin:
		http.Redirect(ctx.Res(), ctx.Req(), "/login?error&u=/cp/users", http.StatusFound)
		return nil
	case nil:
	default:
		return err
	}

	/*
		idStr, ok := ctx.Params()["id"]
		if !ok {
			http.Redirect(ctx.Res(), ctx.Req(), "/cp/users", http.StatusFound)
			return nil
		}
		id, err := strconv.ParseInt(idStr, 10, 32)
		if err != nil {
			http.Redirect(ctx.Res(), ctx.Req(), "/cp/users", http.StatusFound)
			return nil
		}
		u, err := user.FindOne(int(id))
		if err != nil {
			return err
		}
	*/
	u := 1
	layout := ThemeLayout{}
	layout.Content = ThemeBlock{"cp:page/user", struct{ Info interface{} }{u}}
	return ctx.Tpl("cp:page/layout", layout)

}

func UsersCreate(ctx banana.Context) error {
	r := ctx.Req()

	err := Auth(ctx, PrivilegeUserWrite)
	switch err {
	case ErrNoPermit:
		return err
	case ErrNotLogin:
		return err
	case nil:
	default:
		return err
	}

	username, pwd := r.FormValue("username"), r.FormValue("pwd")

	p, err := strconv.ParseInt(r.FormValue("privilege"), 10, 32)
	if err != nil {
		return err
	}
	privilege := int(p) & (PrivilegePostDelete | PrivilegePostDelete | PrivilegePostWrite | PrivilegeUserDelete | PrivilegeUserRead | PrivilegeUserWrite | PrivilegeCategoryRead | PrivilegeCategoryWrite | PrivilegeCategoryDelete)

	err = user.Add(username, pwd, privilege)
	if err != nil {
		return err
	}

	return ctx.Json(struct{}{})
}
