package main

import (
	"time"
	"tools/snet/utils/conn"

	log "github.com/donnie4w/go-logger/logger"
)

func main() {
	conn, err := conn.ConnectServer("", 9000)
	if err != nil {
		log.Error(err)
		return
	}

	go func() {
		buffer := make([]byte, 2014)
		for {
			n, _ := conn.TcpConn.Read(buffer)
			if n > 0 {
				log.Debug(n, string(buffer[:n]))
			}
		}
	}()

	for {
		time.Sleep(time.Second)
		conn.TcpConn.Write([]byte("hello,send from net client"))
	}

}
