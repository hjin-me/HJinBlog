package actions

import (
	"fmt"
	"net/http"

	"code.google.com/p/go-uuid/uuid"
)

func Init(w http.ResponseWriter, r *http.Request, params []string) {

	u := uuid.New()
	fmt.Println(u)
}
