package models

import (
	"encoding/json"
	"html/template"
	"log"
	"strings"
	"time"

	"da"
	"database/sql"
)

type Post struct {
	Id          int
	Title       string
	Content     template.HTML
	Category    string
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

func (p *Post) Save() {
	db, err := da.Connect()
	if err != nil {
		panic(err)
	}
	var (
		stmt     *sql.Stmt
		isInsert = false
	)
	if p.Id == 0 {
		stmt, err = db.Prepare("INSERT INTO blog_posts (content, category, pubtime, title, description, tags) VALUES (?,?,?,?,?,?)") // ? = placeholder
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		isInsert = true
		defer stmt.Close() // Close the statement when we leave main() / the program terminates
	} else {
		stmt, err = db.Prepare("UPDATE blog_posts SET content=?, category=?, pubtime=?, title=?, description=?, tags=? WHERE id=?") // ? = placeholder
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		defer stmt.Close() // Close the statement when we leave main() / the program terminates
	}

	result, err := stmt.Exec(string(p.Content), p.Category, p.PubTime, p.Title, p.Description, p.Keywords.String(), p.Id)
	if err != nil {
		panic(err)
	}
	if isInsert {
		id, err := result.LastInsertId()
		if err != nil {
			panic(err)
		}
		p.Id = int(id)
	}

	// Prepare statement for inserting data
	// db.Do("hmset", "post:"+p.Id, "id", p.Id, "title", p.Title, "content", string(p.Content), "keywords", p.Keywords.Marshal(), "description", p.Description, "pubtime", p.PubTime.Unix(), "category", p.Category)
}

func Read(id int) Post {
	var (
		p   Post
		err error
	)
	db, err := da.Connect()
	if err != nil {
		panic(err)
	}
	// Prepare statement for reading data
	stmt, err := db.Prepare("SELECT id, content, category, pubtime, title, description, tags FROM blog_posts WHERE id = ?")
	if err != nil {
		panic(err) // proper error handling instead of panic in your app
	}
	defer stmt.Close()

	var (
		content     string
		category    string
		pubtime     int64
		title       string
		description string
		tags        string
	)
	err = stmt.QueryRow(id).Scan(&id, &content, &category, &pubtime, &title, &description, &tags)
	if err != nil {
		panic(err)
	}
	p.Id = id
	p.Content = template.HTML(content)
	p.Category = category
	p.PubTime = time.Unix(pubtime, 0)
	p.Title = title
	p.Description = description

	kws := strings.Split(tags, ",")
	for _, v := range kws {
		v = strings.Trim(v, " ")
		if v != "" {
			p.Keywords = append(p.Keywords, Keyword(v))
		}
	}

	return p
}

func New() Post {
	return Post{}
}

func Query(start, limit int) []Post {

	var (
		err error
		ps  []Post
	)
	db, err := da.Connect()
	if err != nil {
		panic(err)
	}
	// Prepare statement for reading data
	stmt, err := db.Prepare("SELECT id, content, category, pubtime, title, description, tags FROM blog_posts ORDER BY pubtime DESC LIMIT ?,?")
	if err != nil {
		panic(err) // proper error handling instead of panic in your app
	}
	defer stmt.Close()

	var (
		id          int
		content     string
		category    string
		pubtime     int64
		title       string
		description string
		tags        string
	)
	rows, err := stmt.Query(start, limit)
	for rows.Next() {
		err = rows.Scan(&id, &content, &category, &pubtime, &title, &description, &tags)
		if err != nil {
			panic(err)
		}
		var (
			p Post
		)
		p.Id = id
		p.Content = template.HTML(content)
		p.Category = category
		p.PubTime = time.Unix(pubtime, 0)
		p.Title = title
		p.Description = description

		kws := strings.Split(tags, ",")
		for _, v := range kws {
			v = strings.Trim(v, " ")
			if v != "" {
				p.Keywords = append(p.Keywords, Keyword(v))
			}
		}
		ps = append(ps, p)
	}

	return ps
}
