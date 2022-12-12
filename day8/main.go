package main

import (
	_ "embed"
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"
)

func stringToInt(s string) (i int) {
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Panic(err)
	}
	return
}

func splitByNewline(s string) []string {
	return strings.Split(strings.TrimSpace(s), "\n")
}

func getCharsSlice(s string) (sl []string) {
	for _, char := range s {
		sl = append(sl, string(char))
	}
	return
}

// check if the first is higher than the rest
func canSee(check int, sl []int) bool {
	if len(sl) == 0 {
		return true
	}

	for _, i := range sl {
		if i >= check {
			return false
		}
	}
	return true
}

func viewingDistance(check int, sl []int) (distance int) {
	if len(sl) == 0 {
		return 0
	}

	for _, i := range sl {
		distance++
		if i >= check {
			return distance
		}
	}
	return
}

func reverseSlice(input []int) []int {
	var output []int

	for i := len(input) - 1; i >= 0; i-- {
		output = append(output, input[i])
	}

	return output
}

type Tree struct {
	value                                    int
	pos                                      []int
	fromTop, fromBottom, fromLeft, fromRight bool
	scsTop, scsBottom, scsLeft, scsRight     int
}

func solve(input string) {
	treeMapRows := make(map[int][]int)
	treeMapColumns := make(map[int][]int)
	Trees := make(map[string]Tree)

	for idx, item := range splitByNewline(input) {
		for _, tree := range getCharsSlice(item) {
			treeMapRows[idx] = append(treeMapRows[idx], stringToInt(tree))
		}
	}

	for _, item := range splitByNewline(input) {
		for idx, tree := range getCharsSlice(item) {
			treeMapColumns[idx] = append(treeMapColumns[idx], stringToInt(tree))
		}
	}

	var treeMapRowsKeys []int
	for key := range treeMapRows {
		treeMapRowsKeys = append(treeMapRowsKeys, key)
	}
	sort.Ints(treeMapRowsKeys)

	var treeMapColumnsKeys []int
	for key := range treeMapColumns {
		treeMapColumnsKeys = append(treeMapColumnsKeys, key)
	}
	sort.Ints(treeMapColumnsKeys)

	for idxMap, key := range treeMapRowsKeys {
		for idxList, tree := range treeMapRows[key] {
			fromLeft := canSee(tree, treeMapRows[key][idxList+1:])
			fromRight := canSee(tree, treeMapRows[key][:idxList])

			scsLeft := viewingDistance(tree, treeMapRows[key][idxList+1:])
			scsRight := viewingDistance(tree, reverseSlice(treeMapRows[key][:idxList]))

			idx := "row" + fmt.Sprint(idxMap) + "column" + fmt.Sprint(idxList)
			Trees[idx] = Tree{
				value:     tree,
				pos:       []int{idxMap, idxList},
				fromLeft:  fromLeft,
				fromRight: fromRight,
				scsLeft:   scsLeft,
				scsRight:  scsRight,
			}
		}
	}
	for idxMap, key := range treeMapColumnsKeys {
		for idxList, tree := range treeMapColumns[key] {
			fromTop := canSee(tree, reverseSlice(treeMapColumns[idxMap][idxList+1:]))
			fromBottom := canSee(tree, treeMapColumns[idxMap][:idxList])

			scsTop := viewingDistance(tree, treeMapColumns[idxMap][idxList+1:])
			scsBottom := viewingDistance(tree, reverseSlice(treeMapColumns[idxMap][:idxList]))

			idx := "row" + fmt.Sprint(idxList) + "column" + fmt.Sprint(idxMap)
			if entry, ok := Trees[idx]; ok {
				entry.fromBottom = fromBottom
				entry.fromTop = fromTop
				entry.scsTop = scsTop
				entry.scsBottom = scsBottom
				Trees[idx] = entry
			}
		}
	}
	// taks 1
	visible := 0
	for _, tree := range Trees {
		if !tree.fromBottom && !tree.fromTop && !tree.fromLeft && !tree.fromRight {
			continue
		}
		visible++
	}
	fmt.Println(visible)
	// taks 2
	topScenicScore := 0
	for _, tree := range Trees {
		scenicScore := tree.scsTop * tree.scsBottom * tree.scsLeft * tree.scsRight
		if scenicScore > topScenicScore {
			topScenicScore = scenicScore
		}
	}
	fmt.Println(topScenicScore)
}

//go:embed input.txt
var input_file string

func main() {
	solve(input_file)
}
