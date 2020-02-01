package serial

import (
	"io"

	"github.com/jacobsa/go-serial/serial"
)

func New(portnum string, baudrate uint) (io.ReadWriteCloser, error) {
	opt := serial.OpenOptions{
		PortName:        portnum,
		BaudRate:        baudrate,
		DataBits:        8,
		StopBits:        1,
		ParityMode:      serial.PARITY_NONE,
		MinimumReadSize: 8,
	}

	var err error
	var tempIO io.ReadWriteCloser

	tempIO, err = serial.Open(opt)

	return tempIO, err
}
