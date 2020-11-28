package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {
	fmt.Println("enter port number")
	var portnum string
	fmt.Scanf("%s", &portnum)

	fmt.Println("start server using telnet 127.0.0.1 portnumber")

	//listening on port entered
	lis, _ := net.Listen("tcp", ":"+portnum)
	// accept connection
	conn, _ := lis.Accept()

	// run loop some message counter c
	var c int
	c = 3
	for i := 0; i < c; i++ {
		// get message, output
		message, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Print("Message Received:", string(message))

	}
}

// step-1
// durgaprasad@durgaprasad-Lenovo-E49:~/go/golang/src$ go run basictcpserver2.go
// enter port number
// 3333
// start server using telnet 127.0.0.1 portnumber
// Message Received:hi
// Message Received:durga

// step-2
// durgaprasad@durgaprasad-Lenovo-E49:~/go/golang$ telnet 127.0.0.1 3333
// Trying 127.0.0.1...
// Connected to 127.0.0.1.
// Escape character is '^]'.
// hi
// durga
// ^Csignal: interrupt
// durgaprasad@durgaprasad-Lenovo-E49:~/go/golang/src$

// Connection closed by foreign host.
// durgaprasad@durgaprasad-Lenovo-E49:~/go/golang$
