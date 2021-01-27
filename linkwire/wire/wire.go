package wire

import (
	"io"
	"log"

	"github.com/jacobsa/go-serial/serial"
)

type Context struct {
	Rwc io.ReadWriteCloser
}

func New(opt serial.OpenOptions) *Context {
	var err error
	log.Println(opt)
	opt.MinimumReadSize = 4
	rwc, err := serial.Open(opt)

	if err != nil {
		log.Println(err)
		return nil

	}

	log.Println("serile is opened")
	return &Context{
		Rwc: rwc,
	}
}
