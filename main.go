package main

import (
	"fmt"
	"net"
	"strconv"
	"time"
)

func main() {
	for i := 1; i < 1024; i++ {
		go func(portNumber int) {
			fmt.Println(portNumber)
			if portNumber == 25 {
				return
			}
			dest := "scanme.nmap.org:" + strconv.Itoa(portNumber)
			conn, err := net.Dial("tcp", dest)

			if err != nil {
				return
			}

			fmt.Printf("~> got conn for %v\n", dest)
			conn.Close()
		}(i)
	}
	time.Sleep(7 * time.Second)
	fmt.Println("done!")
}
