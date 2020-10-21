package main

import "fmt"

func prinComb(a, b int) {
	if a < 9 {
		fmt.Print("0", a)
	} else {
		fmt.Print(a)
	}
	fmt.Print(" ")
	if b < 9 {
		fmt.Print("0", b)
	} else {
		fmt.Print(b)
	}
	if !(a == 98 && b == 99) {
		fmt.Print(", ")
	}
}

func ftPrintComb2() {
	a := 0
	b := 0

	for a < 99 {
		b = a + 1
		for b < 100 {

			prinComb(a, b)
			b++
		}
		a++
	}
}
