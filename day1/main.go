package main

import (
	"bufio"
	"fmt"
	"os"
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

func elfMostCalories(elfCaloriesTotal map[string]int) map[string]int {
	bossElf := ""
	maxCalories := 0
	for elf, calories := range elfCaloriesTotal {
		if calories > maxCalories {
			bossElf = elf
			maxCalories = calories
		}
	}

	elf := make(map[string]int)
	elf[bossElf] = maxCalories
	return elf
}

func main() {
	fmt.Println(elfMostCalories(elfsTotalCalories(readFile("input.txt"))))
}
