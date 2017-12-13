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

	lines := strings.Split(string(c), "\n")

	compact := map[int]int{}
	last := 0
	for _, l := range lines {
		s := strings.Split(l, ":")
		name, err := strconv.Atoi(s[0])
		if err != nil {
			log.Fatal(err)
		}

		size, err := strconv.Atoi(strings.TrimSpace(s[1]))
		if err != nil {
			log.Fatal(err)
		}

		compact[name] = size
		last = name
	}

	flat := []int{}
	for i := 0; i < last+1; i++ {
		if v, ok := compact[i]; ok {
			flat = append(flat, v)
			continue
		}
		flat = append(flat, 0)
	}

	attempt := 0
	for {

		if getSchwifty(flat, attempt) {
			fmt.Println(attempt)
			os.Exit(0)
		}
		attempt++
	}

}

func getSchwifty(x []int, offset int) bool {
	for i := 0; i < len(x); i++ {
		if x[i] != 0 && (i+offset)%((2*x[i])-2) == 0 {
			return false
		}
	}
	return true
}
