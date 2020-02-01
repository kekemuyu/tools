package conn

import (
	"fmt"
	"testing"
	"time"
)

func serve() {
	l, err := Listen("0.0.0.0", 3011)

	fmt.Println(l, err)

	n := 0
	var tmp1conn, tmp2conn *Conn
	for {
		time.Sleep(time.Second)
		fmt.Println(l)
		if err != nil {
			continue
		}
		tmpconn, err := l.GetConn()
		if err != nil {
			continue
		}
		if n == 0 && err == nil {
			tmp1conn = tmpconn
			n = 1
			continue
		}
		if n == 1 && err == nil {
			tmp2conn = tmpconn
			n = 2
		}
		if n == 2 {

			fmt.Println("join start")
			Join(tmp1conn, tmp2conn)
			fmt.Println("join exit")
		}

	}

}

func Test_Join(t *testing.T) {
	go serve()
	isconnected := false
	var err error
	var con1, con2 *Conn

	conn1handle := func() {
		for {
			fmt.Println(con1.TcpConn.Write([]byte("hello")))
			time.Sleep(time.Second)
		}
	}

	for !isconnected {
		time.Sleep(time.Second)
		if con1, err = ConnectServer("127.0.0.1", 3011); err == nil {
			isconnected = true

			fmt.Println("con1:", con1, err, con1.GetLocalAddr())

			go conn1handle()
		}

	}

	isconnected = false
	for !isconnected {
		time.Sleep(time.Second)
		if con2, err = ConnectServer("127.0.0.1", 3011); err == nil {
			isconnected = true

			fmt.Println("con2:", con2, err, con2.GetLocalAddr())
		}
	}

	rdata := make([]byte, 1024)
	for {
		if con2 == nil {
			continue
		}
		fmt.Println("conn2 read")

		n, _ := con2.TcpConn.Read(rdata)
		fmt.Println("read conn2", string(rdata[:n]))

	}
}
