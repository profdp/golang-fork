package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os/exec"
	"time"
)

func main() {
	ph := 5
	for ph = 0; ph < 5; ph++ {
		mes()
	}
}
func mes() {
	cmd := []string{"go", "run", "conest.go"}

	proc := exec.Command(cmd[0], cmd[1], cmd[2])

	stdin, _ := proc.StdinPipe()
	defer stdin.Close()

	stdout, _ := proc.StdoutPipe()
	defer stdout.Close()

	go func() {
		s := bufio.NewScanner(stdout)
		for s.Scan() {
			fmt.Println("Parent says to:" + cmd[2] + "->" + s.Text())
		}
	}()
	// Start the process
	proc.Start()

	var portnum string

	//portnum := "3333"
	fmt.Println("enter the port number")
	fmt.Scanf("%s", &portnum)
	io.WriteString(stdin, portnum+"\n")
	// 	*******************************************
	time.Sleep(time.Second * 2)

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
		fmt.Print("Message Received from:"+cmd[2]+"is=", string(message))

	}
	conn.Close()

	proc.Process.Kill()

}

// <<<<<<<<<<<<<<<<-----------OUTPUT------->>>>>>>>>>>>>>>>>
// durgaprasad@durgaprasad-Lenovo-E49:~/go/golang/src/IPC$ go run basictcpserver6.go
// enter the port number
// 3333
// Parent says to:conest.go->start connection using telnet 127.0.0.1  3333
// Message Received from:conest.gois=hi3-1
// Message Received from:conest.gois=hi3-2
// Message Received from:conest.gois=hi3-2
// enter the port number
// 4444
// Parent says to:conest.go->start connection using telnet 127.0.0.1  4444
// Message Received from:conest.gois=hi4-1
// Message Received from:conest.gois=hi4-2
// Message Received from:conest.gois=hi4-3
// enter the port number
// 5555
// Parent says to:conest.go->start connection using telnet 127.0.0.1  5555
// Message Received from:conest.gois=hi5-1
// Message Received from:conest.gois=hi5-2
// Message Received from:conest.gois=hi5-3
// enter the port number
// 6666
// Parent says to:conest.go->start connection using telnet 127.0.0.1  6666
// Message Received from:conest.gois=hi6-1
// Message Received from:conest.gois=hi6-2
// Message Received from:conest.gois=hi6-3
// enter the port number
// 7777
// Parent says to:conest.go->start connection using telnet 127.0.0.1  7777
// Message Received from:conest.gois=hi7-1
// Message Received from:conest.gois=hi7-2
// Message Received from:conest.gois=hi7-3
// durgaprasad@durgaprasad-Lenovo-E49:~/go/golang/src/IPC$

// <<<<<<<<<<<<<<<<<<<-----OUTPUT BASH----------------------->>>>>>>>>>>>>>

// durgaprasad@durgaprasad-Lenovo-E49:~/go/golang$ telnet 127.0.0.1 3333
// Trying 127.0.0.1...
// Connected to 127.0.0.1.
// Escape character is '^]'.
// hi3-1
// hi3-2
// hi3-2
// Connection closed by foreign host.
// durgaprasad@durgaprasad-Lenovo-E49:~/go/golang$ telnet 127.0.0.1 4444
// Trying 127.0.0.1...
// Connected to 127.0.0.1.
// Escape character is '^]'.
// hi4-1
// hi4-2
// hi4-3
// Connection closed by foreign host.
// durgaprasad@durgaprasad-Lenovo-E49:~/go/golang$ telnet 127.0.0.1 5555
// Trying 127.0.0.1...
// Connected to 127.0.0.1.
// Escape character is '^]'.
// hi5-1
// hi5-2
// hi5-3
// Connection closed by foreign host.
// durgaprasad@durgaprasad-Lenovo-E49:~/go/golang$ telnet 127.0.0.1 6666
// Trying 127.0.0.1...
// Connected to 127.0.0.1.
// Escape character is '^]'.
// hi6-1
// hi6-2
// hi6-3
// Connection closed by foreign host.
// durgaprasad@durgaprasad-Lenovo-E49:~/go/golang$ telnet 127.0.0.1 7777
// Trying 127.0.0.1...
// Connected to 127.0.0.1.
// Escape character is '^]'.
// hi7-1
// hi7-2
// hi7-3
// Connection closed by foreign host.
// durgaprasad@durgaprasad-Lenovo-E49:~/go/golang$
