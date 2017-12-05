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
	stepCount := 0

	f, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}

	c, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}

	instrx := strings.Split(string(c), "\n")
	instr := []int{}
	for _, v := range instrx {
		p, err := strconv.Atoi(v)
		if err != nil {
			log.Fatal(err)
		}
		instr = append(instr, p)
	}

	fmt.Println(instr)

	instrLoc := 0
	for {
		if instrLoc < len(instr) {
			instrValue := instr[instrLoc]
			instr[instrLoc] = instrValue + 1
			instrLoc = instrLoc + instrValue
			fmt.Println(instrLoc)
			stepCount++
			continue
		}
		break
	}

	fmt.Println(stepCount)
}
