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
	problem1(inputs)
	problem2(inputs)
}

func problem1(inputs []string) {
	var ans int
	x := 1
	ticks := 0
	cycle := 0
	val := 0
	for _, v := range inputs {
		if v[:4] == "noop" {
			ticks = 1
			val = 0
		} else {
			ticks = 2
			val, _ = strconv.Atoi(v[5:])
		}

		for t := 0; t < ticks; t++ {
			cycle++
			if ((cycle - 20) % 40) == 0 {
				ans += x * cycle
			}
		}
		x += val
	}
	fmt.Printf("Task 1: %d\n", ans)
}

func problem2(inputs []string) {
	x := 1
	ticks := 0
	cycle := 0
	val := 0
	println("Task 2:")
	for _, v := range inputs {
		if v[:4] == "noop" {
			ticks = 1
			val = 0
		} else {
			ticks = 2
			val, _ = strconv.Atoi(v[5:])
		}

		for t := 0; t < ticks; t++ {
			pos := cycle % 40
			if cycle > 1 && pos == 0 {
				println()
			}

			if pos-x < 2 && pos-x > -2 {
				print(" #")
			} else {
				print(" .")
			}
			cycle++
		}
		x += val
	}
	println()
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
