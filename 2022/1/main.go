package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
)

func main() {
	inputs := readInput()
	problem(inputs)
}

func Max(in *[]int) int {
	m := 0
	for _, e := range *in {
		if e > m {
			m = e
		}
	}
	return m
}

func Max3(in *[]int) int {
	m2 := 0
	m1 := 0
	m := 0
	for _, e := range *in {
		if e > m {
			m2 = m1
			m1 = m
			m = e
		}
	}
	return m + m1 + m2
}

func problem(inputs []string) {
	elves := []int{0}

	var idx int
	for _, v := range inputs {
		if v != "" {
			val, err := strconv.Atoi(v)
			if err != nil {
				log.Fatal(err)
			}
			elves[idx] = elves[idx] + val
		} else {
			idx = idx + 1
			elves = append(elves, 0)
		}
	}

	fmt.Printf("Task 1: %d\n", Max(&elves))
	fmt.Printf("Task 2: %d\n", Max3(&elves))
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
