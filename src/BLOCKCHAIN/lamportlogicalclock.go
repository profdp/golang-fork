package main

import "fmt"

//finds maximum of two
func max1(a int, b int) int {

	if a > b {
		return a
	} else {
		return b
	}
}

// Function to display the logical timestamp
func display(e1 int, e2 int, p1 []int, p2 []int) {
	fmt.Println()
	fmt.Print("The time stamps of events in P1: ")
	for i := 0; i < e1; i++ {
		fmt.Print(p1[i], "  ")
	}
	fmt.Println()
	fmt.Print("The time stamps of events in P2: ")

	// Print the array p2[]
	for i := 0; i < e2; i++ {
		fmt.Print(p2[i], "  ")
	}
	fmt.Println()
}

//lamplort-logic
//**********************************//
func lampclock(e1 int, e2 int, m [5][3]int) {
	var p1 = make([]int, e1)
	var p2 = make([]int, e2)

	// fmt.Println("m=", m)
	//Initialize p1[] and p2[]
	//***************************//
	for i := 0; i < e1; i++ {
		p1[i] = i + 1
	}
	for i := 0; i < e2; i++ {
		p2[i] = i + 1
	}
	//***************************//

	//display as a table
	//***************************//
	fmt.Print("\t")
	for i := 0; i < e2; i++ {

		fmt.Print("e2", "", i+1, "\t")
	}
	for i := 0; i < e1; i++ {
		fmt.Println("\t")
		fmt.Print("e1", "", i+1, "\t")
		for j := 0; j < e2; j++ {
			fmt.Print(m[i][j], "\t")
		}
	}
	//***************************//
	//change timestamps
	for i := 0; i < e1; i++ {

		for j := 0; j < e2; j++ {

			//Change the timestamp if the message is sent
			if m[i][j] == 1 {
				p2[j] = max1(p2[j], p1[i]+1)
				for i := j + 1; i < e2; i++ {
					p2[i] = p2[i-1] + 1
				}
			}
			//Change the timestamp if the message is reeived
			if m[i][j] == -1 {
				p1[i] = max1(p1[i], p2[j]+1)
				for k := i + 1; k < e1; k++ {
					p1[k] = p1[k-1] + 1
				}
			}
		}
	}
	display(e1, e2, p1[:], p2[:])
}

//**********************************//
//main
func main() {
	e1 := 5
	e2 := 3
	// initialize array of events
	var m [5][3]int
	m[0][0] = 0
	m[0][1] = 0
	m[0][2] = 0
	m[1][0] = 0
	m[1][1] = 0
	m[1][2] = 1
	m[2][0] = 0
	m[2][1] = 0
	m[2][2] = 0
	m[3][0] = 0
	m[3][1] = 0
	m[3][2] = 0
	m[4][0] = 0
	m[4][1] = -1
	m[4][2] = 0
	lampclock(e1, e2, m)
}
