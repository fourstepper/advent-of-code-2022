// Credit for this day goes to Satyarth Agrahari
// As I am a beginner to both Go and data structures, I used a lot of his code from his solution at https://github.com/satylogin/aoc/blob/main/2022/day_07.go to implement my solution

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type File struct {
	isDir    bool
	size     int
	parent   *File
	children map[string]*File
}

// converts a string to int
func StrToInt(s string) (i int) {
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Panic(err)
	}
	return
}

func CreateDir(parent *File) *File {
	return &File{
		isDir:    true,
		parent:   parent,
		children: map[string]*File{},
	}
}

func ReadLines(scanner *bufio.Scanner) {
	objects := []*File{}
	pos := map[*File]int{}

	root := CreateDir(nil)
	root.parent = root
	node := root

	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		switch fields[0] {
		case "$":
			switch fields[1] {
			case "cd":
				switch fields[2] {
				case "/":
					node = root
				case "..":
					node = node.parent
				default:
					node = node.children[fields[2]]
				}
			case "ls":
			}
		case "dir":
			node.children[fields[1]] = CreateDir(node)
			pos[node]++
		default:
			node.children[fields[1]] = &File{
				isDir:  false,
				size:   StrToInt(fields[0]),
				parent: node,
			}
			objects = append(objects, node.children[fields[1]])
			pos[node]++
		}
	}
	var sizes []int
	for i := 0; i < len(objects); i++ {
		file := objects[i]
		parent := file.parent
		if parent != file {
			parent.size += file.size
			pos[parent] -= 1
			if pos[parent] == 0 {
				objects = append(objects, parent)
			}
		}
		if file.isDir {
			sizes = append(sizes, file.size)
		}
	}

	sort.Ints(sizes)

	// task 1
	var res1 int
	for _, size := range sizes {
		if size <= 100000 {
			res1 += size
		}
	}
	fmt.Printf("Task 1 result: %v\n", res1)

	// task2
	fsTotal := 70000000 - sizes[len(sizes)-1]
	requiredSpace := 30000000
	var res2 int
	for _, size := range sizes {
		if size+fsTotal >= requiredSpace {
			res2 = size
			break
		}
	}
	fmt.Printf("Task 2 result: %v\n", res2)
}

var input_files = [1]string{"input.txt"}

func main() {
	for _, file := range input_files {
		input, err := os.Open(file)
		if err != nil {
			log.Panic(err)
		}

		scanner := bufio.NewScanner(input)
		ReadLines(scanner)
	}
}
