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

	instructions := strings.Split(string(c), "\r\n")

	membank := map[string]int{}
	maxValue := 0

	for _, instr := range instructions {
		fmt.Println(instr)
		instrcomps := strings.Split(instr, " ")

		memposition := instrcomps[0]
		storedValue := 0
		var ok bool
		if storedValue, ok = membank[memposition]; !ok {
			membank[memposition] = 0
		}
		instrToPerform := instrcomps[1]
		instrValue, err := strconv.Atoi(instrcomps[2])
		if err != nil {
			log.Fatal(err)
		}

		conditional := instrcomps[4:]

		if checkConditional(conditional, membank) {
			switch instrToPerform {
			case "inc":
				membank[memposition] = storedValue + instrValue
			case "dec":
				membank[memposition] = storedValue - instrValue

			}
			if membank[memposition] > maxValue {
				maxValue = membank[memposition]
			}
		}

	}

	fmt.Println(membank)
	fmt.Println(maxValue)

}

var condOpMap = map[string]func(int, string, int) bool{
	"==": func(operand int, op string, cond int) bool { return operand == cond },
	"!=": func(operand int, op string, cond int) bool { return operand != cond },
	">=": func(operand int, op string, cond int) bool { return operand >= cond },
	"<=": func(operand int, op string, cond int) bool { return operand <= cond },
	"<":  func(operand int, op string, cond int) bool { return operand < cond },
	">":  func(operand int, op string, cond int) bool { return operand > cond },
}

func checkConditional(conditional []string, membank map[string]int) bool {
	fmt.Println(conditional)
	memPosCheck := conditional[0]
	storedValue := 0
	ok := false
	if storedValue, ok = membank[memPosCheck]; !ok {
		membank[memPosCheck] = 0
	}
	op := conditional[1]
	comp := conditional[2]
	pcomp, err := strconv.Atoi(comp)
	if err != nil {
		log.Fatal(err)
	}
	return condOpMap[op](storedValue, op, pcomp)
}
