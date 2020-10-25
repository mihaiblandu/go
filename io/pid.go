package main

import (
	"fmt"
	"syscall"
)

func main() {
	pid := syscall.Getppid()

	fmt.Println("Process Id :", pid)
}
