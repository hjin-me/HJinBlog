package db

import (
	"errors"
	"log"
)

var lastErr error

func Err() error {
	return lastErr
}

func redisScan(ch chan interface{}, cursor []byte) {
	reply, err := Do("scan", cursor)
	if err != nil {
		lastErr = err
		close(ch)
		return
	}

	row, ok := reply.([]interface{})
	if !ok {
		lastErr = errors.New("reply error")
		log.Printf("reply is %v\n", reply)
		close(ch)
		return
	}
	cursor, ok = row[0].([]byte)
	if !ok {
		log.Printf("row[0] is %v\n", row[0])
		lastErr = errors.New("cursor error")
		close(ch)
		return
	}
	keys, ok := row[1].([]interface{})
	if !ok {
		log.Printf("row[1] is %v\n", row[1])
		lastErr = errors.New("keys error")
		close(ch)
		return
	}
	for _, k := range keys {
		s, ok := k.([]byte)
		if ok {
			ch <- s
		}
	}
	if string(cursor) == "0" {
		close(ch)
		lastErr = nil
		return
	}

	redisScan(ch, cursor)
}

func Scan() chan interface{} {

	ch := make(chan interface{})
	go func() {
		redisScan(ch, []byte("0"))
	}()
	return ch
}
