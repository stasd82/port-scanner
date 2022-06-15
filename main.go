package main

import (
	"fmt"
	"net"
	"strconv"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	for i := 1; i < 1024; i++ {
		wg.Add(1)

		go func(portNumber int) {
			defer wg.Done()

			if portNumber == 25 {
				return
			}

			dest := "scanme.nmap.org:" + strconv.Itoa(portNumber)
			conn, err := net.Dial("tcp", dest)

			if err != nil {
				return
			}

			conn.Close()
			fmt.Printf("~> got conn for %v\n", dest)
		}(i)
	}
	wg.Wait()
	fmt.Println("done!")
}
