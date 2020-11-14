package main

import (
	"fmt"
)

func cald(name string, chd chan string) {

	chd <- name
}

func calp(name string, chp chan string) {
	chp <- name

}

func main() {

	chddd := make(chan string)
	chppp := make(chan string)
	go cald("durga", chddd)
	go calp("prasad", chppp)
	d, p := <-chddd, <-chppp
	fmt.Println("d=", d, "p=", p)
}
