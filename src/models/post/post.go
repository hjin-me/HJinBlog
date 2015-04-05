package post

import (
	"html/template"
	"models/category"
	"time"

	"da"
	"database/sql"
)

const (
	TABLE_NAME_POSTS      = "blog_posts"
	TABLE_NAME_CATEGORIES = "blog_categories"
	INSERT_POSTS          = "INSERT INTO " + TABLE_NAME_POSTS + "(uid, content, category, pubtime, title, description, keywords) VALUES (?,?,?,?,?,?,?)"
	UPDATE_POSTS          = "UPDATE " + TABLE_NAME_POSTS + " SET uid=?, content=?, category=?, pubtime=?, title=?, description=?, keywords=? WHERE id=?"
	QUERY_POSTS           = "SELECT p.id, p.uid, p.content, p.cid, c.name, c.alias, p.pubtime, p.title, p.description, p.keywords FROM " + TABLE_NAME_POSTS + " as p LEFT JOIN " + TABLE_NAME_CATEGORIES + " as c ON c.id=p.cid ORDER BY pubtime DESC LIMIT ?,?"
	FIND_POSTS            = "SELECT p.id, p.uid, p.content, p.cid, c.name, c.alias, p.pubtime, p.title, p.description, p.keywords FROM " + TABLE_NAME_POSTS + " as p LEFT JOIN " + TABLE_NAME_CATEGORIES + " as c ON c.id=p.cid WHERE id = ?"
)

type RawPost struct {
	Id          int
	UserId      int
	Title       string
	Content     string
	Category    category.Category
	Keywords    Keywords
	Description string
	PubTime     time.Time
}

type Post struct {
	RawPost
	Content template.HTML
}

func (p *RawPost) Save() error {
	db, err := da.Connect()
	if err != nil {
		return err
	}
	var (
		stmt     *sql.Stmt
		isInsert = false
	)
	if p.Id == 0 {
		stmt, err = db.Prepare(INSERT_POSTS) // ? = placeholder
		if err != nil {
			return err
		}
		isInsert = true
		defer stmt.Close() // Close the statement when we leave main() / the program terminates
	} else {
		stmt, err = db.Prepare(UPDATE_POSTS) // ? = placeholder
		if err != nil {
			return err
		}
		defer stmt.Close() // Close the statement when we leave main() / the program terminates
	}

	result, err := stmt.Exec(p.UserId, string(p.Content), p.Category, p.PubTime.Unix(), p.Title, p.Description, p.Keywords.String(), p.Id)
	if err != nil {
		return err
	}
	if isInsert {
		id, err := result.LastInsertId()
		if err != nil {
			return err
		}
		p.Id = int(id)
	}
	return nil

	// Prepare statement for inserting data
	// db.Do("hmset", "post:"+p.Id, "id", p.Id, "title", p.Title, "content", string(p.Content), "keywords", p.Keywords.Marshal(), "description", p.Description, "pubtime", p.PubTime.Unix(), "category", p.Category)
}

func ReadRaw(id int) RawPost {
	var (
		p   RawPost
		err error
	)
	db, err := da.Connect()
	if err != nil {
		panic(err)
	}
	// Prepare statement for reading data
	stmt, err := db.Prepare(FIND_POSTS)
	if err != nil {
		panic(err) // proper error handling instead of panic in your app
	}
	defer stmt.Close()

	var (
		uid         int
		content     string
		category    category.Category
		pubtime     int64
		title       string
		description string
		keywords    string
	)
	// "SELECT p.id, p.uid, p.content, p.cid, c.name, c.alias, p.pubtime, p.title, p.description, p.keywords FROM " + TABLE_NAME_POSTS + " as p LEFT JOIN " + TABLE_NAME_CATEGORIES + " as c ON c.id=p.cid WHERE id = ?"
	err = stmt.QueryRow(id).Scan(&id, &uid, &content, &category.Id, &category.Name, &category.Alias, &pubtime, &title, &description, &keywords)
	if err != nil {
		panic(err)
	}
	p.Id = id
	p.Content = content
	p.Category = category
	p.PubTime = time.Unix(pubtime, 0)
	p.Title = title
	p.Description = description

	p.Keywords.Parse(keywords)

	return p
}

func Read(id int) Post {
	pr := ReadRaw(id)

	p := Post{pr, template.HTML(pr.Content)}

	return p
}

func New() RawPost {
	p := RawPost{}
	p.PubTime = time.Now()
	return p
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
	stmt, err := db.Prepare(QUERY_POSTS)
	if err != nil {
		panic(err) // proper error handling instead of panic in your app
	}
	defer stmt.Close()

	var (
		id          int
		uid         int
		content     string
		category    category.Category
		pubtime     int64
		title       string
		description string
		keywords    string
	)
	rows, err := stmt.Query(start, limit)
	for rows.Next() {
		// "SELECT p.id, p.uid, p.content, p.cid, c.name, c.alias, p.pubtime, p.title, p.description, p.keywords FROM " + TABLE_NAME_POSTS + " as p LEFT JOIN " + TABLE_NAME_CATEGORIES + " as c ON c.id=p.cid WHERE id = ?"
		err = rows.Scan(&id, &uid, &content, &category.Id, &category.Name, &category.Alias, &pubtime, &title, &description, &keywords)
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
		p.Keywords.Parse(keywords)

		ps = append(ps, p)
	}

	return ps
}
