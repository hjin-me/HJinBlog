package models

import (
	"db"
	"encoding/json"
	"log"
)

func ScanByPrefix(prefix string) <-chan string {
	var (
		err error
		ch  = make(chan string)
	)

	go func() {
		defer close(ch)

		for key := range db.Scan(prefix) {

			ch <- string(key)

		}

		if err = db.Err(); err != nil {
			return
		}

	}()

	return ch
}

func Scan() (ps []Post, err error) {
	for id := range db.Scan("post:") {
		x := Read(string(id))
		ps = append(ps, x)
	}

	if err = db.Err(); err != nil {
		return
	}
	sort(ps)

	return
}
func ZScan(key string) (ps []Archive, err error) {

	for rec := range db.ZScan(key) {
		x := Archive{}
		err := json.Unmarshal(rec[0], &x)
		if err != nil {
			log.Printf("unmarshal zsort failed %v %s", err, rec)
		} else {
			ps = append(ps, x)
		}
	}

	if err = db.Err(); err != nil {
		return
	}
	sortArchive(ps)

	return
}

func sortArchive(ps []Archive) {

	if len(ps) <= 1 {
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
		if ps[tgt].PubTime.Before(ps[last].PubTime) {
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
		sortArchive(ps[0:tgt])
	}
	if tgt != len(ps)-2 {
		// log.Print("div big", tgt, curr, last, len(ps))
		sortArchive(ps[tgt+1:])
	}
}

func sort(ps []Post) {
	if len(ps) <= 1 {
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
		if ps[tgt].PubTime.Before(ps[last].PubTime) {
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
