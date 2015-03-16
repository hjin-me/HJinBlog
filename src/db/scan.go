package db

import (
	"errors"
	"log"
)

var lastErr error

func Err() error {
	return lastErr
}

func redisScan(ch chan []byte, cursor []byte, args ...string) {
	var (
		reply interface{}
		err   error
	)
	if len(args) == 1 {
		reply, err = Do("scan", cursor, "MATCH", args[0]+"*")
	} else {
		reply, err = Do("scan", cursor)
	}

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

func redisZScan(ch chan [2][]byte, cursor []byte, k string) {
	var (
		reply interface{}
		err   error
	)
	reply, err = Do("zscan", k, cursor)

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
	for i := 0; i < len(keys); i += 2 {
		p, _ := keys[i+1].([]byte)
		s, _ := keys[i].([]byte)
		ch <- [2][]byte{s, p}

	}
	if string(cursor) == "0" {
		close(ch)
		lastErr = nil
		return
	}

	redisZScan(ch, cursor, k)
}

func Scan(args ...string) chan []byte {

	ch := make(chan []byte)
	go func() {
		redisScan(ch, []byte("0"), args...)
	}()
	return ch
}

func ZScan(key string) chan [2][]byte {

	ch := make(chan [2][]byte)
	go func() {
		redisZScan(ch, []byte("0"), key)
	}()
	return ch
}
