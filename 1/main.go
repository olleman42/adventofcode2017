package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Need to specify input")
	}
	input := os.Args[1]
	dubdub := input + input
	xlen := int(len(input) / 2)

	valid := []int{}

	for i, l := range input {

		if string(l) == string(dubdub[i+xlen]) {
			p, err := strconv.Atoi(string(l))
			if err != nil {
				log.Fatal(err)
			}
			valid = append(valid, p)
		}

	}

	var tot int

	for _, v := range valid {
		tot = tot + v
	}

	fmt.Println(valid)
	fmt.Println(tot)
}
