package main

import (
	"fmt"
)

func hello(done chan bool) {
	fmt.Println("Hello world goroutine")
	done <- 1 < 2 // true  or false
}
func main() {
	done := make(chan bool)
	go hello(done)
	<-done //remove this line clear try executing
	fmt.Println("main function")

}
