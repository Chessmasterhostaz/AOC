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

var total int

func problem1(inputs []string) {
	for _, v := range inputs {
	loopo:
		for _, l1 := range v[:len(v)/2] {
			for _, l2 := range v[len(v)/2:] {
				if l1 == l2 {
					var val int
					if val = int(l1) % 'a'; val > 26 {
						val -= ('A' - 26)
					}
					total = total + 1 + val
					break loopo
				}
			}
		}
	}
	fmt.Println(total)
}

var tmpLst = make([]int, 52)

func problem2(inputs []string) {
	total = 0
	group := 0
	for _, v := range inputs {
		for _, l1 := range v {
			var val int
			if val = int(l1) % 'a'; val > 26 {
				val -= ('A' - 26)
			}
			tmpLst[val] |= (1 << group)
		}
		if group = (group + 1) % 3; group == 0 {
			for i := range tmpLst {
				if tmpLst[i] == 7 {
					total += i + 1
				}
				tmpLst[i] = 0
			}
		}
	}
	fmt.Println(total)
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
