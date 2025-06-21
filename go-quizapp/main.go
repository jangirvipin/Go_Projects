package main

import (
	"bufio"
	"fmt"
	"github.com/jangirvipin/quiz/problem"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

var Counter int

func handleInput(scanner *bufio.Scanner, answer string) {
	for {
		userInput, ok := getUserInput(scanner)
		if !ok {
			fmt.Println("Time's up! Moving to the next question.")
			return
		}
		input := strings.TrimSpace(userInput)
		userAnswerInt, err := strconv.Atoi(input)

		if err != nil {
			fmt.Printf("Please enter a valid number. Error %v", err)
			fmt.Println("Try again.")
			continue
		}

		correctAnswerInt, err := strconv.Atoi(strings.TrimSpace(answer))

		if err != nil {
			fmt.Printf("Error parsing correct answer: %v\n", err)
			return
		}

		if userAnswerInt == correctAnswerInt {
			Counter++
		}
		return
	}
}

func getUserInput(scanner *bufio.Scanner) (string, bool) {
	inputchan := make(chan string)
	done := make(chan struct{})

	go func() {
		if scanner.Scan() {
			inputchan <- scanner.Text()
		}
	}()

	go func() {
		ticker := time.NewTicker(1 * time.Second)
		defer ticker.Stop()

		remaining := 30
		for {
			select {
			case <-ticker.C:
				remaining--
				fmt.Printf("\r ⏳ Time left: %2ds ", remaining)
				if remaining <= 0 {
					return
				}
			case <-done:
				return
			}
		}
	}()

	select {
	case input := <-inputchan:
		close(done)
		return input, true
	case <-time.After(30 * time.Second):
		close(done)
		return "", false
	}
}

func main() {

	problems, err := problem.ProblemPuller()
	if err != nil {
		log.Printf("Problem puller error: %v", err)
	}

	scanner := bufio.NewScanner(os.Stdin)

	for i, p := range problems {
		fmt.Println("Problem ", i+1, ":", p.Question)
		fmt.Print("Answer: ")
		handleInput(scanner, p.Answer)
	}
	fmt.Println("Quiz completed!")
	fmt.Printf("\nYou scored %d out of %d.\n", Counter, len(problems))

}
