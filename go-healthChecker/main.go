package main

import (
	"flag"
	"fmt"
	"os"
	"sync"
	"time"
)

func main() {

	domain := flag.String("d", "", "Domain to check")
	port := flag.Int("p", 0, "Port to hit ")

	flag.Parse()

	if *domain == "" || *port == 0 {
		fmt.Println("Please provide valid domain and port.")
		os.Exit(1)
	}

	if *port < 1 || *port > 65535 {
		fmt.Println("Port must be between 1 and 65535.")
		os.Exit(1)
	}

	if len(os.Args) < 3 {
		fmt.Println("Usage: go run main.go -d <domain> -p <port>")
		os.Exit(1)
	}

	var wg sync.WaitGroup
	wg.Add(1)
	status := make(chan bool)
	go CheckHealth(*domain, *port, &wg, status)

	select {
	case value := <-status:
		if value {
			fmt.Printf("The domain %s is healthy on port %d\n", *domain, *port)
		} else {
			fmt.Printf("The domain %s is not healthy on port %d\n", *domain, *port)
		}
	case <-time.After(5 * time.Second):
		fmt.Println("Request timed out")
	}

	wg.Wait()
}
