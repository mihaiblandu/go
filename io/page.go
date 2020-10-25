package main

import (
	"fmt"
	"os"
)

func main() {
	pagesize := os.Getpagesize()

	fmt.Println("Page size :", pagesize)
}
