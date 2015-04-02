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
	ErrNoPermit  = errors.New("have no permit")
	ErrPrivilege = errors.New("have no permit")
	ErrNotLogin  = errors.New("have not login")
)

func Auth(ctx banana.Context, p ...int) error {

	bnuid, err := ctx.Req().Cookie(UID_COOKIE_NAME)
	if err != nil {
		if err == http.ErrNoCookie {
			log.Println("cookes not found")
			return ErrNotLogin
		}
		return err
	}
	isLogin, username, err := user.DecodeToken(bnuid.Value)
	if err != nil {
		log.Println("decode error")
		return ErrNotLogin
	}
	if !isLogin {
		log.Println("is not login")
		return ErrNotLogin
	}
	privilege := 0
	for _, x := range p {
		privilege = privilege | x
	}

	can, err := user.Authentication(username, privilege)
	if err != nil {
		log.Println("auth error")
		return err
	}
	if !can {
		return ErrNoPermit
	}
	return nil
}

func LoginPage(ctx banana.Context) error {
	return ctx.Tpl("cp:page/login", 0)
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
