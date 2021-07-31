package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

// listen is for creating a server
// and dial is for using a server

// Listen returns a listener. The listener interface requires 3 functions to be
// implemented: Accept, Close, Addr.

// Accept waits for and returns the next connection to the listener.
// Hence, we will have an infinite loop for accepting connections.

// Functions implemented by Conn(Connection) interface ==> Read, Write, Close,
// SetDeadline, etc.
// Since Conn has both Read and Write methods, it implements both Reader and
// Writer interface.

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Panic(err)
	}
	defer li.Close()

	for {
		// If somebody makes a call to this server, we accept
		conn, err := li.Accept()
		if err != nil {
			log.Println(err)
		}
		io.WriteString(conn, "\n Hello from TCP server \n")
		fmt.Println(conn, "How is your day?")
		fmt.Fprintf(conn, "%v", "Well I hope")

		conn.Close()
	}
}

// On start of application
// Do you want the application “01_create-server” to accept incoming network connections?
