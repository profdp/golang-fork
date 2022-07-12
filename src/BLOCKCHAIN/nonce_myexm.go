package main

import (
	"crypto/sha256"
	"fmt"
	"reflect"
)

func main() {
	data := []byte("11")
	hdata := sha256.Sum256(data)
	fmt.Println(hdata)
	fmt.Println(hdata[0])
	fmt.Println(reflect.TypeOf((hdata[0])))
	/*nb := []byte(strconv.Itoa(hdata[0]))
	fmt.Println("nb=", nb)*/
}
