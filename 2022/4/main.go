package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
)

func main() {
	inputs := readInput()
	problem1(inputs)
	problem2(inputs)
}

func problem1(inputs []string) {
	var overlaps int
	for _, v := range inputs {
		pair := strings.Split(v, ",")

		a1 := strings.Split(pair[0], "-")
		a2 := strings.Split(pair[1], "-")

		l1, _ := strconv.Atoi(a1[0])
		l2, _ := strconv.Atoi(a1[1])
		r1, _ := strconv.Atoi(a2[0])
		r2, _ := strconv.Atoi(a2[1])

		if l1 >= r1 {
			if l2 <= r2 {
				overlaps++
				continue
			}
		}

		if r1 >= l1 {
			if r2 <= l2 {
				overlaps++
				continue
			}
		}
	}
	fmt.Printf("Task 1: %d\n", overlaps)
}

func problem2(inputs []string) {
	var overlaps int
	for _, v := range inputs {
		pair := strings.Split(v, ",")

		a1 := strings.Split(pair[0], "-")
		a2 := strings.Split(pair[1], "-")

		l1, _ := strconv.Atoi(a1[0])
		l2, _ := strconv.Atoi(a1[1])
		r1, _ := strconv.Atoi(a2[0])
		r2, _ := strconv.Atoi(a2[1])

		if l1 >= r1 {
			if l1 <= r2 {
				overlaps++
				continue
			}
		}

		if r1 >= l1 {
			if r1 <= l2 {
				overlaps++
				continue
			}
		}
	}
	fmt.Printf("Task 2: %d\n", overlaps)
}

func printDots(start, stop int) {

	for i := 0; i < 99; i++ {
		if i < start || i > stop {
			fmt.Printf(".")
		} else {
			fmt.Printf("#")
		}
	}
	fmt.Println()
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
