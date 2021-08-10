package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"time"
)

func main() {
	csvFile := flag.String("csv", "problems.csv", "Run csv problem from this command")
	timeLimit := flag.Int("limit", 30, "Maximum number of time needed to complete problem")
	shuffle := flag.Bool("shuffle", false, "Shuffle the questions (default 'false')")
	flag.Parse()
	_ = shuffle

	file, err := os.Open(*csvFile)
	if err != nil {
		exit(fmt.Sprintf("Failed to open csv file %s\n", *csvFile))
	}

	r := csv.NewReader(file)

	lines, err := r.ReadAll()
	if err != nil {
		exit("failed to parse provided csv")
	}
	problems := parseLines(lines)
	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)

	fmt.Println(problems)

	correct := 0
problemLoop:
	for i, p := range problems {
		fmt.Printf("Problem #%d : %s = \n ", i+1, p.q)
		answerChan := make(chan string)

		go func() {
			var answer string
			fmt.Scanf("%s\n", &answer)
			answerChan <- answer
		}()

		select {
		case <-timer.C:
			fmt.Printf("You scored %d out of %d\n", correct, len(problems))
			break problemLoop
		case answer := <-answerChan:
			if answer == p.a {
				correct++
			}

		}

	}
	fmt.Printf("You scored %d out of %d.", correct, len(problems))

	if *shuffle {
		shuffles()
	}
}

func parseLines(lines [][]string) []problem {
	ret := make([]problem, len(lines))
	for i, line := range lines {
		ret[i] = problem{
			q: line[0],
			a: line[1],
		}
	}
	return ret
}

func shuffles() {
	fmt.Println("Show")

}

type problem struct {
	q string
	a string
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
