package main

import (
	"fmt"
	"github.com/jangirvipin/go-reminder/Time"
	"os"
	"strings"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	if len(os.Args) < 3 {
		fatalError := "Please provide a time and notification message in the format: '<Time> <Message>'"
		fmt.Println(fatalError)
	}

	time := os.Args[1]
	notificationMessage := strings.Join(os.Args[2:], " ")

	wg.Add(1)
	go TimeParser.TimeParse(time, notificationMessage, &wg)

	fmt.Println("Reminder set. Keep this program running until the notification time.")
	wg.Wait()
}
