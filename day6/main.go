package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

// convert a string to a []string of it's characters
func stringToSlice(s string) (sl []string) {
	for _, char := range s {
		charString := strings.TrimSpace(string(char))
		sl = append(sl, string(charString))
	}
	return
}

// check if all items in []string are unique
func Unique(s []string) bool {
	var tempSlice []string

	for _, item := range s {
		for _, tempItem := range tempSlice {
			if item == tempItem {
				return false
			}
		}
		tempSlice = append(tempSlice, item)
	}
	return true
}

// removes an item from a slice (sl) at index (idx)
func RemoveIndex(sl []string, idx int) []string {
	return append(sl[:idx], sl[idx+1:]...)
}

// read file and return it's contents as string
func ReadFile(path string) (s string) {
	fileContent, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	s = string(fileContent)

	return
}

func Task(path string, distinctChars int) (result int) {
	inputSlice := stringToSlice(ReadFile(path))
	var tempSlice []string

	for idx, marker := range inputSlice {

		// add to tempSlice
		tempSlice = append(tempSlice, marker)

		if len(tempSlice) != distinctChars {
			continue
		} else {
			if Unique(tempSlice) {
				return idx + 1
			}
			// remove first index of slice to accomodate the next one
			tempSlice = RemoveIndex(tempSlice, 0)
		}
	}
	return
}

func main() {
	fmt.Println(Task("input.txt", 4))
	fmt.Println(Task("input.txt", 14))
}
