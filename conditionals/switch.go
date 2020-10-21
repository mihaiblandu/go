package main

import (
	"fmt"
	"time"
)

func main() {

	i := 2
	switch i {
	case 1:
		fmt.Println("1")
	case 2:
		fmt.Println("2")
	case 3:
		fmt.Println("3")
	}

	// Funy, but is verry useful
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's time to work")
	default:
		fmt.Println("It's a time to hork harder")
	}
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("AM Time")
	default:
		fmt.Println("PM Time")
	}
	//Lambda alike function

	lambda := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm an int")
		case string:
			fmt.Println("I'm an string")
		default:
			fmt.Printf("Don't know type of %T\n", t)
		}
	}

	lambda(true)
	lambda(1)
	lambda(0.1)
	lambda("hey")
}
