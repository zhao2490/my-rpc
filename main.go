package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"time"

	"github.com/zhao2490/my-rpc/codec"
	"github.com/zhao2490/my-rpc/config"
	"github.com/zhao2490/my-rpc/server"
)

func startServer(addr chan string) {
	// pick a free port
	l, err := net.Listen("tcp", ":0")
	if err != nil {
		log.Fatal("network error:", err)
	}
	log.Println("start rpc server on", l.Addr())
	addr <- l.Addr().String()
	server.DefaultServer.Accept(l)
}

func main() {
	addr := make(chan string)
	go startServer(addr)

	// in fact, following code is like a simple my-rpc client
	conn, _ := net.Dial("tcp", <-addr)
	//conn, _ := os.OpenFile("./zzz.txt", os.O_CREATE|os.O_WRONLY, 0777)
	defer func() { _ = conn.Close() }()

	time.Sleep(time.Second)
	// send options
	_ = json.NewEncoder(conn).Encode(config.DefaultOption)
	cc := codec.NewGobCodec(conn)
	// send request & receive response
	for i := 0; i < 5; i++ {
		h := &codec.Header{
			ServiceMethod: "Foo.Sum",
			Seq:           uint64(i),
		}
		_ = cc.Write(h, fmt.Sprintf("my-rpc req %d", h.Seq))
		_ = cc.ReadHeader(h)
		log.Println("header:", h)
		var reply string
		_ = cc.ReadBody(&reply)
		log.Println("reply:", reply)
	}
}
