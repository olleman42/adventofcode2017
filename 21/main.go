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

	rules := map[string]string{}
	for _, l := range strings.Split(string(c), "\n") {
		comps := strings.Split(l, "=>")
		comps[0] = strings.TrimSpace(comps[0])
		comps[1] = strings.TrimSpace(comps[1])
		rules[comps[0]] = comps[1]
	}

	currentGrid := [][]bool{
		[]bool{false, true, false},
		[]bool{false, false, true},
		[]bool{true, true, true},
	}

	// fmt.Println(gridToString(flip((currentGrid))))
	// fmt.Println(gridToString(rotate(currentGrid)))
	// fmt.Println(gridToString(rotate(rotate(currentGrid))))
	// check if it can be reduced
	// enhance
	// check if grid matches anything - else rotate and flip to fit

	// getVisual(currentGrid)
	// currentGrid = translateGrid(currentGrid, rules)
	for i := 0; i < 18; i++ {

		getVisual(currentGrid)

		// run analysis on each component
		regrid := [][][][]bool{}
		for _, l := range split(currentGrid) {
			ll := [][][]bool{}
			for _, c := range l {
				// fmt.Println("from")
				// getVisual(c)
				// fmt.Println("to")
				// getVisual(translateGrid(c, rules))
				ll = append(ll, translateGrid(c, rules))
			}
			regrid = append(regrid, ll)
		}

		currentGrid = join(regrid)
	}

	getVisual(currentGrid)
	fmt.Println(getCount(currentGrid))

	// split

}

func getCount(in [][]bool) int {
	total := 0
	for _, l := range in {
		for _, v := range l {
			if v {
				total++
			}
		}
	}
	return total
}

func translateGrid(in [][]bool, rules map[string]string) [][]bool {
	flipAlternator := []func(in [][]bool) [][]bool{
		rotate,
		flip,
		flip,
	}

	tracker := -1
	for {
		tracker++
		// fmt.Println(gridToString(in))
		if val, ok := rules[gridToString(in)]; ok {

			return stringToGrid(val)
		}
		in = flipAlternator[tracker%3](in)
	}
}

func getVisual(grid [][]bool) {
	return
	for _, l := range grid {
		for _, c := range l {
			if c {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Print("\n")
	}
	fmt.Print("--\n")
}

func gridToString(grid [][]bool) string {
	lines := []string{}
	for _, l := range grid {
		line := ""
		for _, b := range l {
			val := ""
			if b {
				val = "#"
			} else {
				val = "."
			}
			line = line + val

		}
		lines = append(lines, line)
	}
	return strings.Join(lines, "/")
}

func stringToGrid(in string) [][]bool {
	lines := strings.Split(in, "/")
	olines := [][]bool{}
	for _, l := range lines {
		line := []bool{}
		for _, r := range l {
			if r == '.' {
				line = append(line, false)
			} else {
				line = append(line, true)
			}
		}
		olines = append(olines, line)
	}

	return olines
}

func rotate(ingrid [][]bool) [][]bool {
	size := len(ingrid)
	outGrid := [][]bool{}
	for i := 0; i < size; i++ {
		ln := []bool{}
		for j := size - 1; j >= 0; j-- {
			ln = append(ln, ingrid[j][i])
		}
		outGrid = append(outGrid, ln)
	}
	return outGrid
}

func flip(ingrid [][]bool) [][]bool {
	size := len(ingrid)
	ogrid := [][]bool{}
	for _, l := range ingrid {

		nl := []bool{}
		for i := size - 1; i >= 0; i-- {
			nl = append(nl, l[i])
		}
		ogrid = append(ogrid, nl)
	}
	return ogrid
}

func split(ingrid [][]bool) [][][][]bool {

	splitsize := 0
	if len(ingrid)%2 == 0 {
		splitsize = 2
	}
	if len(ingrid)%3 == 0 && splitsize == 0 {
		splitsize = 3
	}
	if splitsize == 0 {
		log.Fatal("trying to work with grid of impossible size")
	}

	// qudrants
	qsize := len(ingrid) / splitsize
	allGrids := [][][][]bool{}
	for i := 0; i < qsize; i++ {
		aGridLine := [][][]bool{}
		for j := 0; j < qsize; j++ {
			minigrid := [][]bool{}
			for a := 0; a < splitsize; a++ {
				mgline := []bool{}
				for b := 0; b < splitsize; b++ {
					// fmt.Println((i*qsize)+a, (j*qsize)+b, j, qsize, b)
					mgline = append(mgline, ingrid[(i*splitsize)+a][(j*splitsize)+b])
				}

				minigrid = append(minigrid, mgline)
			}
			aGridLine = append(aGridLine, minigrid)
		}
		allGrids = append(allGrids, aGridLine)
	}

	return allGrids

}

func join(grid [][][][]bool) [][]bool {
	// walk through each row and connection that SHYET
	// outsize := len(grid)
	size := len(grid[0][0][0])
	fmt.Println(size)

	oo := [][]bool{}
	for _, l := range grid {
		lines := [][]bool{}

		// prepare lines
		for i := 0; i < size; i++ {
			lines = append(lines, []bool{})
			for _, c := range l {
				lines[i] = append(lines[i], c[i]...)
			}

		}

		oo = append(oo, lines...)

	}

	return oo
}
