package main

import "fmt"

func main() {
	i := 1
	for i < 10 {
		fmt.Println(i)
		i++
	}
	i = 0
Here:
	if i < 3 {
		fmt.Println("Goto : ", i)
		i++
		goto Here
	}
	//Infinite
	for {
		fmt.Println("infinite loop")
		break
	}
}
