package main

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
)

func generateData2(filename string) (<-chan []string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, errors.New(filename)
	}

	output := make(chan []string)

	go func() {
		cr := csv.NewReader(f)
		cr.FieldsPerRecord = 3

		for {
			record, err := cr.Read()
			if err == io.EOF {
				close(output)
				return
			}
			// fmt.Println(record)
			output <- record
		}
	}()
	return output, nil
}

func convertOne(in <-chan []string) chan []string {
	out := make(chan []string)
	go func(in <-chan []string) {
		defer close(out)
		for c := range in {
			c[0], c[1], c[2] = c[1], c[2], c[0]
			out <- c
		}
	}(in)
	return out
}

func displayData2(in chan []string) chan string {
	output := make(chan string)
	go func(in chan []string) {
		defer close(output)
		for c := range in {
			output <- fmt.Sprintf("First: %s, Second: %s, Third: %s", c[0], c[1], c[2])
		}
	}(in)
	return output
}

func main() {
	//c := displayData(convertTwo(convertOne(generateData("file.csv"))))
	records, err := generateData2("/Users/rashed/ZOOMINFO/Code_Practice/LeetCode-in-Go/educative/concurrency/patterns/pipeline_pattern/file.csv")
	if err != nil {
		log.Fatalf("Could not read csv %v", err)
	}
	for data := range displayData2(convertOne(records)) {
		fmt.Println(data)
	}
}
