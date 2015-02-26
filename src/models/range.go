package models

import (
	"db"
	"encoding/json"
	"log"
)

func ZRange(key string, start, stop int) (ps []Archive, err error) {

	for rec := range db.ZRange(key, start, stop) {
		x := Archive{}
		err := json.Unmarshal(rec, &x)
		if err != nil {
			log.Printf("unmarshal zrange failed %v %s", err, rec)
		} else {
			ps = append(ps, x)
		}
	}

	if err = db.Err(); err != nil {
		return
	}

	return
}
