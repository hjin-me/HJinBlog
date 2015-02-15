package models

import (
	"db"
	"log"
)

func Scan() (ps []Post, err error) {
	for id := range db.Scan() {
		x := Read(string(id))
		ps = append(ps, x)
	}

	if err = db.Err(); err != nil {
		return
	}

	return
}

func sort(ps []Post) {
	if len(ps) <= 1 {
		log.Println("cant divide")
		return
	}

	var (
		tgt  int = 0
		curr int = 0
		last int = 1
	)
	tgt = 0
	curr = 1
	last = 1
	for last < len(ps) {
		if ps[tgt].PubTime.After(ps[last].PubTime) {
			if last > curr {
				ps[tgt], ps[curr], ps[last] = ps[last], ps[tgt], ps[curr]
			} else {
				ps[tgt], ps[last] = ps[last], ps[tgt]
			}
			tgt++
			curr++
			last++
		} else {
			last++
		}
	}
	if tgt != 0 {
		sort(ps[0:tgt])
	}
	if tgt != len(ps)-2 {
		// log.Print("div big", tgt, curr, last, len(ps))
		sort(ps[tgt+1:])
	}
}
