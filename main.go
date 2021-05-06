package main

import (
	"context"
	"log"
	"net"
	"time"

	"github.com/zhao2490/my-rpc/client"
	"github.com/zhao2490/my-rpc/service"
)

type Bar int

func (b Bar) Timeout(argv int, reply *int) error {
	//time.Sleep(time.Second * 2)
	return nil
}

func startServer(addr chan string) {
	var bar Bar
	if err := service.Register(&bar); err != nil {
		log.Fatal("register error:", err)
	}
	// pick a free port
	l, err := net.Listen("tcp", ":0")
	if err != nil {
		log.Fatal("network error:", err)
	}
	log.Println("start rpc server on", l.Addr())
	addr <- l.Addr().String()
	service.DefaultServer.Accept(l)
}

func main() {
	addrCh := make(chan string)
	go startServer(addrCh)
	addr := <-addrCh
	time.Sleep(time.Second)

	c, _ := client.Dial("tcp", addr)
	ctx, _ := context.WithTimeout(context.Background(), time.Second)
	var reply int
	err := c.Call(ctx, "Bar.Timeout", 1, &reply)
	log.Printf("err=%v", err)
}
