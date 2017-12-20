package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {
	f, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}

	c, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}

	grid := [][]rune{}
	for _, r := range strings.Split(string(c), "\n") {
		row := []rune{}
		for _, c := range r {
			row = append(row, c)
		}
		grid = append(grid, row)
	}

	// for _, r := range grid {
	// 	fmt.Println(string(r))
	// }

	direction := "d"
	// crawl along
	x, y := 0, 0
	// get entry point
	for xx, c := range grid[0] {
		if c == '|' {
			y, x = 0, xx
		}
	}
	fmt.Println(x, y)
	fmt.Println(string(grid[y][x]))
	storedLetters := []rune{}

	totalSteps := 0
	for {
		switch direction {
		case "d":
			y = y + 1
		case "u":
			y = y - 1
		case "l":
			x = x - 1
		case "r":
			x = x + 1
		}

		totalSteps++
		currentSymbol := grid[y][x]

		// chech neighbors of this point to see in which direction to move
		// ignore the direction you came from
		if currentSymbol != '+' {
			// check if letter
			if currentSymbol != '|' && currentSymbol != '-' && currentSymbol != ' ' {
				storedLetters = append(storedLetters, currentSymbol)
				fmt.Println("current symbol", string(currentSymbol), currentSymbol, totalSteps)
			}
			continue
		}

		// fmt.Println("Turning", y, x)

		// come upon +, check where to go now
		if string(grid[y+1][x]) != " " && direction != "d" && direction != "u" {
			direction = "d"
			continue
		}
		if string(grid[y-1][x]) != " " && direction != "u" && direction != "d" {
			direction = "u"
			continue
		}
		if string(grid[y][x-1]) != " " && direction != "l" && direction != "r" {
			direction = "l"
			continue
		}
		if string(grid[y][x+1]) != " " && direction != "r" && direction != "l" {
			direction = "r"
			continue
		}

	}

	// track moving direction
	// Check neighbours, starting in traveling direction
}
