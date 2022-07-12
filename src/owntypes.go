package main

import "fmt"

func main() {
	type dept []string
	depts := dept{"CSE", "EEE", "ECE"}

	for i, dep := range depts {
		fmt.Println(i, dep)
	}

	//creation of structure
	type clg struct {
		name     string
		clg_code uint
	}
	clg_engg := clg{"UCET", 4511}
	fmt.Println("DETAILS OF COLLEGE:", clg_engg)
	type clgs_list struct {
	}
	type univ struct {
		name           string
		univ_code      uint
		univ_clg_count uint
		//univ_clgs      []string
	}
	univ_NLG := univ{"MGU", 45, 2}
	fmt.Println("DETAILS OF UNIVERSITY:", univ_NLG)
}
