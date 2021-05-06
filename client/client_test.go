package client

import (
	"fmt"
	"net"
	"strings"
	"testing"
	"time"

	"github.com/zhao2490/my-rpc/config"
)

func _assert(condition bool, msg string, v ...interface{}) {
	if !condition {
		panic(fmt.Sprintf("assertion failed: "+msg, v...))
	}
}

func TestClient_dialTimeout(t *testing.T) {
	t.Parallel()
	l, _ := net.Listen("tcp", ":0")

	f := func(conn net.Conn, opt *config.Option) (client *Client, err error) {
		_ = conn.Close()
		time.Sleep(time.Second * 2)
		return nil, nil
	}
	t.Run("timeout", func(t *testing.T) {
		_, err := dialTimeout(f, "tcp", l.Addr().String(), &config.Option{ConnectTimeout: time.Second})
		t.Logf("err=%+v\n", err)
		_assert(err != nil && strings.Contains(err.Error(), "connect timeout"), "expect a timeout error")
	})
}
