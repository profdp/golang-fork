package main

import (
	"fmt"
	"time"
)

func hello(done chan bool, cname string) {
	fmt.Println("hello go routine is going to sleep")
	time.Sleep(4 * time.Second)
	fmt.Println("hello go routine awake and going to write to channel ", cname)
	done <- true
}
func main() {
	doit := make(chan bool)
	chname := "doit"
	fmt.Println("Main going to call hello go goroutine")
	go hello(doit, chname)
	<-doit
	fmt.Println("Main received data from channel", chname)
}
