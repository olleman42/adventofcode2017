package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}

	c, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(c), "\r\n")

	fixed := []int{}

	// walk through each line and find the biggest difference
	for _, l := range lines {
		row := strings.Split(l, "\t")
		ints := []int{}
		for _, c := range row {
			conv, err := strconv.Atoi(c)
			if err != nil {
				log.Fatal(err)
			}
			ints = append(ints, conv)
		}

		bg := 0

		for _, c := range ints {

			for _, c2 := range ints {
				if c-c2 > bg {
					bg = c - c2
				}
			}
		}
		fixed = append(fixed, bg)
	}

	fmt.Println(fixed)

	sum := 0
	for _, c := range fixed {
		sum = sum + c
	}
	fmt.Println(sum)

}
