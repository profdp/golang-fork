package main

import (
	"fmt"
	"time"
)

func numbers() {
	for i := 1; i <= 5; i++ {
		time.Sleep(250 * time.Millisecond)
		fmt.Printf("%d ", i)
	}
}
func alphabets() {
	for i := 'a'; i <= 'e'; i++ {
		//1 time.Sleep(250 * time.Millisecond)
		fmt.Printf("%c ", i)
	}
}
func main() {
	go numbers()
	//1 go alphabets()
	alphabets()
	time.Sleep(3000 * time.Millisecond)
	fmt.Println("main terminated")
}
