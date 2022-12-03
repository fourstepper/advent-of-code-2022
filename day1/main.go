package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func readFile(file string) map[string][]string {
	readFile, err := os.Open(file)
	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	elfs := make(map[string][]string)

	elf := []string{}
	elf_number := 0
	for fileScanner.Scan() {
		if len(fileScanner.Text()) == 0 {
			elf = nil
			elf_number++
			continue
		}

		elf = append(elf, fileScanner.Text())
		elfs[strconv.Itoa(elf_number)] = elf

	}

	// fmt.Println(elfs)

	readFile.Close()
	return elfs
}

func elfsTotalCalories(elfFoodPacks map[string][]string) map[string]int {
	elfs := make(map[string]int)
	elf_number := 0
	for _, foodPacks := range elfFoodPacks {
		foodTotal := 0
		for _, foodPack := range foodPacks {
			foodPack, err := strconv.Atoi(foodPack)
			if err != nil {
				fmt.Println(err)
			}

			foodTotal = foodTotal + foodPack
		}
		elfs[strconv.Itoa(elf_number)] = foodTotal
		elf_number++
	}
	return elfs
}

func elfsSortCalories(elfCaloriesTotal map[string]int) map[string]int {
	keys := make([]string, 0, len(elfCaloriesTotal))

	for key := range elfCaloriesTotal {
		keys = append(keys, key)
	}

	sort.Slice(keys, func(i, j int) bool {
		return elfCaloriesTotal[keys[i]] > elfCaloriesTotal[keys[j]]
	})

	bossElfs := make(map[string]int)
	iter := 0
	for _, key := range keys {
		if iter < 3 {
			bossElfs[key] = elfCaloriesTotal[key]
		}
		iter++
	}

	total := 0
	for _, value := range bossElfs {
		total = total + value
	}
	bossElfs["total"] = total

	return bossElfs
}

func main() {
	// prints out a map[string]int with the top three elfs, as well as the total value of calories those elfs have
	fmt.Println(elfsSortCalories(elfsTotalCalories(readFile("input.txt"))))
}
