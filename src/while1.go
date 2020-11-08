package main

import "fmt"

func main() {

	var a int = 10
	var b int = 20
	for a < b {
		fmt.Printf("a=%d\n", a)
		a += 1
	}
}
