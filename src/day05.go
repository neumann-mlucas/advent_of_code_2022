package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

const (
	testInput = `
    [D]    
[N] [C]    
[Z] [M] [P]
 1   2   3 

move 1 from 2 to 1
move 3 from 1 to 3
move 2 from 2 to 1
move 1 from 1 to 2`
)

func main() {
	resultTest := firstProblem(testInput)
	if resultTest != "CMZ" {
		fmt.Fprintf(os.Stderr, "fail test on first problem: %s != %s", resultTest, "CMZ")
		os.Exit(0)
	}

	resultTest = secondProblem(testInput)
	if resultTest != "MCD" {
		fmt.Fprintf(os.Stderr, "fail test on first problem: %s != %s", resultTest, "MCD")
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

type Move struct {
	qty int
	src int
	dst int
}

func firstProblem(input string) string {
	stacks, moves := parseInput(input)
	for _, move := range moves {
		stacks = applyMove(stacks, move)
	}
	result := ""
	for _, stack := range stacks {
		result += string(stack[len(stack)-1])

	}
	return result
}

func secondProblem(input string) string {
	stacks, moves := parseInput(input)
	for _, move := range moves {
		stacks = applyMove9001(stacks, move)

	}
	result := ""
	for _, stack := range stacks {
		result += string(stack[len(stack)-1])

	}
	return result
}

func parseInput(input string) ([]string, []Move) {
	rawStack := strings.Split(input, "\n\n")[0]
	stacks := parseStacks(rawStack)

	rawMoves := strings.Split(input, "\n\n")[1]
	var moves []Move
	for _, move := range strings.Split(rawMoves, "\n") {
		if move != "" {
			moves = append(moves, parseMove(move))
		}
	}
	return stacks, moves
}

func parseStacks(input string) []string {
	rawStacks := strings.Split(input, "\n")

	revString := make([]string, 100, 1000)
	for j := 1; j < 100; j += 2 {
		for _, rawStack := range rawStacks[:len(rawStacks)-1] {
			if j < len(rawStack) {
				revString[j] += string(rawStack[j])
			}
		}
	}

	var stacks []string
	for _, str := range revString {
		str = strings.TrimSpace(str)
		if str != "" {
			stacks = append(stacks, reverse(str))
		}
	}
	return stacks
}

func parseMove(moveStr string) Move {
	words := strings.Split(moveStr, " ")
	qty, _ := strconv.Atoi(words[1])
	src, _ := strconv.Atoi(words[3])
	dst, _ := strconv.Atoi(words[5])
	return Move{qty, src - 1, dst - 1}
}

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func applyMove(stacks []string, move Move) []string {
	for i := 0; i < move.qty; i++ {
		if len(stacks[move.src]) == 0 {
			break
		}
		ls := len(stacks[move.src])
		item := stacks[move.src][ls-1]
		stacks[move.src] = stacks[move.src][:ls-1]
		stacks[move.dst] += string(item)
	}
	return stacks
}

func applyMove9001(stacks []string, move Move) []string {
	if len(stacks[move.src]) == 0 {
		return stacks
	}
	upper := len(stacks[move.src])
	lower := upper - move.qty
	if lower < 0 {
		lower = 0
	}
	item := stacks[move.src][lower:upper]
	stacks[move.src] = stacks[move.src][:lower]
	stacks[move.dst] += item
	return stacks
}
