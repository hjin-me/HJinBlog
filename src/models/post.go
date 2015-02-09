package models

import (
	"encoding/json"
	"html/template"
	"log"
	"strings"
	"time"

	"github.com/garyburd/redigo/redis"
)

type Post struct {
	Id          string
	Title       string
	Content     template.HTML
	Keywords    Keywords
	Description string
	PubTime     time.Time
}
type Keywords []Keyword

func (ks Keywords) String() string {
	tmps := make([]string, len(ks))
	for i, v := range ks {
		tmps[i] = string(v)
	}
	return strings.Join(tmps, ",")
}
func (kw Keywords) Marshal() string {
	ret, err := json.Marshal(kw)
	if err != nil {
		log.Fatalln(err)
	}
	return string(ret)
}
func (kw *Keywords) Unmarshal(str string) {
	err := json.Unmarshal([]byte(str), kw)
	if err != nil {
		log.Fatalln(err)
	}
}

type Keyword string

func (k Keyword) Alias() string {
	return strings.ToLower(string(k))
}

func (p Post) Save() {
	conn, err := redis.Dial("tcp", "10.211.55.8:6379")
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()
	conn.Do("hmset", p.Id, "id", p.Id, "title", p.Title, "content", string(p.Content),
		"keywords", p.Keywords.Marshal(), "description", p.Description, "pubtime", p.PubTime.Unix())
}

func Read(id string) Post {
	conn, err := redis.Dial("tcp", "10.211.55.8:6379")
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()
	reply, err := conn.Do("hmget", id, "id", "title", "content",
		"keywords", "description", "pubtime")
	if err != nil {
		log.Fatalln(err)
	}
	tarr, ok := reply.([]interface{})
	if !ok {
		log.Fatalln("convert to arr failed")
	}
	p := Post{}
	p.Id, _ = redis.String(tarr[0], nil)
	p.Title, _ = redis.String(tarr[1], nil)
	content, _ := redis.String(tarr[2], nil)
	p.Content = template.HTML(content)
	kw, _ := redis.String(tarr[3], nil)
	p.Keywords.Unmarshal(kw)
	p.Description, _ = redis.String(tarr[4], nil)
	t, _ := redis.Int64(tarr[5], nil)
	p.PubTime = time.Unix(t, 0)
	return p
}
