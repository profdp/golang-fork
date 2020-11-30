package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)

	for sc.Scan() {
		a := sc.Text()

		fmt.Println(a)
	}
}
