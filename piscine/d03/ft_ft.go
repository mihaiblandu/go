package main

import "fmt"

func pointer(number *int) {
	*number = 42
}

func main() {
	var a int
	pointer(&a)
	fmt.Println(a)
}
