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

	cc := strings.Replace(string(c), "\n", "", -1)
	ps := strings.Split(cc, "\t")

	parts := []int{}

	for _, v := range ps {
		p, err := strconv.Atoi(v)
		if err != nil {
			log.Fatal(err)
		}
		parts = append(parts, p)
	}

	// first, do pass to find max
	// then do pass to smear out contents
	seenCache := []string{}
	count := 0
	seenAt := 0
	done := false
	for {
		count++
		max := -1
		pos := -1
		for i, v := range parts {
			if v > max {
				max = v
				pos = i
			}
		}

		// do spread
		parts[pos] = 0
		for i := max; i > 0; i-- {
			pos++
			if pos > len(parts)-1 {
				pos = 0
			}
			parts[pos]++
		}
		fmt.Println(parts)

		// check if this exsists in cache
		// break
		t := []string{}
		for _, v := range parts {
			t = append(t, string(v))
		}
		cached := strings.Join(t, "-")

		for i, v := range seenCache {
			if v == cached {
				// done, break
				done = true
				fmt.Println(i)
				seenAt = i
			}
		}

		if done {
			break
		}
		// store in cache
		seenCache = append(seenCache, cached)

	}

	fmt.Println(count - seenAt - 1)

}
