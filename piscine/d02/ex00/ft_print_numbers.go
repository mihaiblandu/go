package main

import "fmt"

func ftPrintNumbers() {
	c := '0'
	for c <= '9' {
		fmt.Printf("%c", c)
		c++
	}
}
