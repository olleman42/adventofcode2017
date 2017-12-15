package main

import "fmt"

func main() {
	facA := 16807
	facB := 48271

	// initA := 65
	// initB := 8921

	initA := 783
	initB := 325

	matchCount := 0
	for i := 0; i < 5e6; i++ {
		if i%1000000 == 0 {
			fmt.Println(i)
		}

		a := (initA * facA) % 2147483647
		for a%4 != 0 {
			a = a * facA % 2147483647
		}
		b := (initB * facB) % 2147483647
		for b%8 != 0 {
			b = b * facB % 2147483647
		}

		initA = a
		initB = b

		sA := fmt.Sprintf("%016b", a)
		sB := fmt.Sprintf("%016b", b)

		lsA := sA[len(sA)-16:]
		lsB := sB[len(sB)-16:]

		if lsA == lsB {
			matchCount++
		}

	}

	fmt.Println(matchCount)
}
