package http

import (
	"fmt"
	"log"
	"net"
	"testing"
)

func ExampleServe(t *testing.T) {
	server := Server{Handler: func(req Request, res *Response) {
		defer close(res.Body)
		res.Body <- []byte(fmt.Sprintf("You requested %v", req.URL))
	}}
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	server.Serve(ln)
}
