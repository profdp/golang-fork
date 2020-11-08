package main

import "fmt"

func main() {
	var a bool = true
	var b bool = false

	if !a == b {
		fmt.Println("false")
	}
	if 10 < 20 {
		fmt.Println(a)
	}
	if (10 < 20) == a {
		fmt.Println(a)
	} else if (10 < 20) == b {
		fmt.Println(b)
	}
	var x bool = 20 > 30
	if x {
		fmt.Println(b)
	}
	if x == a {
		fmt.Println(a)
	} else if x == b {
		fmt.Println(b)
	}
}
