package main

import (
	"fmt"
	"net"
	"os"
)

const (
	CONN_HOST = "localhost"
	//CONN_PORT = "3333"
	CONN_TYPE = "tcp"
)

func main() {
	// Listen for incoming connections.
	var CONN_PORT string
	fmt.Println("Enter the port")
	fmt.Scanf("%s", &CONN_PORT)
	var CONN_MSG string
	fmt.Println("Enter the message")
	fmt.Scanf("%s", &CONN_MSG)
	l, err := net.Listen(CONN_TYPE, CONN_HOST+":"+CONN_PORT)
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}
	// Close the listener when the application closes.
	defer l.Close()
	fmt.Println("Listening on " + CONN_HOST + ":" + CONN_PORT)
	for {
		// Listen for an incoming connection.
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			os.Exit(1)
		}

		// Handle connections in a new goroutine.
		go handleRequest(conn, CONN_MSG)
	}
}

// Handles incoming requests.
func handleRequest(conn net.Conn, msg string) {
	// Make a buffer to hold incoming data.
	buf := make([]byte, 8192)
	// Read the incoming connection into the buffer.
	_, err := conn.Read(buf)
	if err != nil {
		fmt.Println("Error reading:", err.Error())
	}
	// Send a response back to person contacting us.
	conn.Write([]byte("Message received."))
	conn.Write([]byte(msg)) //
	// Close the connection when you're done with it.
	conn.Close()
}

//execution steps
// // step-1
// durgaprasad@durgaprasad-Lenovo-E49:~/go/golang/src$ go run tcpp3.go
// Enter the port
// 3333
// Enter the message
// durgaprasad
// Listening on localhost:3333
// // step-2
// open another bash
// durgaprasad@durgaprasad-Lenovo-E49:~/go/golang$ echo -n "test"|nc localhost 3333
// Message received.durgaprasad
