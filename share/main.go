package main

import (
	"fmt"
	"time"

	// "errors"
	// "fmt"

	"share/node"
)

func main() {
	udp := node.NewUdp()
	go udp.Run()
	node := node.NewNode(udp)
	go func() {
		for {
			time.Sleep(time.Second)
			fmt.Println(node.GetClientList())
		}
	}()
	node.Run()
}
