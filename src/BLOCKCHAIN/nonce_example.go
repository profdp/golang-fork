package main

import (
	"crypto/sha256"
	"fmt"
	"strconv"
	"time"
)

func main() {
	data := []byte("11")
	//hash of string
	hash := sha256.Sum256(data)
	fmt.Println("data=", data)
	fmt.Println("hash=", hash)
	fmt.Println(hash[0])
	//hash of integer
	n := 25
	nb := []byte(strconv.Itoa(n))
	fmt.Println("nb=", nb)
	hash1 := sha256.Sum256(nb)
	fmt.Println("hash1[0]=", hash1[0])
	//checking for target value
	target := 222
	if int(hash1[0]) > target {
		fmt.Println("yes grater than target")
	} else {
		fmt.Println("not grater than target")
	}
	//****************************
	hash2 := sha256.Sum256(append(data, nb...))

	fmt.Println("hash2[0]=", hash2[0])
	time.Sleep(3000 * time.Millisecond)
	fmt.Println("target= ", target)
	fmt.Println((int(hash2[0]) > target))
	//repeat for nonce
	for i := true; i; i = (int(hash2[0]) > target) {
		n = n + 1
		nbn := []byte(strconv.Itoa(n))
		fmt.Println("nbn=", nbn)
		hash2 = sha256.Sum256(append(data, nbn...))
		fmt.Println("n=", n)
		fmt.Println("hash2[0]=", hash2[0])
		time.Sleep(3000 * time.Millisecond)
	}
	fmt.Println("nonce=", n)
}
