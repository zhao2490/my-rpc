package main

import (
	"fmt"
	"log"
	"net"
	"sync"
	"time"

	"github.com/zhao2490/my-rpc/client"
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
	//conn, _ := os.OpenFile("./zzz.txt", os.O_CREATE|os.O_WRONLY, 0777)
	client, _ := client.Dial("tcp", <-addr)
	defer func() { _ = client.Close() }()

	time.Sleep(time.Second)
	// send request & receive response
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			args := fmt.Sprintf("geerpc req %d", i)
			var reply string
			if err := client.Call("Foo.Sum", args, &reply); err != nil {
				log.Fatal("call Foo.Sum error:", err)
			}
			log.Println("reply:", reply)
		}(i)
	}
	wg.Wait()
}
