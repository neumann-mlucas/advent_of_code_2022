package main

import (
	"fmt"
	"io/ioutil"
	"os"
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
	var score int
	games := strings.Split(string(content), "\n")
	for _, rawGame := range games {
		game, err := parseGame(rawGame)
		if err == nil {
			score += game.winScore()
			score += game.choiceScore()
		}
	}
	fmt.Println("part one:\t", score)

	score = 0
	for _, rawGame := range games {
		game, err := parseGame(rawGame)
		if err == nil {
			game.followStrategy(rawGame)
			score += game.winScore()
			score += game.choiceScore()
		}
	}
	fmt.Println("part two:\t", score)
}

type Choice uint8

const (
	Rock     Choice = 0
	Paper    Choice = 1
	Scissors Choice = 2
)

type Game struct {
	p1 Choice
	p2 Choice
}

func parseGame(s string) (*Game, error) {
	choices := strings.Split(s, " ")
	if len(choices) != 2 {
		return nil, fmt.Errorf("Bad input: %s", s)
	}

	game := new(Game)
	var err error

	game.p1, err = parseChoice(choices[0])
	if err != nil {
		return nil, err
	}
	game.p2, err = parseChoice(choices[1])
	if err != nil {
		return nil, err
	}
	return game, nil
}

func parseChoice(c string) (Choice, error) {
	if c == "A" || c == "X" {
		return Rock, nil
	} else if c == "B" || c == "Y" {
		return Paper, nil
	} else if c == "C" || c == "Z" {
		return Scissors, nil
	}
	return 0, fmt.Errorf("Cannot Parse %s", c)
}

func inc(h Choice) Choice {
	return (h + 1) % 3
}

func dec(h Choice) Choice {
	return (h - 1) % 3
}

func (g *Game) winScore() int {
	if loseTo(g.p1) == g.p2 {
		return 6
	} else if g.p1 == g.p2 {
		return 3
	} else {
		return 0
	}
}

func (g *Game) choiceScore() int {
	return int(g.p2) + 1
}

func (g *Game) followStrategy(s string) {
	if strings.Contains(s, "Z") {
		g.makeWin()
	} else if strings.Contains(s, "Y") {
		g.makeDraw()
	} else {
		g.makeLoss()
	}
}

func (g *Game) makeDraw() {
	g.p2 = g.p1
}

func (g *Game) makeWin() {
	g.p2 = loseTo(g.p1)
}

func (g *Game) makeLoss() {
	g.p2 = loseTo(loseTo(g.p1))
}

func loseTo(c Choice) Choice {
	switch c {
	case Rock:
		return Paper
	case Paper:
		return Scissors
	case Scissors:
		return Rock
	default:
		return 0
	}
}
