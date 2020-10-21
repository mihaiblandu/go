package main

import "fmt"

func ftPrintComb() {

	var a byte = 0
	var b byte = 0
	var c byte = 0

	for a <= 7 {
		b = a + 1
		for b <= 8 {
			c = b + 1
			for c <= 9 {
				//fmt.Print(a, b, c)
				print(a)
				print(b)
				print(c)
				if !(a == 7 && b == 8 && c == 9) {
					fmt.Print(", ")
				}
				c++
			}
			b++
		}
		a++
	}
}
