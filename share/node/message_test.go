package node

import (
	"testing"
)

func Test_generateTopicFromIps(t *testing.T) {
	re := generateTopicFromIps("192.168.31.97", "192.168.31.94")
	if re != 63 {
		t.Log("gen topic error,the result is:", re)
	}
}
