package actions

import (
	"models/post"

	"github.com/hjin-me/banana"
)

type TimeArchive struct {
	Year int
	List []post.Post
}

func Scan(ctx banana.Context) {
	/*
		posts, err := models.ZScan("pubtime")
		if err != nil {
			log.Fatal(err)
		}
		y2i := make(map[int]int)
		ta := []TimeArchive{}

		for _, p := range posts {
			y := p.PubTime.Year()
			_, ok := y2i[y]
			if !ok {
				a := TimeArchive{}
				a.Year = y
				ta = append(ta, a)
				y2i[y] = len(ta) - 1
			}
			ta[y2i[y]].List = append(ta[y2i[y]].List, p)
		}
		ctx.Tpl("archives.html", ta)
	*/
}
