package main

import "fmt"

func main() {

	var roll [3]int //type 1 decleration
	roll[0] = 11
	roll[1] = 22
	roll[2] = 33

	//type 2 declearation

	age := [3]int{17, 18, 19}

	//type 3 declearation

	height := []int{56, 57, 58}

	//type 4 declearation
	width := [3]int{10, 20}

	fmt.Println("rolls are", roll)
	fmt.Println("ages are", age)
	fmt.Println("heights are", height)
	fmt.Println("widths are", width)
}
