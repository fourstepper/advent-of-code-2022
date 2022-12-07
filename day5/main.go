package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func stacksAmount(input string) int {
	return len(strings.ReplaceAll(input, " ", ""))
}

// strips useless text around number commands
func stripIntoCommand(input string) string {
	return strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(input, "move", ""), "from", ""), "to", "")
}

func deleteSpaces(s []string) []string {
	var r []string
	for _, str := range s {
		if str != " " {
			r = append(r, str)
		}
	}
	return r
}

func removeIndex[T any](slice []T, i int) []T {
	return append(slice[:i], slice[i+1:]...)
}

func mapSlice0thItem(input map[int][]string) string {
	var keys []int
	var r string
	for key := range input {
		keys = append(keys, key)
	}
	sort.Ints(keys)

	for _, i := range keys {
		r = r + input[i][0]
	}
	return r
}

func parseStacks(input []string) map[int][]string {
	mapping := make(map[int][]string)
	final := make(map[int][]string)

	whitespaceCounter := 0

	// for each line in the stacks map
	for idx, stack := range input {
		var tempSlice []string
		// for each char in the line
		for _, char := range stack {
			if whitespaceCounter >= 3 {
				tempSlice = append(tempSlice, " ")
				whitespaceCounter = 0
				continue
			}
			if string(char) == "[" || string(char) == "]" {
				continue
			}
			if string(char) == " " {
				whitespaceCounter++
				continue
			}
			tempSlice = append(tempSlice, string(char))
			whitespaceCounter = 0
		}
		mapping[idx] = tempSlice
	}

	var keys []int
	for key := range mapping {
		keys = append(keys, key)
	}
	sort.Ints(keys)

	stacksAmount := len(mapping[1])

	for i := 0; i < stacksAmount; i++ {
		var tempSlice []string

		for _, value := range keys {
			tempSlice = append(tempSlice, mapping[value][i])
		}
		final[i+1] = tempSlice
	}

	return final
}

func readFile(file string) (string, string) {
	readFile, err := os.Open(file)
	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	defer readFile.Close()

	var totalFirst string
	totalSecond := ""

	// first task
	// for each stack, create a key in the map

	var stackStrings []string
	var mapping map[int][]string

	for fileScanner.Scan() {
		prefix := strings.HasPrefix(fileScanner.Text(), " 1")
		movePrefix := strings.HasPrefix(fileScanner.Text(), "move")

		// get a list of the stack lines for parsing
		if !prefix {
			stackStrings = append(stackStrings, fileScanner.Text())

		}
		// get a parsed map
		if prefix {
			mapping = parseStacks(stackStrings)
			for i := range mapping {
				mapping[i] = deleteSpaces(mapping[i])
			}
		}
		// if a line starts with move, process command
		if movePrefix {
			command := stripIntoCommand(fileScanner.Text())

			var commandSlice []int
			trimmed := strings.Split(strings.TrimSpace(command), " ")

			for _, i := range trimmed {
				Int, err := strconv.Atoi(i)
				if err != nil {
					// skip anything that can't be converted to int
					continue
				}
				commandSlice = append(commandSlice, Int)
			}

			for i := 0; i < commandSlice[0]; i++ {
				if len(mapping[commandSlice[1]]) == 0 {
					break
				}
				mapping[commandSlice[2]] = append([]string{mapping[commandSlice[1]][0]}, mapping[commandSlice[2]]...)
				mapping[commandSlice[1]] = removeIndex(mapping[commandSlice[1]], 0)
			}
		}
	}
	// append the first item of each slice in map to string, in order
	totalFirst = mapSlice0thItem(mapping)

	return totalFirst, totalSecond
}

func main() {
	fmt.Println(readFile("input.txt"))
}
