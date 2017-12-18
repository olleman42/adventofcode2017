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

	in0, in1, count := make(chan int, 20), make(chan int, 20), make(chan int)
	wait := make(chan bool)

	go runProgram(0, in0, in1, count, itrs)
	go runProgram(1, in1, in0, count, itrs)

	tally0, tally1 := 0, 0
	go func(count chan int, tally0, tally1 int) {
		for c := range count {
			// fmt.Println("got count from", c)
			if c == 0 {
				tally0++
			} else {
				tally1++
			}
			// fmt.Println(tally0, tally1)
			// fmt.Println(len(in0), len(in1))
		}
	}(count, tally0, tally1)
	// os.Exit(0)
	<-wait

}

type instruction struct {
	name string
	a    interface{}
	b    interface{}
}

func runProgram(id int, in, out, count chan int, itrs []instruction) {
	reg := map[string]int{
		"i": 0,
		"a": 0,
		"p": id,
		"b": 0,
		"f": 0,
	}

	// latestSoundPlayed := 0
	// latestRecoveredFreq := 0
	currentInstruction := 0

	funcMap := map[string]func(a, b interface{}){
		"snd": func(a, b interface{}) {
			// fmt.Println(id, "sending", reg[a.(string)], a.(string), reg)
			out <- reg[a.(string)]
			count <- id
			// latestSoundPlayed = reg[a.(string)]
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
			// fmt.Println(id, "waiting for data")
			// fmt.Println(id, "received", incoming, a.(string), reg)
			reg[a.(string)] = <-in
			// if reg[a.(string)] != 0 {
			// 	fmt.Println("recovered")
			// 	latestRecoveredFreq = latestSoundPlayed
			// 	done = true
			// }
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

	i := 0
	for {

		x := itrs[currentInstruction]
		// fmt.Println(currentInstruction, x, reg)
		fmt.Println(id, "running", currentInstruction, x, reg, len(in))
		funcMap[x.name](x.a, x.b)
		currentInstruction++
		if currentInstruction == len(itrs) {
			fmt.Println("program exiting")
			break
		}
		i++

	}

}

// each should work until both are waiting for more
