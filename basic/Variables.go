package main

import "fmt"

func printThis(str string) {
	fmt.Println(str)
}

func main() {
	f := "apple"
	printThis(f)
	printThis("banana")
}
