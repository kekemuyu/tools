package net

import (
	"fmt"
	"net"

	log "github.com/donnie4w/go-logger/logger"
)

func New(bindAddr string, bindPort int64) (conn *net.TCPConn, err error) {
	tcpAddr, err := net.ResolveTCPAddr("tcp4", fmt.Sprintf("%s:%d", bindAddr, bindPort))
	listener, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		return nil, err
	}

	log.Debug("server listen:", fmt.Sprintf("%s:%d", bindAddr, bindPort))
	conn, err = listener.AcceptTCP()
	if err != nil {
		return nil, err
	}
	return conn, nil

}
