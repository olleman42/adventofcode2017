package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func main() {
	input := "187,254,0,81,169,219,1,190,19,102,255,56,46,32,2,216"
	// input := "3,4,1,5"

	il := strings.Split(input, ",")

	lengths := []int{}

	for _, length := range il {
		plength, err := strconv.Atoi(length)
		if err != nil {
			log.Fatal(err)
		}
		lengths = append(lengths, plength)
	}

	nums := []string{}

	for i := 0; i < 256; i++ {
		nums = append(nums, strconv.Itoa(i))
	}

	workingIndex := 0
	skipSize := 0

	for _, length := range lengths {
		fmt.Println(workingIndex)
		fmt.Println(length)
		// grab appropriate length, reverse
		startIndex := workingIndex
		stopIndex := workingIndex + length

		numbuf := make([]string, len(nums))
		copy(numbuf, nums)

		if stopIndex >= len(nums) {
			fmt.Println("err double")
			stopIndex := workingIndex + length - len(nums)
			firstPart := numbuf[workingIndex:]
			secondPart := numbuf[:stopIndex]
			fmt.Println(firstPart, secondPart)
			firstPart, secondPart = reverseDouble(firstPart, secondPart)
			fmt.Println(firstPart, secondPart)

			nums = append(secondPart, numbuf[stopIndex:startIndex]...)
			nums = append(nums, firstPart...)
			fmt.Println(nums)
		} else {
			fmt.Println("err single")
			nums = append(numbuf[:startIndex], reverse(numbuf[startIndex:stopIndex])...)
			nums = append(nums, numbuf[stopIndex:]...)
		}
		fmt.Println(nums)

		// advance the workingIndex an appropriate amount and iterate skipSize
		workingIndex = workingIndex + length + skipSize
		fmt.Println(workingIndex, length+skipSize, len(nums))
		if workingIndex >= len(nums) { // wrap around
			workingIndex = workingIndex - len(nums)
		}
		skipSize++

		fmt.Println("----")

	}
}

func reverse(stuff []string) []string {
	x := []string{}
	for i := len(stuff) - 1; i > -1; i-- {
		x = append(x, stuff[i])
	}
	return x
}

func reverseDouble(a, b []string) ([]string, []string) {
	merged := append(a, b...)

	ax, bx := []string{}, []string{}
	for i := len(merged) - 1; i > -1; i-- {
		if i >= len(b) {
			ax = append(ax, merged[i])
			continue
		}
		bx = append(bx, merged[i])
	}

	return ax, bx

}
