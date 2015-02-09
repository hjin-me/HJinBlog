package actions

import (
	"html/template"
	"log"
	"models"
	"testing"
	"time"
)

func TestSavePost(t *testing.T) {

	x := models.Post{}
	x.Title = "This is title"
	x.Id = "761"
	x.Content = template.HTML(`<p><img src="http://bcs.duapp.com/huangj-in/images/IMG_1289.jpg" alt="编写可维护的 JavaScript 封面"></p><p>这本书不是工具书，他不像《JavaScript 权威指南》可以用来当武器。在这本书里面不会告诉你 JavaScript 的语法是什么、闭包是什么、原型链是什么，相反地，她是以 JavaScript 为基础来介绍编程风格（缩进，注释，换行符等的使用）、保持代码可维护性的方法、以及面对不断增长的代码进行自动化测试部署等。</p>`)
	x.Description = "3月25日在http://ued.taobao.com/blog/2013/03/maintainable-javascript/得知了这本书的发售，3月27日购得此书。于是在第一时间阅读后，写下一些笔记，把这本书推荐给..."
	x.PubTime = time.Now()

	kws := []string{"可维护", "JavaScript", "权威指南", "书", "阅读", "前端"}
	x.Keywords = make([]models.Keyword, len(kws))
	for i, w := range kws {
		x.Keywords[i] = models.Keyword(w)
	}

	x.Save()
}

func TestReadPost(t *testing.T) {
	x := models.Read("761")
	if x.Id != "761" {
		log.Fatalln("id not 761")
	}
	if x.Title != "This is title" {
		log.Fatalln("title no ok")
	}
	if x.Keywords[0] != "可维护" {
		log.Fatalln("keywords[0] not ok")
	}
	if x.Keywords[5] != "前端" {
		log.Fatalln("keywords[5] not ok")
	}
}
