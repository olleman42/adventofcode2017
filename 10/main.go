package main

import (
	"encoding/hex"
	"fmt"
)

func main() {
	// input := "187,254,0,81,169,219,1,190,19,102,255,56,46,32,2,216"
	input := append([]byte("187,254,0,81,169,219,1,190,19,102,255,56,46,32,2,216"), []byte{17, 31, 73, 47, 23}...)
	// input := "3,4,1,5"

	// binput := bytes.NewBufferString(input)

	// for _, v := range input {
	// 	fmt.Println(byte(v))
	// }
	// fmt.Println(input)

	// li := strings.Split(input, ",")

	lengths := []int{}

	for _, length := range input {
		lengths = append(lengths, int(length))
	}

	// for _, length := range li {
	// 	p, err := strconv.Atoi(length)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	lengths = append(lengths, p)
	// }

	nums := []int{}

	for i := 0; i < 256; i++ {
		nums = append(nums, i)
	}

	workingIndex := 0
	skipSize := 0

	for z := 0; z < 64; z++ {
		for _, length := range lengths {
			// fmt.Println(workingIndex)
			// fmt.Println(length)
			// grab appropriate length, reverse
			startIndex := workingIndex
			stopIndex := workingIndex + length

			numbuf := make([]int, len(nums))
			copy(numbuf, nums)

			if stopIndex >= len(nums) {
				// fmt.Println("err double")
				stopIndex := workingIndex + length - len(nums)
				firstPart := numbuf[workingIndex:]
				secondPart := numbuf[:stopIndex]
				// fmt.Println(firstPart, secondPart)
				firstPart, secondPart = reverseDouble(firstPart, secondPart)
				// fmt.Println(firstPart, secondPart)

				nums = append(secondPart, numbuf[stopIndex:startIndex]...)
				nums = append(nums, firstPart...)
				// fmt.Println(nums)
			} else {
				// fmt.Println("err single")
				nums = append(numbuf[:startIndex], reverse(numbuf[startIndex:stopIndex])...)
				nums = append(nums, numbuf[stopIndex:]...)
			}

			// advance the workingIndex an appropriate amount and iterate skipSize
			workingIndex = workingIndex + length + skipSize
			// fmt.Println(workingIndex, length+skipSize, len(nums))
			for workingIndex >= len(nums) { // wrap around
				workingIndex = workingIndex - len(nums)
			}
			skipSize++

			// fmt.Println("----")

		}

	}
	fmt.Println(nums)

	// reduce to dense hash
	dense := []byte{}
	for i := 0; i < 16; i++ {
		collapsed := nums[i*16]
		for j := 1; j < 16; j++ {
			collapsed = collapsed ^ nums[(i*16)+j]
		}
		dense = append(dense, byte(collapsed))
	}

	fmt.Println(dense)

	fmt.Println(hex.EncodeToString(dense))
}

func reverse(stuff []int) []int {
	x := []int{}
	for i := len(stuff) - 1; i > -1; i-- {
		x = append(x, stuff[i])
	}
	return x
}

func reverseDouble(a, b []int) ([]int, []int) {
	merged := append(a, b...)

	ax, bx := []int{}, []int{}
	for i := len(merged) - 1; i > -1; i-- {
		if i >= len(b) {
			ax = append(ax, merged[i])
			continue
		}
		bx = append(bx, merged[i])
	}

	return ax, bx

}
