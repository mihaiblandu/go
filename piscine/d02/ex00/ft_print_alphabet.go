package main

import "fmt"

func ftPrintAlphabet() {
	var c byte
	c = 'a'
	for c <= 'z' {
		fmt.Printf("%c", c)
		c++
	}
}
