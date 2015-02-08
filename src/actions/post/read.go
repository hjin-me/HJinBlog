package actions

import (
	"framework"
	"html/template"
	"strings"
	"time"
)

type Post struct {
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

type Keyword string

func (k Keyword) Alias() string {
	return strings.ToLower(string(k))
}

func Read(ctx fw.Context) {
	x := Post{}
	x.Title = "This is title"
	x.Content = template.HTML(`<p><img src="http://bcs.duapp.com/huangj-in/images/IMG_1289.jpg" alt="编写可维护的 JavaScript 封面"></p><p>这本书不是工具书，他不像《JavaScript 权威指南》可以用来当武器。在这本书里面不会告诉你 JavaScript 的语法是什么、闭包是什么、原型链是什么，相反地，她是以 JavaScript 为基础来介绍编程风格（缩进，注释，换行符等的使用）、保持代码可维护性的方法、以及面对不断增长的代码进行自动化测试部署等。</p>`)
	x.Description = "3月25日在http://ued.taobao.com/blog/2013/03/maintainable-javascript/得知了这本书的发售，3月27日购得此书。于是在第一时间阅读后，写下一些笔记，把这本书推荐给..."
	x.PubTime = time.Now()

	kws := []string{"可维护", "JavaScript", "权威指南", "书", "阅读", "前端"}
	x.Keywords = make([]Keyword, len(kws))
	for i, w := range kws {
		x.Keywords[i] = Keyword(w)
	}
	ctx.Tpl("post.html", x)
}
