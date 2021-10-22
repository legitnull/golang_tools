package main

import (
	"fmt"
	"net"
	"sync"
)

func scan(port int, wsyn *sync.WaitGroup) {
	wsyn.Done()
	IP := "github.com"
	addr := fmt.Sprintf(IP+":%d", port)
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		return
	}
	fmt.Printf("%d open\n", port)
	conn.Close()
}

func main() {
	var wsyn sync.WaitGroup
	for i := 1; i < 65550; i++ {
		wsyn.Add(1)
		go scan(i, &wsyn)
	}
	wsyn.Wait()
}
