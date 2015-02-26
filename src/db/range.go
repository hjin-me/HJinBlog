package db

import (
	"errors"
	"log"
)

func redisZRange(ch chan []byte, k string, start, stop int) {

	var (
		reply interface{}
		err   error
	)
	defer close(ch)
	reply, err = Do("zrevrange", k, start, stop)

	if err != nil {
		lastErr = err
		return
	}
	lastErr = nil

	row, ok := reply.([]interface{})
	if !ok {
		lastErr = errors.New("reply error")
		log.Printf("reply is %v\n", reply)
		return
	}
	for _, r := range row {
		data, ok := r.([]byte)
		if !ok {
			log.Printf("r is %v\n", r)
			lastErr = errors.New("row error")
			return
		} else {
			ch <- data
		}
	}

}

func ZRange(key string, start, end int) chan []byte {

	ch := make(chan []byte)
	go func() {
		redisZRange(ch, key, start, end)
	}()
	return ch
}
