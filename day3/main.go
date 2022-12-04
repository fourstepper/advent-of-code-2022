package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func intValueOf(letter string) int {
	var alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	return strings.Index(alphabet, letter) + 1
}

func readFile(file string) (int, int) {
	readFile, err := os.Open(file)
	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	defer readFile.Close()

	// first task
	totalItemValue := 0
	// second task
	var tempGroupSlice []string
	totalItemValueGroups := 0

	for fileScanner.Scan() {
		// first task
		value := fileScanner.Text()
		middle := len(value) / 2
		var substrings = []string{value[:middle], value[middle:]}

		// iterate over runes of the first compartment and check which
		// are also present in the second compartment
		// Break once we find the first match
		for _, char := range substrings[0] {
			if strings.Contains(substrings[1], string(char)) {
				totalItemValue = totalItemValue + intValueOf(string(char))
				break
			}
		}

		// second task
		tempGroupSlice = append(tempGroupSlice, value)

		// once our tempGroupSlice slice reaches the length of 3, iterate over the chars of the first string and check if any
		// of the characters are present in both the second and the third string
		// if it does, add it's int value to the total group value and nil the slice
		if len(tempGroupSlice) == 3 {
			for _, char := range tempGroupSlice[0] {
				if strings.Contains(tempGroupSlice[1], string(char)) && strings.Contains(tempGroupSlice[2], string(char)) {
					totalItemValueGroups = totalItemValueGroups + intValueOf(string(char))
					tempGroupSlice = nil
					break
				}
			}
		}

	}
	return totalItemValue, totalItemValueGroups
}

func main() {
	fmt.Println(readFile("input.txt"))
}
