package main

import (
	"fmt"
	"net"
	"strconv"
)

func main() {
	var openPorts []int
	portsChan := make(chan int, 10)
	resultsChan := make(chan int)

	for i := 0; i < cap(portsChan); i++ {
		go worker(portsChan, resultsChan)
	}

	go func() {
		for i := 1; i <= 24; i++ {
			portsChan <- i
		}
	}()

	for i := 0; i < 24; i++ {
		p := <-resultsChan
		if p != 0 {
			openPorts = append(openPorts, p)
		}
	}

	close(portsChan)
	close(resultsChan)

	for _, v := range openPorts {
		fmt.Printf("%d open\n", v)
	}
}

func worker(ports chan int, result chan int) {
	for p := range ports {
		dest := "scanme.nmap.org:" + strconv.Itoa(p)
		conn, err := net.Dial("tcp", dest)

		if err != nil {
			result <- 0
			continue
		}

		result <- p
		conn.Close()
	}
}
