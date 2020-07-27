// Package http implements an HTTP/1.1 server.
package http

import (
	"fmt"
	"io"
	"log"
	"math/rand"
	"net"
	"strings"
	"tools/comsurf/msg"
)

const (
	reqBuffLen    = 2 * 1024
	reqMaxBuffLen = 64 * 1024
)

var (
	socketCounter = 0
	verbose       = false
)

// Server defines the Handler used by Serve.
type Server struct {
	Handler func(Request, *Response)
}

// Serve starts the HTTP server listening on port. For each request, handle is
// called with the parsed request and response in their own goroutine.
func (s Server) Serve(ln net.Listener, myio io.ReadWriteCloser) error {
	r := rand.New(rand.NewSource(99))

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println("Error while accepting new connection", err)
			continue
		}

		socketCounter++
		if verbose {
			log.Println("handleConnection #", socketCounter)
		}
		req := Request{Headers: make(map[string]string)}
		res := Response{conn: conn, serialIO: myio, connID: r.Uint32()}
		go handleConnection(req, &res, s.Handler)
	}
	return nil
}

func readRequest(req Request, res *Response) (requestBuff []byte, err error) {
	requestBuff = make([]byte, 0, 8*1024)
	var reqLen int

	for {
		buff := make([]byte, reqBuffLen)
		reqLen, err = res.conn.Read(buff)
		var tempmsg msg.Msg
		tempmsg.Id = uint32(0x00000001)
		tempmsg.Data = append(tempmsg.Data, buff[:reqLen]...)
		tempmsg.Datalen = uint32(reqLen)
		packdata, _ := msg.Pack(tempmsg)
		res.serialIO.Write(packdata)
		break

		requestBuff = append(requestBuff, buff...)
		break
		if len(requestBuff) > reqMaxBuffLen {
			log.Println("Request is too big, ignoring the rest.")
			break
		}

		if err != nil && err != io.EOF {
			log.Println("Connection error:", err)
			break
		}

		if err == io.EOF || reqLen < reqBuffLen {
			break
		}
	}
	return
}

func handleConnection(req Request, res *Response, handle func(Request, *Response)) {
	defer func() {
		socketCounter--
		if verbose {
			log.Println(fmt.Sprintf("Closing socket:%d. Total connections:%d", res.connID, socketCounter))
		}
	}()
	reqLen := 0
	buffTotal := make([]byte, 0)
	for {
		readRequest(req, res)
		// requestBuff, err := readRequest(req, res)

		// if len(requestBuff) == 0 {
		// 	return
		// }

		// if err != nil && err != io.EOF {
		// 	log.Println("Error while reading socket:", err)
		// 	return
		// }

		// if verbose {
		// 	log.Println(string(requestBuff[0:]))
		// }

		for {
			buff := make([]byte, 8-reqLen)
			n, _ := res.serialIO.Read(buff)

			reqLen = +n
			if reqLen < 8 {
				continue
			}

			tempmsg, err := msg.Unpack(buff)
			if err != nil {
				log.Println(err)
				break
			}

			reqLen = 0
			break
		}

		for {
			buff = make([]byte, tempmsg.Datalen-reqLen)
			n, _ = res.serialIO.Read(buff) //读取串口数据
			reqLen = +n
			if reqLen < tempmsg.Datalen {
				continue
			}
			break
		}

		requestLines := strings.Split(string(buff[0:reqLen]), crlf)
		req.parseHeaders(requestLines[1:])
		err = req.parseInitialLine(requestLines[0])

		res.Body = make(chan []byte)
		go res.respondOther(req)

		if err != nil {
			res.Status = 400
			res.ContentType = "text/plain"

			res.Body <- []byte(err.Error() + "\n")
			close(res.Body)

			continue
		}

		requestIsValid := true
		log.Println(fmt.Sprintf("%s", requestLines[0]))

		if len(req.Headers["Host"]) == 0 {
			res.ContentType = "text/plain"
			res.Status = 400
			close(res.Body)
			requestIsValid = false
		}

		switch req.Method {
		case "GET", "HEAD":
		default:
			res.ContentType = "text/plain"
			res.Status = 501
			close(res.Body)
			requestIsValid = false
		}

		if requestIsValid {
			if req.Method == "HEAD" {
				close(res.Body)
			} else {
				handle(req, res)
			}
		}

		if req.Headers["Connection"] == "close" {
			break
		}
	}
}
