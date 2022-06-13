package main

import (
	"fmt"
	"net"
	"strconv"
)

func main() {
	for i := 1; i < 1024; i++ {
		if i == 25 {
			continue
		}
		dest := "scanme.nmap.org:" + strconv.Itoa(i)
		conn, err := net.Dial("tcp", dest)

		if err != nil {
			continue
		}

		fmt.Printf("~> got conn for %v\n", dest)
		conn.Close()
	}
	fmt.Println("done!")
}
