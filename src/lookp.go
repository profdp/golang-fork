/* LookupPort
 */

package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Fprintf(os.Stderr,
			"Usage: %s network-type service\n",
			os.Args[0])
		os.Exit(1)
	}
	networkType := os.Args[1]
	service := os.Args[2]

	port, err := net.LookupPort(networkType, service)
	if err != nil {
		fmt.Println("Error: ", err.Error())
		os.Exit(2)
	}

	fmt.Println("Service port ", port)
	os.Exit(0)
}

// durgaprasad@durgaprasad-Lenovo-E49:~/go/golang/src$ go run lookp.go tcp domain
// Service port  53

// durgaprasad@durgaprasad-Lenovo-E49:~/go/golang/src$ go run lookp.go tcp telnet
// Service port  23

// durgaprasad@durgaprasad-Lenovo-E49:~/go/golang/src$ go run lookp.go udp telnet
// Error:  lookup udp/telnet: Servname not supported for ai_socktype
// exit status 2

// durgaprasad@durgaprasad-Lenovo-E49:~/go/golang/src$ go run lookp.go udp domain
// Service port  53
