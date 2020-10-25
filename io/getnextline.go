package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	b, err := ioutil.ReadFile("./io/cd.go")
	if err != nil {
		fmt.Println(err)
	}
	str := string(b)

	fmt.Println(str)

	file, err := os.Create("golang.txt")
	if err != nil {
		log.Fatal("We got error", err)
	}
	defer file.Close()
	fmt.Fprintf(file, "Hello go!!")
}
