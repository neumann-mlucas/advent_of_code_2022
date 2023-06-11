package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprint(os.Stderr, "error: wrong number of arguments")
	}
	content, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s", err)
	}

	var calories []int
	for _, inv := range strings.Split(string(content), "\n\n") {
		calories = append(calories, caloriesInInventory(inv))
	}
	sort.Ints(calories)

	fmt.Printf("most calories:\t%d cals\n", calories[len(calories)-1])
	var topThree int
	for i := len(calories) - 1; i > len(calories)-4 && i >= 0; i-- {
		topThree += calories[i]
	}
	fmt.Printf("sum of top 3:\t%d cals\n", topThree)
}

func caloriesInInventory(bag string) int {
	var total int
	for _, item := range strings.Split(bag, "\n") {
		calories, _ := strconv.Atoi(item)
		total += calories
	}
	return total
}
