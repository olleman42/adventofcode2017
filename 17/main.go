package main

import (
	"fmt"
)

func main() {
	cb := []int{0}

	input := 382
	// input := 3
	currentIndex := 0

	for i := 1; i <= 50e6; i++ {

		// step forward input steps
		// for x := 0; x < input; x++ {
		// 	currentIndex++
		// 	if currentIndex > len(cb)-1 {
		// 		currentIndex = 0
		// 	}
		// }
		currentIndex = (currentIndex + input) % len(cb)
		// fmt.Println(currentIndex)

		//insert value in current index+1 (splice into place)
		cb = append(cb[:currentIndex+1], append([]int{i}, cb[currentIndex+1:]...)...)

		// set starting index input value
		currentIndex++
		if i%100000 == 0 {
			fmt.Println(i)
		}
		// increase input value
	}
	fmt.Println(currentIndex)
	fmt.Println(cb[currentIndex])
	fmt.Println(">>", cb[currentIndex+1])
	fmt.Println(cb[1])

}
