package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type vector struct {
	x, y, z int
}

func (v *vector) equals(v2 vector) bool {
	return v.x == v2.x && v.y == v2.y && v.z == v2.z
}

type particle struct {
	p, v, a vector
	dead    bool
}

func (p *particle) getDistance() int {
	return int(math.Abs(float64(p.p.x)) + math.Abs(float64(p.p.y)) + math.Abs(float64(p.p.z)))
}

func main() {
	f, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	c, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}

	ps := strings.Split(string(c), "\n")

	particles := []*particle{}

	for _, p := range ps {
		comp := regexp.MustCompile(`([0-9])\w+|-([0-9]+)|[0-9]`)
		x := comp.FindAllString(p, -1)
		fmt.Println(x)
		prsed := []int{}
		for _, q := range x {
			pp, err := strconv.Atoi(q)
			if err != nil {
				log.Fatal(err)
			}
			prsed = append(prsed, pp)
		}
		// fmt.Println(comp.FindAllString(p, -1))
		particles = append(particles, &particle{
			p: vector{prsed[0], prsed[1], prsed[2]},
			v: vector{prsed[3], prsed[4], prsed[5]},
			a: vector{prsed[6], prsed[7], prsed[8]},
		})

	}

	// walk through many many times and check who is the closest to 0,0,0, do it
	for i := 0; i < 500; i++ {
		closest := [2]int{0, 100000000000}
		for pname, p := range particles {
			if p.dead {
				continue
			}
			p.v = vector{p.v.x + p.a.x, p.v.y + p.a.y, p.v.z + p.a.z}
			p.p = vector{p.v.x + p.p.x, p.v.y + p.p.y, p.v.z + p.p.z}
			if p.getDistance() < closest[1] {
				closest = [2]int{pname, p.getDistance()}
			}
		}
		fmt.Println(particles[0])
		fmt.Println(closest)
		// check for collisions
		collisions := []vector{}
		for pname, p := range particles {
			// check is position collides with anyone else
			if p.dead {
				continue
			}
			for ppname, pp := range particles {
				if pp.dead {
					continue
				}
				if pname != ppname && p.p.equals(pp.p) {
					collisions = append(collisions, p.p)
					fmt.Println("collision at ", p)
				}
			}
		}
		// if yes, remove everyone with this position
		for _, cv := range collisions {
			for _, p := range particles {
				if cv.equals(p.p) {
					p.dead = true
				}
			}
		}

	}

	alive := 0
	for _, p := range particles {
		if p.dead {
			continue
		}
		alive++
	}
	fmt.Println(alive)
}
