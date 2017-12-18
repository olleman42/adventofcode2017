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
	reg := map[string]int{
		"i": 0,
		"a": 0,
		"p": 0,
		"b": 0,
		"f": 0,
	}

	latestSoundPlayed := 0
	latestRecoveredFreq := 0
	currentInstruction := 0
	done := false

	funcMap := map[string]func(a, b interface{}){
		"snd": func(a, b interface{}) {
			fmt.Println("playing sound", reg[a.(string)])
			latestSoundPlayed = reg[a.(string)]
		},
		"set": func(a, b interface{}) {
			switch b.(type) {
			case int:
				reg[a.(string)] = b.(int)
			case string:
				reg[a.(string)] = reg[b.(string)]
			}
		},
		"add": func(a, b interface{}) {
			switch b.(type) {
			case int:
				reg[a.(string)] = reg[a.(string)] + b.(int)
			case string:
				reg[a.(string)] = reg[a.(string)] + reg[b.(string)]
			}
		},
		"mul": func(a, b interface{}) {
			reg[a.(string)] = reg[a.(string)] * b.(int)
		},
		"mod": func(a, b interface{}) {
			switch b.(type) {
			case int:
				reg[a.(string)] = reg[a.(string)] % b.(int)
			case string:
				reg[a.(string)] = reg[a.(string)] % reg[b.(string)]
			}
		},
		"rcv": func(a, b interface{}) {
			fmt.Println("attempting to recover freq")
			if reg[a.(string)] != 0 {
				fmt.Println("recovered")
				latestRecoveredFreq = latestSoundPlayed
				done = true
			}
		},
		"jgz": func(a, b interface{}) {
			if reg[a.(string)] > 0 {
				currentInstruction = currentInstruction - 1
				switch b.(type) {
				case int:
					currentInstruction = currentInstruction + b.(int)
				case string:
					currentInstruction = currentInstruction + reg[b.(string)]
				}
			}
		},
	}

	f, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}

	c, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}

	// set up instructions
	itrs := []instruction{}
	for _, instr := range strings.Split(string(c), "\n") {
		pi := strings.Split(instr, " ")
		switch pi[0] {
		case "rcv":
			fallthrough
		case "snd":
			itrs = append(itrs, instruction{name: pi[0], a: pi[1]})
		case "set":
			fallthrough
		case "mul":
			fallthrough
		case "jgz":
			fallthrough
		case "mod":
			fallthrough
		case "add":
			pb, err := strconv.Atoi(pi[2])
			if err != nil {
				itrs = append(itrs, instruction{name: pi[0], a: pi[1], b: pi[2]})
				continue
			}
			itrs = append(itrs, instruction{name: pi[0], a: pi[1], b: pb})
		}

	}

	for _, i := range itrs {
		fmt.Println(i)

	}
	// os.Exit(0)

	i := 0
	for !done {

		x := itrs[currentInstruction]
		fmt.Println(currentInstruction, x, reg)
		funcMap[x.name](x.a, x.b)
		currentInstruction++
		if currentInstruction > len(itrs) {
			break
		}
		i++

	}

	fmt.Println(latestRecoveredFreq)
	fmt.Println(latestSoundPlayed)
	fmt.Println(i)

}

type instruction struct {
	name string
	a    interface{}
	b    interface{}
}
