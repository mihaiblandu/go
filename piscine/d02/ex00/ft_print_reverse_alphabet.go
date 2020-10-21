package main

import "fmt"

func ftPrintReverseAlphabet() {
	c := 'z'
	for c >= 'a' {
		fmt.Printf("%c", c)
		c--
	}
}
