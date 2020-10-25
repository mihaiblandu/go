package main

import (
	"fmt"
	"net"
)

func main() {
	addr, err := net.LookupIP("mihaiblandu.com")
	if err != nil {
		fmt.Println("unknown error :", err)
	} else {
		fmt.Println("Ip ", addr)
	}
}
