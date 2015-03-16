package models

import (
	"db"
	"testing"
	"time"
)

var redisIp = "cp01-rdqa-dev301.cp01.baidu.com:8989"

func TestPostSave(t *testing.T) {
	err := db.Connect(redisIp)
	if err != nil {
		t.Fatal(err)
	}

	p := New()
	p.Id = "1"
	p.Title = "这是一篇测试文章"
	p.Content = `# Dillinger

Dillinger is a cloud-enabled, mobile-ready, offline-storage, AngularJS powered HTML5 Markdown editor.

  - Type some Markdown on the left
  - See HTML in the right
  - Magic

Markdown is a lightweight markup language based on the formatting conventions that people naturally use in email.  As [John Gruber] writes on the [Markdown site] [1]:

> The overriding design goal for Markdown's
> formatting syntax is to make it as readable
> as possible. The idea is that a
> Markdown-formatted document should be
> publishable as-is, as plain text, without
> looking like it's been marked up with tags
> or formatting instructions.

This text you see here is *actually* written in Markdown! To get a feel for Markdown's syntax, type some text into the left window and watch the results in the right.

### Version
3.0.2

### Tech

Dillinger uses a number of open source projects to work properly:

* [AngularJS] - HTML enhanced for web apps!
* [Ace Editor] - awesome web-based text editor
* [Marked] - a super fast port of Markdown to JavaScript
* [Twitter Bootstrap] - great UI boilerplate for modern web apps
* [node.js] - evented I/O for the backend
* [Express] - fast node.js network app framework [@tjholowaychuk]
* [Gulp] - the streaming build system
* [keymaster.js] - awesome keyboard handler lib by [@thomasfuchs]
* [jQuery] - duh

### Installation

You need Gulp installed globally:
` + "```sh\n" + "$ npm i -g gulp" + "```" + `

### Plugins

Dillinger is currently extended with the following plugins

* Dropbox
* Github
* Google Drive
* OneDrive

Readmes, how to use them in your own application can be found here:

* [plugins/dropbox/README.md](https://github.com/joemccann/dillinger/tree/master/plugins/dropbox/README.md)
* [plugins/github/README.md](https://github.com/joemccann/dillinger/tree/master/plugins/github/README.md)
* [plugins/googledrive/README.md](https://github.com/joemccann/dillinger/tree/master/plugins/googledrive/README.md)
* [plugins/onedrive/README.md](https://github.com/joemccann/dillinger/tree/master/plugins/onedrive/README.md)

### Development

Want to contribute? Great!

Dillinger uses Gulp + Webpack for fast developing.
Make a change in your file and instantanously see your updates!

Open your favorite Terminal and run these commands.


### Todo's

 - Write Tests
 - Rethink Github Save
 - Add Code Comments
 - Add Night Mode

License
----

MIT


**Free Software, Hell Yeah!**

[john gruber]:http://daringfireball.net/
[@thomasfuchs]:http://twitter.com/thomasfuchs
[1]:http://daringfireball.net/projects/markdown/
[marked]:https://github.com/chjj/marked
[Ace Editor]:http://ace.ajax.org
[node.js]:http://nodejs.org
[Twitter Bootstrap]:http://twitter.github.com/bootstrap/
[keymaster.js]:https://github.com/madrobby/keymaster
[jQuery]:http://jquery.com
[@tjholowaychuk]:http://twitter.com/tjholowaychuk
[express]:http://expressjs.com
[AngularJS]:http://angularjs.org
[Gulp]:http://gulpjs.com
`
	p.Description = "这里是一段描述"
	p.PubTime = time.Now()
	p.Save()

	p.Id = "2"
	p.Title = "文章2呵呵呵呵呵呵呵额呵呵呵呵呵呵呵"
	p.Save()
	p.Id = "3"
	p.Title = "文章3zxcvjklasdfupiqwerjklzxcvuipasdfjklwer呵呵呵呵呵呵呵额呵呵呵呵呵呵呵"
	p.Save()
	p.Id = "4"
	p.Title = "文章4呵呵呵呵呵呵呵额呵呵呵呵呵呵呵"
	p.Save()
	p.Id = "5"
	p.Title = "文章5呵呵呵呵呵呵呵额呵呵呵呵呵呵呵"
	p.Save()
	p.Id = "6"
	p.Title = "文章6呵呵呵呵呵呵呵额呵呵呵呵呵呵呵"
	p.Save()

	Rebuild()

}
