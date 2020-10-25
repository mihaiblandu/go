package main

import "fmt"

func euler() float64 {
	var h float64 = 1
	var e float64 = 1
	n := 1
	for n < 1000000000 {
		h = 1.0 / float64(n*n)
		e = e + h
		n++
	}
	return e

}

func main() {

	fmt.Println(euler())

}
