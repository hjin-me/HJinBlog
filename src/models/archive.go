package models

import (
	"db"
	"encoding/json"
	"log"
	"sync"
	"time"
)

type Archive struct {
	Id          string
	Title       string
	PubTime     time.Time
	Description string
}

func (a Archive) String() string {
	s, err := json.Marshal(a)
	if err != nil {
		log.Println(err)
		return ""
	}
	return string(s)
}

func (a *Archive) Unmarshal(b []byte) {
	x := make(map[string]interface{})
	err := json.Unmarshal(b, a)
	if err != nil {
		log.Println(err)
	}
	if pt, ok := x["PubTime"]; ok {
		if t, ok := pt.(int64); ok {
			a.PubTime = time.Unix(t, 0)
		}
	} else {
		log.Println("PubTime key not exists")
	}

}

func CleanArchive(ch <-chan string) {
	for k := range ch {
		_, err := db.Do("del", k)
		if err != nil {
			log.Println("del", err)
		}
	}
}

func AddArchive(ch <-chan string) {
	for k := range ch {
		p := Read(string(k))
		a := Archive{p.Id, p.Title, p.PubTime, p.Description}
		zsort := "z:" + p.Category
		// 插入分类集合
		_, err := db.Do("zadd", zsort, a.PubTime.Unix(), a)
		if err != nil {
			log.Println("zadd", err)
		}
		// 插入总集合
		_, err = db.Do("zadd", "z:archive", a.PubTime.Unix(), a)
		if err != nil {
			log.Println("zadd", err)
		}
	}
}

func RebuildArchive() error {

	// 删除存在的 ZSort
	zin := ScanByPrefix("z:")
	CleanArchive(zin)

	// 遍历所有文章
	pin := ScanByPrefix("post:")
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			AddArchive(pin)
			wg.Done()
		}()
	}
	wg.Wait()

	return nil
}
