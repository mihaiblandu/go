package main

import (
	"fmt"
	"syscall"
)

func main() {
	cwd, _ := syscall.Getwd()
	fmt.Println("Current cd: ", cwd)
	err := syscall.Mkdir("test", 777)
	if err == nil {
		println("OK")
	} else {
		fmt.Println("Error", err)
	}
}
