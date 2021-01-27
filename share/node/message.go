package node

import (
	"encoding/json"
	"strconv"
	"strings"
)

type Message struct {
	Topic   int //gen by the lase byte of srcIP and desIP with method bit xor
	Addr    string
	Content []byte
}

func bytesToMsg(bs []byte) (Message, error) {
	var msg Message
	var err error
	err = json.Unmarshal(bs, &msg)
	return msg, err
}

//from soure ip and target ip generate a byte with bit xor method
func generateTopicFromIps(ip1, ip2 string) int {
	index1 := strings.LastIndex(ip1, ".")
	index2 := strings.LastIndex(ip2, ".")

	lastbyte1, _ := strconv.Atoi(ip1[index1+1:])
	lastbyte2, _ := strconv.Atoi(ip2[index2+1:])

	return lastbyte1 ^ lastbyte2
}
