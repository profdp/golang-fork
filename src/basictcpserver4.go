package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {
	mes("3333")

	mes("4444")

}
func mes(portnum string,) {

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
