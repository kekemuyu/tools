package main

import (
	"tools/snet/config"
	"tools/snet/net"
	"tools/snet/serial"
	"tools/snet/utils/conn"

	log "github.com/donnie4w/go-logger/logger"
)

func main() {
	portname := config.Cfg.Section("serial").Key("PortName").String()
	baudrate, _ := config.Cfg.Section("serial").Key("BaudRate").Uint()
	netip := config.Cfg.Section("server").Key("ip").String()
	netport, _ := config.Cfg.Section("server").Key("port").Int64()

	log.Debug(portname, " ", baudrate, " ", netip, " ", netport)

	serialIO, err := serial.New(portname, baudrate)
	if err != nil {
		log.Error(err)
		return
	}

	netIO, err := net.New(netip, netport)
	if err != nil {
		log.Error(err)
		return
	}

	conn.Join(serialIO, netIO)

}
