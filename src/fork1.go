package main

import (
	"os"
	"os/exec"
	"syscall"
)

func main() {
	// check if ls command exists in the PATH
	binary, lookErr := exec.LookPath("ls")
	if lookErr != nil {
		panic(lookErr)
	}

	// the program we want to execute
	args := []string{"ls", "-a", "-l", "-h"}

	// fork new process and execute our program
	execErr := syscall.Exec(binary, args, os.Environ())

	// catch error if any
	if execErr != nil {
		panic(execErr)
	}
}
