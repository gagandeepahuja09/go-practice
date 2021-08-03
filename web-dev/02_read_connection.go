package main

// TCP - Read from connection
// We are going to use go routines so that our server can accept multiple connections.
// We will use bufio.Scanner for that
// The text that is going to come in will adhere to the HTTP(RFC 7230 ietf)

// Revise
// 1. Listen
// 2. Accept
// 3. Read & Write

import (
	"bufio"
	"fmt"
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
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
	}
	defer conn.Close()
}

// On start of application
// Do you want the application “01_create-server” to accept incoming network connections?
