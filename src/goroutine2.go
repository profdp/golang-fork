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
func mssg(msg string) {
	for i := 0; i < 5; i++ {
		//fmt.Println("sleep for 3 sec")
		//time.Sleep(time.Second * 3)
		fmt.Println("message is ", msg)
	}
}
func main() {
	ms()
	fmt.Println("goroutine")
	go ms()
	time.Sleep(time.Second)
	go mssg("hii")
	time.Sleep(time.Second * 3)
}
