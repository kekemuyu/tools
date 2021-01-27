package node

import (
	"fmt"
	"net"
	"os"
	"strings"

	"share/config"
	"time"
)

type Udp struct {
	Addr             *net.UDPAddr //local address
	Conn             *net.UDPConn
	Msg              chan []byte
	ClientIP         chan string
	HeartBeatTimeOut int64
	ClientHeart      map[string]int64
}

func NewUdp() *Udp {
	return &Udp{
		Addr:             &net.UDPAddr{IP: net.ParseIP(config.MustGetString("localaddr.ip")), Port: config.MustGetInt("localaddr.port")},
		Msg:              make(chan []byte),
		ClientIP:         make(chan string, 1024),
		HeartBeatTimeOut: 10,
		ClientHeart:      make(map[string]int64),
	}
}

func (t *Udp) Run() {
	addr := &net.UDPAddr{IP: net.ParseIP(config.MustGetString("localaddr.ip")), Port: config.MustGetInt("localaddr.port")}
	go t.pingServe()
	go func() {
		listener, err := net.ListenUDP("udp", addr)
		if err != nil {
			fmt.Println(err)
			return
		}

		t.Conn = listener
		go read(listener)

	}()

	b := make([]byte, 1)
	os.Stdin.Read(b)
}

func (t *Udp) Send(msg []byte, ip string) {
	addr := &net.UDPAddr{IP: net.ParseIP(ip), Port: config.MustGetInt("remoteaddr.port")}
	t.Conn.WriteToUDP(msg, addr)
}

func (t *Udp) pingServe() {
	addr := &net.UDPAddr{IP: net.ParseIP(config.MustGetString("localaddr.ip")), Port: config.MustGetInt("localaddr.ping_port")}
	listener, err := net.ListenUDP("udp", addr)
	if err != nil {
		fmt.Println(err)
		return
	}
	go ping(listener)

	go t.pingread(listener)

	b := make([]byte, 1)
	os.Stdin.Read(b)

}

func ping(conn *net.UDPConn) {
	// addr1 := &net.UDPAddr{IP: net.ParseIP(config.MustGetString("localaddr.ip")), Port: config.MustGetInt("localaddr.ping_port")}
	addr2 := &net.UDPAddr{IP: net.ParseIP(config.MustGetString("remoteaddr.ip")), Port: config.MustGetInt("remoteaddr.ping_port")}
	localIP := getLocalIp(config.MustGetString("remoteaddr.ip"))
	fmt.Println(localIP)
	// conn2, err := net.DialUDP("udp", addr1, addr2)

	// if err != nil {
	// 	fmt.Println(err)
	// }
	// defer conn2.Close()

	for {
		time.Sleep(5 * time.Second)
		conn.WriteToUDP([]byte(localIP), addr2)
		// conn2.Write([]byte(localIP))

	}

}

func read(conn *net.UDPConn) {
	for {

		data := make([]byte, 1024)
		n, remoteAddr, err := conn.ReadFromUDP(data)
		if err != nil {
			fmt.Printf("error during read: %s", err)

		}
		if n > 0 {

			fmt.Println("msg:", string(data[:n]), remoteAddr)

		}
		//fmt.Printf("receive %s from <%s>\n", data[:n], remoteAddr)
	}
}

func (t *Udp) pingread(conn *net.UDPConn) {
	data := make([]byte, 1024)
	for {

		for k, v := range t.ClientHeart {
			if v > 0 {
				t.ClientHeart[k] = t.ClientHeart[k] - 1
			}
		}
		n, remoteAddr, err := conn.ReadFromUDP(data)
		if err != nil {
			fmt.Printf("error during read: %s", err)

		}

		if n > 0 {

			fmt.Println("ping:", string(data[:n]), remoteAddr)
			t.ClientIP <- remoteAddr.IP.String()
			t.ClientHeart[remoteAddr.IP.String()] = t.HeartBeatTimeOut

		}

		//fmt.Printf("receive %s from <%s>\n", data[:n], remoteAddr)
	}
}

//ip:192.168.10.x
func getLocalIp(ip string) string {
	addrs, _ := net.InterfaceAddrs()
	for _, addr := range addrs {

		strIpv4 := addr.String()
		index := strings.LastIndex(strIpv4, ".")
		length := strings.Index(strIpv4, `/`)

		if index == -1 {
			continue
		}
		index2 := strings.LastIndex(ip, ".")

		if strIpv4[:index] == ip[:index2] {
			return strIpv4[:length]
		}

	}
	return ""
}
