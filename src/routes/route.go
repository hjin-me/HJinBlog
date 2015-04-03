package routes

import (
	"errors"

	"github.com/hjin-me/banana"
)

type RouteList []struct {
	Rule interface{}
	Func banana.ControllerType
}
type ActionList map[string]banana.ControllerType

var (
	actionList          = make(ActionList)
	ErrActionNameExists = errors.New("action name exists")
	ErrActionNotFound   = errors.New("action not found")
)

func Reg(name string, fn banana.ControllerType) error {
	if _, ok := actionList[name]; ok {
		return ErrActionNameExists
	}
	actionList[name] = fn
	return nil
}

func Handle() {
	rawConf := config{}
	_, err := banana.Config("routes.yaml", &rawConf)
	if err != nil {
		panic(err)
	}
	for _, r := range rawConf.Routes {
		for _, m := range r.Method {
			act, ok := actionList[r.Action]
			if !ok {
				panic(ErrActionNotFound)
			}
			switch m {
			case "POST":
				banana.Post(r.Rule, act)
			case "GET":
				banana.Get(r.Rule, act)
			case "PUT":
				banana.Put(r.Rule, act)
			case "DELETE":
				banana.Delete(r.Rule, act)
			case "OPTION":
				banana.Option(r.Rule, act)
			}
		}
	}
}
