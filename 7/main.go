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

	// fmt.Println(string(c))

	var origo xnode
	data := strings.Split(string(c), "\n")
	// walk through list and see if they have children,
	for _, x := range data {
		orphan := true
		subject := getComponents(x)
		// check if anyone has them as child
		for _, y := range data {
			node := getComponents(y)
			if node.children == nil {
				continue
			}
			if hasChild(node.children, subject.name) {
				orphan = false
				break
			}
			// if hasChild(ynodeChildren, x)
		}
		if orphan {
			origo = subject
		}
	}

	fmt.Println(origo)

	// get the weight of children recursively - check where broken link is on the way
	// subtower weights

	// parse all programs and index them
	prgs := map[string]*xnode{}
	for _, x := range data {
		xx := getComponents(x)
		prgs[xx.name] = &xx
	}

	for _, c := range origo.children {
		fmt.Println(c, getChildTowerWeight(*prgs[c], prgs))
	}

}

func getChildTowerWeight(node xnode, prgs map[string]*xnode) int {
	countedWeight := node.weight

	if node.children != nil {
		// throw in a check to see if unbalanced
		childWeights := []int{}

		for _, c := range node.children {
			childWeight := getChildTowerWeight(*prgs[c], prgs)
			childWeights = append(childWeights, childWeight)
			// check if one child weight is broken
			countedWeight = countedWeight + childWeight
		}
		for _, cw := range childWeights {
			for _, cw2 := range childWeights {
				if cw != cw2 {
					fmt.Println(node, childWeights)
				}
			}
		}
	}

	return countedWeight
}

func hasChild(children []string, nodeName string) bool {
	for _, c := range children {
		if nodeName == c {
			return true
		}
	}
	return false
}

type xnode struct {
	name     string
	weight   int
	children []string
}

func getComponents(node string) xnode {
	comp := strings.Split(node, "->")

	subcomp := strings.Split(comp[0], "(")
	name := strings.TrimSpace(subcomp[0])

	weight, err := strconv.Atoi(strings.TrimSpace(strings.Replace(subcomp[1], ")", "", -1)))
	if err != nil {
		log.Fatal(err)
	}

	if len(comp) > 1 {
		return xnode{
			name:     name,
			weight:   weight,
			children: strings.Split(strings.TrimSpace(comp[1]), ", "),
		}
	}

	return xnode{
		name:     name,
		weight:   weight,
		children: nil,
	}
}
