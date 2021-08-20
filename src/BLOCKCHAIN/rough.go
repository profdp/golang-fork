package main

import (
	"crypto/sha256"
	"fmt"
	"time"
)

func main() {

	currentTime := time.Now()
	fmt.Println(currentTime)
	hash := sha256.Sum256([]byte(currentTime.String()))
	fmt.Println("hash of TIME")
	fmt.Println(hash)

	input := "durga" + currentTime.String()
	fmt.Println(input)

	inputhash := sha256.Sum256([]byte(input))
	fmt.Println("hash of name+TIME")
	fmt.Println(inputhash)

}
