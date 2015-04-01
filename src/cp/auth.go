package cp

import (
	"errors"
	"log"
	"models/user"
	"net/http"
	"time"

	"github.com/hjin-me/banana"
)

const (
	UID_COOKIE_NAME = "bnuid"
)
const (
	PrivilegePostRead int = 1 << iota
	PrivilegePostWrite
	PrivilegePostDelete
	PrivilegeUserRead
	PrivilegeUserWrite
	PrivilegeUserDelete
	PrivilegeCategoryRead
	PrivilegeCategoryWrite
	PrivilegeCategoryDelete
)

var (
	ErrPrivilege = errors.New("have no permit")
	ErrNotLogin  = errors.New("have not login")
)

func Auth(ctx banana.Context, p ...int) (bool, error) {

	bnuid, err := ctx.Req().Cookie(UID_COOKIE_NAME)
	if err != nil {
		if err == http.ErrNoCookie {
			log.Println("cookes not found")
			return false, nil
		}
		return false, err
	}
	isLogin, username, err := user.DecodeToken(bnuid.Value)
	if err != nil {
		log.Println("decode error")
		return false, err
	}
	if !isLogin {
		log.Println("is not login")
		return false, nil
	}
	privilege := 0
	for _, x := range p {
		privilege = privilege | x
	}
	log.Println(privilege)

	can, err := user.Authentication(username, privilege)
	if err != nil {
		log.Println("auth error")
		return false, err
	}
	return can, nil
}

func LoginPage(ctx banana.Context) error {
	return ctx.Tpl("cp:page/login", 0)
}

func Create(ctx banana.Context) error {

	r := ctx.Req()

	can, err := Auth(ctx, PrivilegeUserWrite)
	if err != nil {
		return err
	}
	if !can {
		return ErrPrivilege
	}

	username, pwd := r.FormValue("username"), r.FormValue("pwd")

	privilege := 0
	switch r.FormValue("privilege") {
	case "admin":
		privilege = PrivilegePostDelete | PrivilegePostDelete | PrivilegePostWrite | PrivilegeUserDelete | PrivilegeUserRead | PrivilegeUserWrite | PrivilegeCategoryRead | PrivilegeCategoryWrite | PrivilegeCategoryDelete
	case "editor":
		privilege = PrivilegePostDelete | PrivilegePostDelete | PrivilegePostWrite | PrivilegeCategoryRead | PrivilegeCategoryWrite | PrivilegeCategoryDelete
	case "visitor":
		privilege = PrivilegePostRead
	}

	err = user.Add(username, pwd, privilege)
	if err != nil {
		return err
	}

	return ctx.Json(struct{}{})
}

func Login(ctx banana.Context) error {

	r := ctx.Req()
	username, pwd := r.FormValue("username"), r.FormValue("pwd")
	result, sign, err := user.Check(username, pwd)
	if err != nil {
		return err
	}
	if result {
		timeout := time.Now().Add(user.Expires)
		userCookie := &http.Cookie{}
		userCookie.Expires = timeout
		userCookie.Name = UID_COOKIE_NAME
		userCookie.Value = sign
		http.SetCookie(ctx.Res(), userCookie)
		http.Redirect(ctx.Res(), ctx.Req(), "/cp/dashboard", http.StatusFound)
	} else {
		http.Redirect(ctx.Res(), ctx.Req(), "/login?error", http.StatusFound)
	}

	return nil
}
