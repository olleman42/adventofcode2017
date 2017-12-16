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

	steps := strings.Split(string(c), ",")

	subject := []rune{}
	for _, r := range "abcdefghijklmnop" {
		subject = append(subject, r)
	}

	crunch := func(su []rune) []rune {
		sx := make([]rune, 16)
		copy(sx, su)
		for _, step := range steps {
			// parse step
			switch step[0] {
			case 's':
				spinPositions, _ := strconv.Atoi(step[1:])
				sx = append(sx[len(sx)-spinPositions:], sx[:len(sx)-spinPositions]...)

			case 'x':
				//swap positions
				posData := step[1:]
				pd := strings.Split(posData, "/")
				p1, _ := strconv.Atoi(pd[0])
				p2, _ := strconv.Atoi(pd[1])
				sx[p1], sx[p2] = sx[p2], sx[p1]
			case 'p':
				posData := step[1:]
				pd := strings.Split(posData, "/")
				// find index for each position
				var p1, p2 int
				for i, r := range sx {
					if string(r) == pd[0] {
						p1 = i
					}
					if string(r) == pd[1] {
						p2 = i
					}
				}
				sx[p1], sx[p2] = sx[p2], sx[p1]

			}
		}
		return sx
	}

	order := []string{"abcdefghijklmnop"}
	permCache := map[string][]rune{}
	for x := 0; x < 1e9; x++ {
		if x%1000000 == 0 {
			fmt.Println(x / 1e6)
		}

		// if x%1e2 == 0 {
		// 	fmt.Println(x)
		// }
		in := string(subject)

		if out, ok := permCache[in]; ok {
			subject = out
			break
			continue
		}
		subject = crunch(subject)
		// fmt.Println(in, string(subject))
		permCache[in] = subject
		order = append(order, string(subject))

	}
	order = order[:60]
	fmt.Println("order len", len(order))
	for _, x := range order {
		fmt.Println(x)
	}

	fx := 0
	for i := 0; i <= 1e9; i++ {
		fx++
		if fx >= 60 {
			fx = 0
		}
	}
	fmt.Println(fx - 1)
	fmt.Println(order[fx-1])
	// fmt.Println(string(subject))
	// fmt.Println(len(permCache))
	// for i, x := range order {
	// 	if x == string(subject) {
	// 		fmt.Println(i)
	// 		fmt.Println(1e9 % i)

	// 	}
	// }
	// check that they all link back

	// for k, v := range permCache {
	// 	if _, ok := permCache[string(v)]; ok {
	// 		fmt.Println(k, string(v))
	// 		continue
	// 	}
	// 	fmt.Println("Next step not found", string(v))
	// 	fmt.Println(len(v))
	// }
}
