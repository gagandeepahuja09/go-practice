package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

func do(conn net.Conn) {
	buf := make([]byte, 1024)
	// blocking call
	_, err := conn.Read(buf)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("processing logic")
	time.Sleep(5 * time.Second)

	// blocking call
	conn.Write([]byte("HTTP/1.1 200 0K\r\n\r\nHello World!\r\n"))

	conn.Close()
}

func main() {
	// reserve port
	listener, err := net.Listen("tcp", ":1729")
	if err != nil {
		log.Fatal(err)
	}

	for {
		log.Println("waiting for client to accept connection")
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}

		log.Println("client connected")
		go do(conn)
		fmt.Println(conn)
	}
}

// IMPROVEMENTS
// 1. We should avoid CPU over-utilization by limiting the number of threads that are created.
// This can be done by setting a limit on the number of threads.
// We have never really done that in any of our system.
// 2. We should add a thread pool to save on thread creation time.
// 3. Add timeouts.
