package main

import (
	"fmt"
	"io/ioutil"
	"os"
	// "strconv"
	// "strings"
)

const (
	testInput = "mjqjpqmgbljsphdztnvjfqwrcgsmlb"
)

func main() {
	resultTest := firstProblem(testInput)
	if resultTest != 7 {
		fmt.Fprintf(os.Stderr, "fail test on first problem: %d != %d", resultTest, 7)
		os.Exit(0)
	}

	resultTest = secondProblem(testInput)
	if resultTest != 19 {
		fmt.Fprintf(os.Stderr, "fail test on first problem: %d != %d", resultTest, 19)
		os.Exit(0)
	}

	if len(os.Args) != 2 {
		fmt.Fprint(os.Stderr, "error: wrong number of arguments")
	}
	content, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s", err)
	}

	sumPriorities := firstProblem(string(content))
	fmt.Println("part one:\t", sumPriorities)

	sumGroupPriorities := secondProblem(string(content))
	fmt.Println("part two:\t", sumGroupPriorities)
}

func secondProblem(inp string) int {
	for i := 0; i < len(inp)-14; i++ {
		set := genSet(inp[i : i+14])
		if lenSet(set) == 14 {
			return i + 14
		}
	}
	return -1
}

func firstProblem(inp string) int {
	for i := 0; i < len(inp)-4; i++ {
		set := genSet(inp[i : i+4])
		if lenSet(set) == 4 {
			return i + 4
		}
	}
	return -1
}

func genSet(str string) map[rune]bool {
	set := make(map[rune]bool)
	for _, c := range str {
		set[c] = true
	}
	return set
}

func lenSet(set map[rune]bool) int {
	var length int
	for _, v := range set {
		if v {
			length++
		}
	}
	return length
}
