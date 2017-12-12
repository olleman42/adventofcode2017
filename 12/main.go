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
	f, err := os.Open("input")
	if err != nil {
		log.Fatal(f)
	}

	c, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}

	programs := strings.Split(string(c), "\n")
	xnews := false
	spentPrograms := []int{}
	group := 0

	for !xnews {
		xnews = true
		validrelations := []int{}

		nonews := false
		for !nonews {
			nonews = true

			for _, p := range programs {
				program, friends := getComponents(p)

				if len(validrelations) == 0 && !contains(validrelations, spentPrograms, program) {

					validrelations = append(validrelations, program)
					validrelations = append(validrelations, friends...)
					xnews = false
					continue
				}

				for _, relation := range validrelations {
					if relation == program {
						// validrelations = append(validrelations, friends...) // no unique lol
						for _, friend := range friends {
							if !contains(validrelations, spentPrograms, friend) {
								validrelations = append(validrelations, friend)
								nonews = false
							}
						}
					}
				}
			}
		}
		// fmt.Println(validrelations)
		// fmt.Println(len(validrelations))
		spentPrograms = append(spentPrograms, validrelations...)
		group++
	}
	fmt.Println(group - 1)
}

func contains(list, l2 []int, val int) bool {
	for _, v := range list {
		if v == val {
			return true
		}
	}
	for _, v := range l2 {
		if v == val {
			return true
		}
	}
	return false
}

func getComponents(p string) (int, []int) {
	comps := strings.Split(p, "<->")
	comps[0] = strings.TrimSpace(comps[0])
	progName, err := strconv.Atoi(comps[0])
	if err != nil {
		log.Fatal(err)
	}

	fs := strings.Split(comps[1], ",")
	friends := []int{}
	for _, f := range fs {
		p, err := strconv.Atoi(strings.TrimSpace(f))
		if err != nil {
			log.Fatal(err)
		}
		friends = append(friends, p)
	}

	return progName, friends
}
