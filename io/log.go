package main

import (
	"fmt"
	"log"
)

func main() {
	fmt.Println("Before fatal error")
	log.Fatalln("Fatal error")
	fmt.Println("After fatal error")
}
