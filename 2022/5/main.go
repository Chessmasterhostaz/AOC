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
	stacklen := (len(inputs[0]) + 1) / 4
	stacks := make([][]string, stacklen)

	var inputStart int
	for i, v := range inputs {
		if v[1] == '1' {
			inputStart = i + 2
			break
		}

		for i := 0; i < stacklen; i++ {
			c := v[(1 + 4*i):(2 + 4*i)]
			if c != " " {
				stacks[i] = append([]string{c}, stacks[i]...)
			}
		}
	}

	for i, v := range inputs {
		if i < inputStart {
			continue
		}
		values := strings.Split(v, " ")
		num, _ := strconv.Atoi(values[1])
		src, _ := strconv.Atoi(values[3])
		src--
		tar, _ := strconv.Atoi(values[5])
		tar--

		for j := 0; j < num; j++ {
			n := len(stacks[src]) - 1
			stacks[tar] = append(stacks[tar], stacks[src][n])
			stacks[src] = stacks[src][:n]
		}

	}

	fmt.Printf("Task 1: ")
	for _, s := range stacks {
		fmt.Printf(s[len(s)-1])
	}
	fmt.Println()
}

func problem2(inputs []string) {
	stacklen := (len(inputs[0]) + 1) / 4
	stacks := make([][]string, stacklen)

	var inputStart int
	for i, v := range inputs {
		if v[1] == '1' {
			inputStart = i + 2
			break
		}

		for i := 0; i < stacklen; i++ {
			c := v[(1 + 4*i):(2 + 4*i)]
			if c != " " {
				stacks[i] = append([]string{c}, stacks[i]...)
			}
		}
	}

	for i, v := range inputs {
		if i < inputStart {
			continue
		}
		values := strings.Split(v, " ")
		num, _ := strconv.Atoi(values[1])
		src, _ := strconv.Atoi(values[3])
		src--
		tar, _ := strconv.Atoi(values[5])
		tar--

		n := len(stacks[src]) - num
		stacks[tar] = append(stacks[tar], stacks[src][n:]...)
		stacks[src] = stacks[src][:n]
	}

	fmt.Printf("Task 2: ")
	for _, s := range stacks {
		fmt.Printf(s[len(s)-1])
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
