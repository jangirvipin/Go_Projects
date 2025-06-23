package main

import (
	"fmt"
	"net/http"
	"strconv"
	"sync"
	"time"
)

func CheckHealth(domain string, port int, wg *sync.WaitGroup, status chan bool) {
	str := strconv.Itoa(port)
	var url = "http://" + domain + ":" + str
	fmt.Println(url)

	defer wg.Done()

	client := http.Client{
		Timeout: 5 * time.Second,
	}
	_, err := client.Get(url)
	if err != nil {
		status <- false
		return
	}
	status <- true
}
