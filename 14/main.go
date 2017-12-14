package main

import (
	"encoding/hex"
	"fmt"
	"strconv"
)

func main() {
	grid := getGrid()

	// count all filled
	cn := 0
	for _, r := range grid {
		for _, c := range r {
			if c {
				cn++
			}

		}
	}

	fmt.Println(cn)

}

func getGrid() [][]bool {
	rows := [][]bool{}
	for i := 0; i < 128; i++ {

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
