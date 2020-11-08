package main

import "fmt"

func main() {

	var len float32
	var bre float32
	var area float32

	fmt.Println("Enter the length")
	fmt.Scanf("%g", &len)
	fmt.Println("Enter the braedth")
	fmt.Scanf("%g", &bre)
	area = len * bre
	fmt.Println("area is ", area)
}
