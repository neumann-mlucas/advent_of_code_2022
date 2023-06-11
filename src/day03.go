package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

const (
	testInput = "vJrwpWtwJgWrhcsFMMfFFhFp\njqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL\nPmmdzqPrVvPwwTWBwg\nwMqvLMZHhHMvwLHjbvcjnnSBnvTQFn\nttgJtRGJQctTZtZT\nCrZsJsPPZsGzwwsLwLmpwMDw"
)

func main() {
	resultTest := firstProblem(testInput)
	if resultTest != 157 {
		fmt.Fprintf(os.Stderr, "fail test on first problem: %d != %d", resultTest, 157)
		os.Exit(0)
	}

	resultTest = secondProblem(testInput)
	if resultTest != 70 {
		fmt.Fprintf(os.Stderr, "fail test on first problem: %d != %d", resultTest, 70)
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
	for _, bag := range strings.Split(input, "\n") {
		comp1, comp2 := genCompartment(bag)
		common := intersection(genSet(comp1), genSet(comp2))
		for k := range common {
			sum += findPriority(k)
		}
	}
	return sum
}

func secondProblem(input string) int {
	var sum int
	bags := strings.Split(input, "\n")
	for i := 0; i < len(bags)-2; i += 3 {
		set := intersection(genSet(bags[i]), genSet(bags[i+1]))
		set = intersection(set, genSet(bags[i+2]))
		for k := range set {
			sum += findPriority(k)
		}
	}
	return sum
}

func genCompartment(bag string) (string, string) {
	lc := len(bag) / 2
	return bag[:lc], bag[lc:]
}

func genSet(s string) map[rune]bool {
	set := make(map[rune]bool)
	for _, c := range s {
		set[c] = true
	}
	return set
}

func intersection(a, b map[rune]bool) map[rune]bool {
	result := make(map[rune]bool)
	for k, v := range a {
		if _, ok := b[k]; ok && v {
			result[k] = true
		}
	}
	return result
}

func findPriority(r rune) int {
	if r >= 65 && r <= 96 {
		return int(r-'A') + 27 // A to Z
	} else if r >= 97 && r <= 122 {
		return int(r-'a') + 1 // a to z
	}
	return 0
}
