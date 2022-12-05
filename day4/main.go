package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func splitFour(input string) []int {
	var slice []int

	halfs := strings.Split(input, ",")

	for _, half := range halfs {
		quarters := strings.Split(half, "-")
		for _, quarter := range quarters {
			quarterInt, err := strconv.Atoi(quarter)
			if err != nil {
				fmt.Println(err)
			}
			slice = append(slice, quarterInt)
		}
	}

	return slice
}

// returns a slice containing all ints between the first and second int
func getRangeSlice(first, second int) []int {
	var finalSlice []int
	for i := first; i <= second; i++ {
		finalSlice = append(finalSlice, i)
	}
	return finalSlice
}

// calls sort.Ints on input []int and returns a sorted []int
func sortIntSlice(input []int) []int {
	sort.Ints(input)
	return input
}

func readFile(file string) (int, int) {
	readFile, err := os.Open(file)
	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	defer readFile.Close()

	totalFirst := 0
	totalSecond := 0

	for fileScanner.Scan() {
		// first task
		line := splitFour(fileScanner.Text())

		if (line[0] <= line[2] && line[1] >= line[3]) || (line[2] <= line[0] && line[3] >= line[1]) {
			totalFirst++
		}

		Slice1 := sortIntSlice(getRangeSlice(line[0], line[1]))
		Slice2 := sortIntSlice(getRangeSlice(line[2], line[3]))

		for _, Int := range Slice1 {
			// Check at what index the int could be inserted into the slice so the slice would remain sorted
			searchIntsResult := sort.SearchInts(Slice2, Int)
			// Check if the resulting index is lower than the total length of the current slice
			if searchIntsResult < len(Slice2) {
				// If it is lower, check if the index position already contains the int we are looking for
				// If it does, add to the match counter and break out of the loop
				if Slice2[searchIntsResult] == Int {
					totalSecond++
					break
				}
			}
		}
	}

	return totalFirst, totalSecond
}

func main() {
	fmt.Println(readFile("input.txt"))
}
