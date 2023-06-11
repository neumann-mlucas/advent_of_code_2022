package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

const (
	testInput = "2-4,6-8\n2-3,4-5\n5-7,7-9\n2-8,3-7\n6-6,4-6\n2-6,4-8"
)

func main() {
	resultTest := firstProblem(testInput)
	if resultTest != 2 {
		fmt.Fprintf(os.Stderr, "fail test on first problem: %d != %d", resultTest, 2)
		os.Exit(0)
	}

	resultTest = secondProblem(testInput)
	if resultTest != 4 {
		fmt.Fprintf(os.Stderr, "fail test on first problem: %d != %d", resultTest, 4)
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

func firstProblem(input string) int {
	var sum int
	for _, duo := range strings.Split(input, "\n") {
		if duo == "" {
			continue
		}
		elfs := strings.Split(duo, ",")
		elfA := genSet(genRange(elfs[0]))
		elfB := genSet(genRange(elfs[1]))
		if isSubSet(elfA, elfB) || isSubSet(elfB, elfA) {
			sum++
		}

	}
	return sum
}

func secondProblem(input string) int {
	var sum int
	for _, duo := range strings.Split(input, "\n") {
		if duo == "" {
			continue
		}
		elfs := strings.Split(duo, ",")
		elfA := genSet(genRange(elfs[0]))
		elfB := genSet(genRange(elfs[1]))
		if hasOverlap(elfA, elfB) {
			sum++
		}
	}
	return sum
}

func genRange(str string) []int {
	interval := strings.Split(str, "-")
	lower, _ := strconv.Atoi(interval[0])
	upper, _ := strconv.Atoi(interval[1])
	var result []int
	for i := lower; i <= upper; i++ {
		result = append(result, i)
	}
	return result
}

func genSet(s []int) map[int]bool {
	set := make(map[int]bool)
	for _, c := range s {
		set[c] = true
	}
	return set
}

func hasOverlap(a, b map[int]bool) bool {
	for k, v := range b {
		if _, ok := a[k]; ok && v {
			return true
		}
	}
	return false
}

func isSubSet(a, b map[int]bool) bool {
	for k, v := range b {
		if _, ok := a[k]; !ok || !v {
			return false
		}
	}
	return true
}
