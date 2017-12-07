package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
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
			fmt.Println(subject)
		}
	}

	// if noone - you've found em

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
	children []string
}

func getComponents(node string) xnode {
	comp := strings.Split(node, "->")

	name := strings.TrimSpace(strings.Split(comp[0], "(")[0])

	if len(comp) > 1 {
		return xnode{
			name:     name,
			children: strings.Split(strings.TrimSpace(comp[1]), ", "),
		}
	}

	return xnode{
		name:     name,
		children: nil,
	}
}
