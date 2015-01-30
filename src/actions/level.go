package actions

import (
	"bytes"
	"encoding/json"
	"fmt"
	"framework"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/syndtr/goleveldb/leveldb"
)

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
	return dbIns
}

func ReadBatch(w http.ResponseWriter, r *http.Request, context fw.Context) {

	db := getLevelIns()

	iter := db.NewIterator(nil, nil)
	var res [][]byte
	for iter.Next() {
		// Remember that the contents of the returned slice should not be modified, and
		// only valid until the next call to Next.
		var b [][]byte
		fmt.Println(string(iter.Key()[:]))
		fmt.Println(string(iter.Value()[:]))
		b = append(b, iter.Key())
		b = append(b, iter.Value())
		s := bytes.Join(b, []byte(":"))
		fmt.Println(s)

		res = append(res, s)
	}
	iter.Release()

	j := []string{}
	for _, v := range res {
		var buffer bytes.Buffer
		_, err := buffer.Write(v)
		if err != nil {
			log.Fatal(err)
		}
		j = append(j, buffer.String())
	}

	json, err := json.Marshal(j)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(r.Header.Get("accept"))
	fmt.Fprintf(w, "%s", json)
}

func Read(w http.ResponseWriter, r *http.Request, context fw.Context) {
	t := time.Now()
	start := t.UnixNano()
	db := getLevelIns()
	id := context.Params["id"]
	data, err := db.Get([]byte(id), nil)
	if err == leveldb.ErrNotFound {
		http.NotFound(w, r)
		return
	}
	if err != nil {
		log.Fatal(err)
		return
	}

	t = time.Now()
	end := t.UnixNano()
	art := Article{}
	json.Unmarshal(data, &art)
	context.Json(art)
	fmt.Printf("%d - %d = %d", end, start, (end-start)/1000)

}

func Delete(w http.ResponseWriter, r *http.Request, context fw.Context) {
	id := context.Params["id"]

	db := getLevelIns()
	err := db.Delete([]byte(id), nil)
	if err != nil {
		log.Fatal(err)
	}
	h := w.Header()
	h.Add("Content-Type", "application/json")
	str := "{\"err\":0}"
	w.Write([]byte(str))
}

func Update(w http.ResponseWriter, r *http.Request, context fw.Context) {
	id := context.Params["id"]
	// b, _ := ioutil.ReadAll(r.Body)
	dec := json.NewDecoder(r.Body)
	//fmt.Printf("%s", b)
	var art Article
	// json.Unmarshal(b, &art)

	dec.Decode(&art)
	art.Id = id
	str, _ := json.Marshal(art)
	db := getLevelIns()
	err := db.Put([]byte(id), str, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(str))
	h1 := w.Header()
	h1.Add("Content-Type", "application/json")
	w.Write(str)
}

func Create(w http.ResponseWriter, r *http.Request, context fw.Context) {

	db := getLevelIns()
	err := db.Put([]byte("test-"+strconv.FormatInt(time.Now().Unix(), 10)), []byte("hello level"), nil)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	value, err := db.Get([]byte("test"), nil)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	w.Write(value)

}
func LevelInit(w http.ResponseWriter, r *http.Request, context fw.Context) {
	fmt.Fprintf(w, "hello level")
}
