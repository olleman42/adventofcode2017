package main

import (
	"encoding/hex"
	"fmt"
	"strconv"
)

type gcs struct {
	s []string
}

func (gcs *gcs) Append(in string) {
	gcs.s = append(gcs.s, in)
}

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

	// groupCells := []*string{}
	gc := &gcs{}
	groupCount := 0
	// walk through all cells and try to climb around
	for curRow, r := range grid {
		for curCol, c := range r {
			// create climber that tries to crawl through members
			ident := strconv.Itoa(curRow) + "-" + strconv.Itoa(curCol)
			if c && !alreadyNoted(&ident, gc) {
				// new group
				groupCount++
				fmt.Println("Crawling group"+strconv.Itoa(groupCount), len(gc.s))
				crawl(curRow, curCol, 0, "", grid, gc)
			}
			//no new group, keep going

			// // if val, check neighbours to see if a new region is to be started - check if neighbours are in group cells already
			// if c {
			// 	// check for neighbors above below, left, right
			// 	// ABOVE
			// 	topHit := curRow != 0 && grid[curRow-1][curCol]
			// 	// LEFT
			// 	leftHit := curCol != 0 && grid[curRow][curCol-1]

			// 	//RIGHT
			// 	rightHit := curCol != 127 && grid[curRow][curCol+1]
			// 	// there is a dude to the right of me is part of my group
			// 	// if neight me or my neighbour are accounted for (and top and right didn't hit), we might be new

			// 	//BOTTOM
			// 	bottomHit := curRow != 127 && grid[curRow+1][curCol]

			// 	// if top and left didn't hit, we might be new
		}
	}
	fmt.Println("done")

}

func crawl(x, y, d int, ignoreDir string, grid [][]bool, groupCells *gcs) {
	// fmt.Println(len(groupCells))
	ident := strconv.Itoa(x) + "-" + strconv.Itoa(y)
	if alreadyNoted(&ident, groupCells) {
		return
	}
	groupCells.Append(ident)
	// groupCells = append(groupCells, &ident)
	// check my neighbours
	if x != 0 && grid[x-1][y] && ignoreDir != "u" {
		crawl(x-1, y, d+1, "d", grid, groupCells)
	}
	if x != 127 && grid[x+1][y] && ignoreDir != "d" {
		crawl(x+1, y, d+1, "u", grid, groupCells)
	}
	if y != 0 && grid[x][y-1] && ignoreDir != "r" {
		crawl(x, y-1, d+1, "l", grid, groupCells)
	}
	if y != 127 && grid[x][y+1] && ignoreDir != "l" {
		crawl(x, y+1, d+1, "r", grid, groupCells)
	}
}

func alreadyNoted(in *string, list *gcs) bool {
	for _, v := range list.s {
		if v == *in {
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
