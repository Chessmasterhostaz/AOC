package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

func main() {
	inputs := readInput()
	part1(inputs)
	part2(inputs)
}

func part1(inputs []string) {
	total := 0
	for _, v := range inputs {
		s := strings.Split(v, " ")
		opp := int(s[0][0] - 64)
		me := int(s[1][0] - 87)
		switch {
		case (me-opp) == 1 || (me-opp) == -2:
			total = total + me + 6
		case me == opp:
			total = total + me + 3
		}
	}
	fmt.Printf("Part 1: %d\n", total)
}

func part2(inputs []string) {
	total := 0
	for _, v := range inputs {
		s := strings.Split(v, " ")
		opp := int(s[0][0] - 64)
		me := int(s[1][0] - 87)
		switch me {
		case 1: //lose
			if opp == 1 {
				total = total + 3
			} else {
				total = total + opp - 1
			}
		case 2: //draw
			total = total + 3 + opp
		case 3: //win
			if opp == 3 {
				total = total + 6 + 1
			} else {
				total = total + opp + 6 + 1
			}
		}
	}
	fmt.Printf("Part 2: %d\n", total)
}

func readInput() []string {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		log.Fatal("unable to get the current filename")
	}
	dirname := filepath.Dir(filename)

	file, err := os.Open(dirname + "/input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var inputs []string

	scan := bufio.NewScanner(file)
	for scan.Scan() {
		str := scan.Text()
		inputs = append(inputs, str)
	}

	return inputs
}
