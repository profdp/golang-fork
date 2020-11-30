package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)

	for sc.Scan() {
		portnum := sc.Text()

		fmt.Println("start connection using telnet 127.0.0.1 ", portnum)
	}
}
