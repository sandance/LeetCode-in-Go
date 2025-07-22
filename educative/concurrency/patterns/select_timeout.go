package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"time"
)

// struct
type problem struct {
	q, a string
}

// The qa.csv is a sample file, you can create your own quiz

// error handling function
func check(e error) {
	if e != nil {
		panic(e)
	}
}

// function read csv file
func readCSVFile(filename string) [][]string {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Errorf("%v", err)
	}

	csvReader := csv.NewReader(bufio.NewReader(file))
	records, err := csvReader.ReadAll()

	return records

}

// function to store result into problem array
func parseProblems(records [][]string) []problem {
	problems := make([]problem, len(records))

	for _, record := range records {
		if len(record) != 2 {
			continue
		}

		//answer, _ := strconv.Atoi(record[1])
		//fmt.Println(answer)
		//
		//inputs := strings.Split(record[0], "+")
		//sum := 0
		//for _, num := range inputs {
		//	num, _ := strconv.Atoi(num)
		//	sum += num
		//}

		problem := problem{
			q: record[0],
			a: record[1],
		}
		problems = append(problems, problem)
	}
	return problems
}

func main() {
	// timer
	timer := time.NewTimer(time.Duration(5) * time.Second)
	correct, incorrect := 0, 0
	// main logic to timesout after 30 seconds.
	records := readCSVFile("inputs.csv")
	problems := parseProblems(records)

	answerCh := make(chan string)

	for i, p := range problems {
		fmt.Printf("Problem #%d: %s ", i+1, p.q)
		go func() {
			var answer string
			fmt.Scanf("%s\n", &answer)
			answerCh <- answer
		}()

		select {
		case <-timer.C:
			fmt.Printf("\nCorrect %d, Incorrect %d, Total %d\n", correct, incorrect, len(problems))
			return
		case answer := <-answerCh:
			if answer == p.a {
				correct++
			} else {
				incorrect++
			}
		}

	}

	fmt.Printf("\nCorrect %d, Incorrect %d, Total %d\n", correct, incorrect, len(problems))
}
