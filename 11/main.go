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

	groups := strings.Split(string(c), "\n")

	for _, g := range groups {

		steps := strings.Split(g, ",")
		fmt.Println(findMinimalSteps(steps))
	}

}

var moveMap = map[string]func(int, int) (int, int){
	"n":  func(x, y int) (int, int) { return x, y + 1 },
	"s":  func(x, y int) (int, int) { return x, y - 1 },
	"ne": func(x, y int) (int, int) { return x + 1, y },
	"sw": func(x, y int) (int, int) { return x - 1, y },
	"se": func(x, y int) (int, int) { return x + 1, y - 1 },
	"nw": func(x, y int) (int, int) { return x - 1, y + 1 },
}

func findMinimalSteps(steps []string) int {

	x, y := 0, 0
	furthest := 0
	for _, s := range steps {
		x, y = moveMap[s](x, y)
		if getDistance(x, y) > furthest {
			furthest = getDistance(x, y)
		}
	}

	fmt.Println(x, y)

	fmt.Println(furthest)
	return getDistance(x, y)
}

func getDistance(x, y int) int {
	total := 0
	if x > 0 && y > 0 {
		total = x + y
	}
	if x < 0 && y < 0 {
		total = -1 * (x + y)
	}
	if x < 0 && y > 0 {
		total = -x + y
	}
	if x > 0 && y < 0 {
		total = x
	}

	if y == 0 {
		total = x
	}
	if x == 0 {
		total = y
	}

	return total
}
