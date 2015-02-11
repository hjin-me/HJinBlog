package db

import "github.com/garyburd/redigo/redis"

type Cmd struct {
	Cmd  string
	Args []interface{}
	Resp chan CmdResp
}
type CmdResp struct {
	reply interface{}
	err   error
}

var (
	conn redis.Conn
	ch   chan Cmd
)

func Connect(redisServer string) error {
	var err error
	conn, err = redis.Dial("tcp", redisServer)
	if err != nil {
		return err
	}
	ch = make(chan Cmd)

	go func() {
		for cmd := range ch {
			resp := CmdResp{}
			resp.reply, resp.err = conn.Do(cmd.Cmd, cmd.Args...)
			cmd.Resp <- resp
			close(cmd.Resp)
		}
		conn.Close()
	}()

	return nil
}

func Close(conn redis.Conn) {
	close(ch)
}

func Do(cmd string, args ...interface{}) (interface{}, error) {
	c := Cmd{cmd, args, make(chan CmdResp)}
	go func() {
		ch <- c
	}()
	resp := <-c.Resp
	return resp.reply, resp.err
}
