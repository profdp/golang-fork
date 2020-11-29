package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {
	ph := 5
	for ph = 0; ph < 5; ph++ {
		mes()
	}
}
func mes() {
	var portnum string
	fmt.Println("enter the port number")
	fmt.Scanf("%s", &portnum)

	fmt.Println("start server using telnet 127.0.0.1 ", portnum)

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
	conn.Close()
}

// output
// durgaprasad@durgaprasad-Lenovo-E49:~/go/golang/src$ go run basictcpserver5.go
// enter the port number
// 3333
// start server using telnet 127.0.0.1  3333
// Message Received:hi3
// Message Received:hi3
// Message Received:hi3
// enter the port number
// 4444
// start server using telnet 127.0.0.1  4444
// Message Received:hi4
// Message Received:hi4
// Message Received:hi4
// enter the port number
// 5555
// start server using telnet 127.0.0.1  5555
// Message Received:hi5
// Message Received:hi5
// Message Received:hi5
// enter the port number
// 6666
// start server using telnet 127.0.0.1  6666
// Message Received:hi6
// Message Received:hi6
// Message Received:hi6
// enter the port number
// 7777
// start server using telnet 127.0.0.1  7777
// Message Received:hi7
// Message Received:hi7
// Message Received:hi7
// durgaprasad@durgaprasad-Lenovo-E49:~/go/golang/src$

// *********************

// durgaprasad@durgaprasad-Lenovo-E49:~/go/golang$ telnet 127.0.0.1 3333
// Trying 127.0.0.1...
// Connected to 127.0.0.1.
// Escape character is '^]'.
// hi3
// hi3
// hi3
// Connection closed by foreign host.
// durgaprasad@durgaprasad-Lenovo-E49:~/go/golang$ telnet 127.0.0.1 4444
// Trying 127.0.0.1...
// Connected to 127.0.0.1.
// Escape character is '^]'.
// hi4
// hi4
// hi4
// Connection closed by foreign host.
// durgaprasad@durgaprasad-Lenovo-E49:~/go/golang$ telnet 127.0.0.1 5555
// Trying 127.0.0.1...
// Connected to 127.0.0.1.
// Escape character is '^]'.
// hi5
// hi5
// hi5
// Connection closed by foreign host.
// durgaprasad@durgaprasad-Lenovo-E49:~/go/golang$ telnet 127.0.0.1 6666
// Trying 127.0.0.1...
// Connected to 127.0.0.1.
// Escape character is '^]'.
// hi6
// hi6
// hi6
// Connection closed by foreign host.
// durgaprasad@durgaprasad-Lenovo-E49:~/go/golang$ telnet 127.0.0.1 7777
// Trying 127.0.0.1...
// Connected to 127.0.0.1.
// Escape character is '^]'.
// hi7
// hi7
// hi7
// Connection closed by foreign host.
// durgaprasad@durgaprasad-Lenovo-E49:~/go/golang$
