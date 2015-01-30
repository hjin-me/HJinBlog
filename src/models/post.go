package models

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/util"
)

type Article struct {
	Id          string `json:"id"`
	Alias       string `json:"alias"`
	Content     string
	Category    string
	Pubtime     int32
	Title       string
	Description string
	Tags        []string
	Keywords    []string
}

var dbIns *leveldb.DB

func getLevelIns() *leveldb.DB {
	if dbIns != nil {
		return dbIns
	}
	dbPath := os.Getenv("GOPATH") + "/src/db"
	db, err := leveldb.OpenFile(dbPath, nil)
	if err != nil {
		dbIns = nil
		panic(err)
	}
	dbIns = db
	fmt.Println("leveldb get")
	return dbIns
}

var ErrNotFound = leveldb.ErrNotFound

func Read(id string) (art Article, err error) {
	t := time.Now()
	start := t.UnixNano()
	db := getLevelIns()
	fmt.Println("level db got", id)

	data, err := db.Get([]byte(id), nil)
	if err == leveldb.ErrNotFound {
		return
	}
	if err != nil {
		// log.Fatal(err)
		return
	}

	t = time.Now()
	end := t.UnixNano()
	json.Unmarshal(data, &art)

	fmt.Println(art)
	fmt.Printf("%d - %d = %d", end, start, (end-start)/1000)
	return art, nil
}

func ReadBatch() (arts []Article, err error) {
	db := getLevelIns()

	iter := db.NewIterator(util.BytesPrefix([]byte("archive-")), nil)
	for iter.Next() {
		// Remember that the contents of the returned slice should not be modified, and
		// only valid until the next call to Next.
		art := Article{}
		json.Unmarshal(iter.Value(), &art)
		arts = append(arts, art)
	}
	iter.Release()
	if err == leveldb.ErrNotFound {
		return
	}
	if err != nil {
		// log.Fatal(err)
		return
	}

	return
}
