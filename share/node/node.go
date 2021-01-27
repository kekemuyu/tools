package node

import (
	"net"
	"time"
)

type Client struct {
	IP     string
	OnLine bool
}

type Node struct {
	List      []Client
	LocalNode *Udp
	MsgBuff   map[int][]Message
}

func NewNode(udp *Udp) *Node {
	node := &Node{
		List:      make([]Client, 0),
		LocalNode: udp,
		MsgBuff:   make(map[int][]Message),
	}
	return node
}

//message send
func (t *Node) Send(msg []byte, addrs ...string) {
	for _, addr := range addrs {
		t.LocalNode.Send(msg, addr)
	}
}

func (t *Node) Run() {
	for {
		time.Sleep(time.Second)
		t.ClientUpdate()
		select {
		case re := <-t.LocalNode.Msg:
			if msg, err := bytesToMsg(re); err == nil {
				t.MsgBuff[msg.Topic] = append(t.MsgBuff[msg.Topic], msg)
			}
		case clientIP := <-t.LocalNode.ClientIP: //添加新的ip
			if !t.isIPExist(clientIP) {
				client := Client{
					IP:     clientIP,
					OnLine: true,
				}
				t.List = append(t.List, client)
			}
		}
	}
}

//get msg from udp
func (t *Node) Get(addr net.UDPAddr) ([]byte, error) {
	var bs []byte
	var err error
	return bs, err
}

func (t *Node) isIPExist(ip string) bool {
	for _, v := range t.List {
		if v.IP == ip {
			return true
		}
	}
	return false
}

func (t *Node) ClientUpdate() {
	for k, v := range t.List {
		if heartimeout, ok := t.LocalNode.ClientHeart[v.IP]; ok {
			if heartimeout == 0 {
				t.List[k].OnLine = false //掉线
			} else {
				t.List[k].OnLine = true //在线
			}
		}
	}
}

func (t *Node) GetClientList() []Client {
	return t.List
}

func (t *Node) GetClientMsg() map[int]Message {
	msgs := make(map[int]Message)
	return msgs
}
