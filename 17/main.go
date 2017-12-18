package main

import (
	"fmt"
	"strconv"
)

func main() {
	// cb := []int{0}
	// cblen := 1

	input := 382
	// input := 3
	currentIndex := 0
	latestNeighbor := 0

	for i := 1; i <= 50e6; i++ {

		// step forward input steps
		// for x := 0; x < input; x++ {
		// 	currentIndex++
		// 	if currentIndex > len(cb)-1 {
		// 		currentIndex = 0
		// 	}
		// }
		currentIndex = (currentIndex + input) % i
		// fmt.Println(currentIndex)

		//insert value in current index+1 (splice into place)
		// cb = append(cb[:currentIndex+1], append([]int{i}, cb[currentIndex+1:]...)...)
		// fmt.Println(currentIndex)

		if currentIndex == 0 {
			latestNeighbor = i
		}
		// set starting index input value
		currentIndex++
		if i%100000 == 0 {
			fmt.Println(i)
		}
		// increase input value
	}
	fmt.Println(currentIndex)
	// fmt.Println(cb[currentIndex])
	// fmt.Println(">>", cb[currentIndex+1])
	// fmt.Println(cb[1])
	fmt.Println("next to 0", strconv.Itoa(latestNeighbor))

}
