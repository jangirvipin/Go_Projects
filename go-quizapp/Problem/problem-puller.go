package problem

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

type Problem struct {
	Question string
	Answer   string
}

var problems []Problem

func ProblemPuller() ([]Problem, error) {
	file, err := os.Open("sample_qa.csv")
	if err != nil {
		return nil, fmt.Errorf("cannot open file: %w", err)
	}

	// close the file when the function returns
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Printf("error closing file: %v\n", err)
		} else {
			fmt.Println("File closed successfully.")
		}
	}(file)

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("error reading CSV file: %w", err)
	}

	if len(records) == 0 {
		return nil, fmt.Errorf("CSV file is empty")
	}

	//Skipping Headers
	startIndex := 1

	for i, record := range records[startIndex:] {
		if len(record) < 2 {
			log.Printf("Skipping record %d: not enough fields", i+startIndex+1)
		}

		Problem := Problem{
			Question: record[0],
			Answer:   record[1],
		}

		problems = append(problems, Problem)
	}
	fmt.Println("\nFinished parsing CSV file.")
	return problems, nil
}
