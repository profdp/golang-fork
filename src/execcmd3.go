package main

import (
	"fmt"
	"os/exec"
)

func main() {

	lsCmd := exec.Command("bash", "-c", "ls")
	lsOut, err := lsCmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println("> ls")
	fmt.Println(string(lsOut))
}
