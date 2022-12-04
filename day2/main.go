package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Rules struct {
	Tie, Win, Loss string
	Value          int
}

var rules = map[string]Rules{
	"X": {
		"A", "C", "B", 1,
	},
	"Y": {
		"B", "A", "C", 2,
	},
	"Z": {
		"C", "B", "A", 3,
	},
}

func readFile(file string) (first, second int) {
	readFile, err := os.Open(file)
	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	count_first := 0
	count_second := 0
	for fileScanner.Scan() {
		round := strings.Split(fileScanner.Text(), " ")

		if round[0] == rules[round[1]].Tie {
			count_first = count_first + 3 + rules[round[1]].Value
		} else if round[0] == rules[round[1]].Win {
			count_first = count_first + 6 + rules[round[1]].Value
		} else {
			count_first = count_first + rules[round[1]].Value
		}

		if round[1] == "Z" {
			for _, rule := range rules {
				if round[0] == rule.Win {
					count_second = count_second + 6 + rule.Value
				}
			}
		} else if round[1] == "Y" {
			for _, rule := range rules {
				if round[0] == rule.Tie {
					count_second = count_second + 3 + rule.Value
				}
			}
		} else {
			for _, rule := range rules {
				if round[0] == rule.Loss {
					count_second = count_second + rule.Value
				}
			}
		}
	}

	readFile.Close()
	// first number for solution 1, second number for solution 2
	return count_first, count_second
}

func main() {
	fmt.Println(readFile("input.txt"))
}
