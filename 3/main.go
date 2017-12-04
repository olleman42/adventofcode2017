package main

import (
	"fmt"
	"log"
	"math"
	"strconv"
)

func main() {
	//  1, 1, 2, 2, 3, 3, 4, 4, 5, 5
	// 1, 2, 1, 2, 1, 2, 1, 2, 3, 2, 3, 4, 3, 2, 3, 4, 3, 2, 3, 4, 3, 2, 3, 4, 5
	// 2 -> 2, 3 -> 2, 4->4, 5->4, 6->6, 7->6
	// generate matrix
	genMatrix(100)

}

// count h-steps and v-steps

func genMatrix(size int) {
	storage := make(map[string]int)
	storage["0,0"] = 1

	v, h := 0, 0

	dir := 0
	dirs := [4]func(v, h int) (int, int){
		func(v, h int) (int, int) { return v, h + 1 }, // RIGHT
		func(v, h int) (int, int) { return v + 1, h }, // UP
		func(v, h int) (int, int) { return v, h - 1 }, //LEFT
		func(v, h int) (int, int) { return v - 1, h }, //DOWN
	}

	turner := 2
	turnCounter := 0
	sizeCounter := 1
	counter := 0

	for i := 2; i <= size; i++ {
		if counter == sizeCounter {
			// change direction
			if dir == 3 {
				dir = 0
			} else {
				dir++
			}

			// check if tail needs to be increased
			turnCounter++
			if turnCounter == turner {
				// increase tail
				turnCounter = 0
				sizeCounter++
			}

			counter = 0
		}

		v, h = dirs[dir](v, h)

		// get neighbours
		// x x x
		// x y x
		// x x x
		partsum := 0
		if val, ok := storage[strconv.Itoa(v+1)+","+strconv.Itoa(h+1)]; ok {
			partsum = partsum + val
		}
		if val, ok := storage[strconv.Itoa(v)+","+strconv.Itoa(h+1)]; ok {
			partsum = partsum + val
		}
		if val, ok := storage[strconv.Itoa(v+1)+","+strconv.Itoa(h)]; ok {
			partsum = partsum + val
		}
		if val, ok := storage[strconv.Itoa(v-1)+","+strconv.Itoa(h+1)]; ok {
			partsum = partsum + val
		}
		if val, ok := storage[strconv.Itoa(v+1)+","+strconv.Itoa(h-1)]; ok {
			partsum = partsum + val
		}
		if val, ok := storage[strconv.Itoa(v)+","+strconv.Itoa(h-1)]; ok {
			partsum = partsum + val
		}
		if val, ok := storage[strconv.Itoa(v-1)+","+strconv.Itoa(h)]; ok {
			partsum = partsum + val
		}
		if val, ok := storage[strconv.Itoa(v-1)+","+strconv.Itoa(h-1)]; ok {
			partsum = partsum + val
		}
		storage[strconv.Itoa(v)+","+strconv.Itoa(h)] = partsum
		fmt.Println(partsum)

		counter++

		if partsum > 277678 {
			fmt.Println(partsum)
			break
		}

	}
	log.Println(v, h)
	distance := math.Abs(float64(v)) + math.Abs(float64(h))
	log.Println(distance)
	log.Println(storage)
}
