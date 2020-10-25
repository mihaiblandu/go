package main

import "fmt"

func fibonacci(nb int) int {
	fmt.Println(nb)
	if nb < 2 {
		return 1
	}
	return fibonacci(nb-1) + fibonacci(nb-2)
}

func main() {

	fmt.Println(fibonacci(9))

}
