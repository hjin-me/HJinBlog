package models

import (
	"db"
	"encoding/json"
	"log"
	"time"
)

type Archive struct {
	Id      string
	Title   string
	PubTime time.Time
}

func (a Archive) String() string {
	s, err := json.Marshal(a)
	if err != nil {
		log.Println(err)
		return ""
	}
	return string(s)
}

func Clean() error {
	ps, err := Scan()
	if err != nil {
		return err
	}

	zsort := "pubtime"

	_, err = db.Do("del", zsort)
	if err != nil {
		log.Println(err)
		return err
	}

	for _, p := range ps {
		a := Archive{p.Id, p.Title, p.PubTime}
		_, err := db.Do("zadd", zsort, a.PubTime.Unix(), a)
		if err != nil {
			log.Println(err)
			return err
		}
	}

	return nil
}
