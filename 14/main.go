package main

import (
	"encoding/hex"
	"fmt"
	"strconv"
)

func main() {
	grid := getGrid()

	// fmt.Println(grid[0])
	// fmt.Println(grid[1])
	// os.Exit(0)
	// count all filled
	// cn := 0
	// for _, r := range grid {
	// 	for _, c := range r {
	// 		if c {
	// 			cn++
	// 		}

	// 	}
	// }

	// fmt.Println(cn)

	groupCells := &[]*string{}
	groupCount := 0
	// walk through all cells and try to climb around
	for curRow, r := range grid {
		for curCol, c := range r {
			// create climber that tries to crawl through members
			ident := strconv.Itoa(curRow) + "-" + strconv.Itoa(curCol)
			if c && !alreadyNoted(&ident, groupCells) {
				// new group
				groupCount++
				fmt.Println("Crawling group " + strconv.Itoa(groupCount))
				crawl(curRow, curCol, "", grid, groupCells)
			}

		}
	}
	fmt.Println("done")

}

func crawl(x, y int, ignoreDir string, grid [][]bool, groupCells *[]*string) {
	ident := strconv.Itoa(x) + "-" + strconv.Itoa(y)
	if alreadyNoted(&ident, groupCells) {
		return
	}
	*groupCells = append(*groupCells, &ident)
	if x != 0 && grid[x-1][y] && ignoreDir != "u" {
		crawl(x-1, y, "d", grid, groupCells)
	}
	if x != 127 && grid[x+1][y] && ignoreDir != "d" {
		crawl(x+1, y, "u", grid, groupCells)
	}
	if y != 0 && grid[x][y-1] && ignoreDir != "r" {
		crawl(x, y-1, "l", grid, groupCells)
	}
	if y != 127 && grid[x][y+1] && ignoreDir != "l" {
		crawl(x, y+1, "r", grid, groupCells)
	}
}

func alreadyNoted(in *string, list *[]*string) bool {
	for _, v := range *list {
		if *v == *in {
			return true
		}
	}
	return false
}

func getGrid() [][]bool {
	rows := [][]bool{}
	for i := 0; i < 128; i++ {

		// input := getKnotHash("flqrgnkx-" + strconv.Itoa(i))
		input := getKnotHash("hxtvlmkl-" + strconv.Itoa(i))

		dc, _ := hex.DecodeString(input)
		rowData := ""
		for _, r := range dc {
			rowData = rowData + fmt.Sprintf("%08b", r)
		}

		rarr := []bool{}
		for _, c := range rowData {
			rarr = append(rarr, c == '1')
		}
		rows = append(rows, rarr)
	}
	return rows

}
