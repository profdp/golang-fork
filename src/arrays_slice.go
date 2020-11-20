package main

import "fmt"

func main() {

	rolls := [5]int{10, 20, 30, 40, 50}
	for i := 0; i < len(rolls); i++ {

		fmt.Println("rolls are", rolls[i])
	}
	//slice-1  basci method from arrays
	fmt.Println("slice from  2 to last", rolls[2:])
	//slice-1 using make mathod
	roll1 := make([]int, 2, 4)
	fmt.Println("roll1", roll1)
	// slice-2 using make method-2

	roll2 := make([]int, 3)
	fmt.Println("roll2", roll2)

	// slice-3 using without make method-3

	roll3 := []int{10, 20, 30}
	fmt.Println("roll3", roll3)

	//append to slice
	roll3 = append(roll3, 40, 50)
	fmt.Println(" roll3", roll3)
	roll2 = append(roll2, 60, 70)
	fmt.Println(" roll2", roll2)

	//METHOD-1-->Append slice to another slice

	roll3 = append(roll3, roll2...)
	fmt.Println(" roll3", roll3)
	//aMETHOD-2-->ppend slice to another slice
	roll2 = append(roll2, roll1[0], roll1[1])
	fmt.Println("roll2", roll2)
	//METHOD-3-->append slice to another slice
	for i := 0; i < len(roll1); i++ {
		roll2 = append(roll2, roll1[i])
	}
	fmt.Println("roll2", roll2)
}
