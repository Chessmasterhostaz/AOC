package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

func main() {
	inputs := readInput()
	problem1(inputs)
	problem2(inputs)
}

func problem1(inputs []string) {
	var c1, c2, c3, c4 string
	var ans int

	for i, c := range inputs[0] {
		c1 = c2
		c2 = c3
		c3 = c4
		c4 = string(c)

		if i > 3 && c1 != c2 && c1 != c3 && c1 != c4 && c2 != c3 && c2 != c4 && c3 != c4 {
			ans = i + 1
			break
		}
	}

	fmt.Printf("Task 1: %d\n", ans)
}

func problem2(inputs []string) {
	var nums [14]string
	var ans int
	for i, c := range inputs[0] {
		for j := 0; j < 13; j++ {
			nums[j] = nums[j+1]
		}
		nums[13] = string(c)

		test := make(map[string]any)
		for j := 0; j < 14; j++ {
			test[fmt.Sprintf(nums[j])] = "y"
		}
		if len(test) == 14 {
			ans = i + 1
			break
		}
	}

	fmt.Printf("Task 2: %d\n", ans)
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
