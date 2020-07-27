package main

import (
	"flag"
	"fmt"
	"io"
	"log"

	"net"
	"strconv"
	"tools/comsurf/config"
	http "tools/comsurf/http"
	"tools/comsurf/serial"
)

const usage = `Usage: gofile port [dir]`

var optRoot = ""

var serialIO io.ReadWriteCloser

func init() {
	portname := config.Cfg.Section("serial_client").Key("PortName").String()
	baudrate, _ := config.Cfg.Section("serial_client").Key("BaudRate").Uint()

	var err error
	serialIO, err = serial.New(portname, baudrate)
	if err != nil {
		log.Fatal(err)
		return
	}
}

func main() {
	flag.Usage = func() {
		log.Fatal(usage)
	}

	flag.Parse()

	if flag.NArg() < 1 || flag.NArg() > 2 {
		flag.Usage()
	}

	port, err := strconv.Atoi(flag.Args()[0])
	if err != nil {
		log.Fatal("Bad port number:", flag.Args()[0])
	}

	if flag.NArg() == 2 {
		optRoot = flag.Args()[1]
	}

	ln, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatal(err)
	}

	server := http.Server{Handler: fileServerHandleRequestGen(optRoot)}
	log.Println("Starting server on port", port)
	server.Serve(ln, serialIO)
}
