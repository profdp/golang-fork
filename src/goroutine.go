package main

import (
	"fmt"
	"time"
)

func ms() {

	for i := 0; i < 5; i++ {
		fmt.Println("i is", i)
	}
}
func main() {
	ms()
	fmt.Println("before  goroutine")
	go ms()
	time.Sleep(time.Second)
	fmt.Println("after goroutine")
}
