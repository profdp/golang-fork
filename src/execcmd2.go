package main

import (
	"fmt"
	"io/ioutil"
	"os/exec"
)

func main() {

	grepCmd := exec.Command("grep", "durga")

	grepIn, _ := grepCmd.StdinPipe()
	grepOut, _ := grepCmd.StdoutPipe()
	grepCmd.Start()
	grepIn.Write([]byte("hello grep\ngoodbye durga grep"))
	grepIn.Close()
	grepBytes, _ := ioutil.ReadAll(grepOut)
	grepCmd.Wait()

	fmt.Println("> grep durga")
	fmt.Println(string(grepBytes))
}
