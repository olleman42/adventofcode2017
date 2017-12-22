package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

var (
	Infected        = "infected"
	Flagged         = "flagged"
	Weakened        = "weakened"
	Clean           = "clean"
	statusTransform = map[string]string{
		Clean:    Weakened,
		Weakened: Infected,
		Infected: Flagged,
		Flagged:  Clean,
	}
	infectedCount = 0
	turnRight     = map[string]string{"u": "r", "r": "d", "d": "l", "l": "u"}
	turnLeft      = map[string]string{"u": "l", "r": "u", "d": "r", "l": "d"}
	turnBack      = map[string]string{"u": "d", "l": "r", "r": "l", "d": "u"}
)

type node struct {
	status string
}

func (n *node) changeDirection(currentDirection *string) {
	if n.status == Infected {
		*currentDirection = turnRight[*currentDirection]
	}
	if n.status == Clean {
		*currentDirection = turnLeft[*currentDirection]
	}
	if n.status == Flagged {
		*currentDirection = turnBack[*currentDirection]
	}

	n.changeStatus()
}

func (n *node) changeStatus() {
	n.status = statusTransform[n.status]
	// fmt.Println("Became ", n.status)
	if n.status == Infected {
		infectedCount++
	}
}

func main() {

	f, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}

	c, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}

	indata := [][]node{}
	for _, l := range strings.Split(string(c), "\n") {
		line := []node{}
		for _, c := range l {
			if c == '#' {
				line = append(line, node{status: Infected})
				continue
			}
			line = append(line, node{status: Clean})

		}
		indata = append(indata, line)
	}

	// set up imaginary coordinate system
	// get middle point of input and set as origo
	size := len(indata)

	half := int(math.Floor(float64(size / 2)))

	fmt.Println(half)

	// fill imaginary grid
	grid := map[string]*node{}
	r := 0
	cx := 0
	for x := -half; x <= half; x++ {
		for y := -half; y <= half; y++ {
			// fmt.Println(r, cx)
			grid[coordToString(x, y)] = &indata[r][cx]
			cx++
		}
		r++
		cx = 0
	}

	x, y := 0, 0

	direction := "u"

	for i := 0; i < 10000000; i++ {

		if _, ok := grid[coordToString(x, y)]; !ok {
			grid[coordToString(x, y)] = &node{status: Clean}
		}

		n := grid[coordToString(x, y)]
		n.changeDirection(&direction)

		// advance
		switch direction {
		case "u":
			x--
		case "d":
			x++
		case "l":
			y--
		case "r":
			y++
		}
		// fmt.Println(x, y)
	}
	fmt.Println(infectedCount)

}

func coordToString(x, y int) string {
	return fmt.Sprintf("%d,%d", x, y)
}

func stringToCoord(in string) (int, int) {
	psd := strings.Split(in, ",")
	x, err := strconv.Atoi(psd[0])
	if err != nil {
		log.Fatal(err)
	}
	y, err := strconv.Atoi(psd[1])
	if err != nil {
		log.Fatal(err)
	}
	return x, y
}
