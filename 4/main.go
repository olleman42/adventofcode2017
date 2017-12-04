package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sort"
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
	lines := strings.Split(string(c), "\n")

	vCount := 0

	for _, l := range lines {
		parts := strings.Fields(l)
		found := false
		letterbuckets := [][2]string{}
		for _, p := range parts {
			if found {
				break
			}
			// check if normalized version exsists in letterbucket
			normalized := normalizeString(p)
			for _, x := range letterbuckets {
				if normalized == x[0] {
					found = true
					// fmt.Println(l)
					fmt.Println(normalized + " - " + p)
					fmt.Println(l)
					fmt.Println(letterbuckets)
					vCount++
					break
				}
			}

			letterbuckets = append(letterbuckets, [2]string{normalized, p})

		}
		// fmt.Println(letterbuckets)

	}

	fmt.Println(len(lines) - vCount)
}

type Hmm []rune

func (s Hmm) Len() int           { return len(s) }
func (s Hmm) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s Hmm) Less(i, j int) bool { return s[i] < s[j] }

func normalizeString(in string) string {
	r := []rune(in)
	sort.Sort(Hmm(r))
	return string(r)
}
