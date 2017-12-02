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

	valid := []int{}

	for i, l := range input {
		if i < len(input)-1 {
			if string(l) == string(input[i+1]) {
				p, err := strconv.Atoi(string(l))
				if err != nil {
					log.Fatal(err)
				}
				valid = append(valid, p)
			}
			continue
		}
		if string(l) == string(input[0]) {
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
