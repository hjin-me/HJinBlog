package post

import (
	"encoding/json"
	"log"
	"strings"
)

type Keywords []Keyword

func (ks Keywords) String() string {
	tmps := make([]string, len(ks))
	for i, v := range ks {
		tmps[i] = string(v)
	}
	return strings.Join(tmps, ",")
}
func (ks *Keywords) Parse(s string) {
	s = strings.Trim(s, " ,")
	ss := strings.Split(s, ",")
	ts := make(Keywords, len(ss))
	for i, s := range ss {
		ts[i] = Keyword(s)
	}
	*ks = ts
}

func (kw Keywords) Marshal() string {
	ret, err := json.Marshal(kw)
	if err != nil {
		log.Println(err)
		panic(err)
	}
	return string(ret)
}
func (kw *Keywords) Unmarshal(str string) {
	err := json.Unmarshal([]byte(str), kw)
	if err != nil {
		log.Println(err)
		panic(err)
	}
}

type Keyword string

func (k Keyword) Alias() string {
	return strings.ToLower(string(k))
}
