package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	f, err := os.Open("input")

	if err != nil {
		log.Fatal(err)
	}

	c, err := ioutil.ReadAll(f)

	if err != nil {
		log.Fatal(err)
	}

	groupStack := 0
	totalScore := 0
	skipNextStep := false
	currentlyInGarbage := false
	for _, step := range string(c) {
		if skipNextStep {
			skipNextStep = false
			continue
		}
		if step == '!' {
			skipNextStep = true
			continue
		}
		if !currentlyInGarbage {
			if step == '<' {
				currentlyInGarbage = true
				continue
			}
			if step == '{' {
				groupStack++
				continue
			}
			if step == '}' {
				totalScore = totalScore + groupStack
				groupStack--
				continue
			}
		}
		if currentlyInGarbage {
			if step == '>' {
				currentlyInGarbage = false
			}
		}

	}

	fmt.Println(totalScore)
}
