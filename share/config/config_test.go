package config

import (
	"testing"
)

func Test_init(t *testing.T) {
	t.Log("test config")
	t.Log(MustGetString("localaddr.ip"))
	t.Log(MustGetString("localaddr.port"))
	t.Log(MustGetString("localaddr.ping_port"))
	t.Log(MustGetString("remoteaddr.ip"))
	t.Log(MustGetString("remoteaddr.port"))
}
