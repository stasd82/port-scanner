package main

import (
	"fmt"
	"net"
	"strconv"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	portsChan := make(chan int, 10)

	for i := 0; i < cap(portsChan); i++ {
		go worker(portsChan, &wg)
	}

	for i := 1; i <= 24; i++ {
		wg.Add(1)
		portsChan <- i
	}

	close(portsChan)
	wg.Wait()

	fmt.Println("done!")
}

func worker(ports chan int, wg *sync.WaitGroup) {
	for p := range ports {
		dest := "scanme.nmap.org:" + strconv.Itoa(p)
		conn, err := net.Dial("tcp", dest)

		if conn != nil {
			fmt.Printf("~> got conn for %v\n", dest)
			conn.Close()
		}
		if err != nil {
			fmt.Printf("~> error for %v\n", dest)
		}

		wg.Done()
	}
}
