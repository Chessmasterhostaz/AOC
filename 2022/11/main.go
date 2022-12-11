package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"sort"
	"strconv"
	"strings"
)

type Monkey struct {
	items []int64
	op    func(*Monkey, int64) int64
	val   int
	test  int
	s     int
	f     int
	ins   int
}

func (m *Monkey) Multiply(in int64) int64 {
	return in * int64(m.val)
}

func (m *Monkey) Square(in int64) int64 {
	return in * in
}

func (m *Monkey) Add(in int64) int64 {
	return in + int64(m.val)
}

func main() {
	inputs := readInput()
	problem1(inputs)
	problem2(inputs)
}

func problem1(inputs []string) {
	monkeys := makeMonkeys(inputs)

	ins := []int{}
	for i := 0; i <= len(inputs)/7; i++ {
		ins = append(ins, 0)
	}

	for r := 0; r < 20; r++ {
		for mi, m := range monkeys {
			for _, i := range m.items {
				val := m.op(&m, i) / 3
				if val%int64(m.test) != 0 {
					monkeys[m.f].items = append(monkeys[m.f].items, val)
				} else {
					monkeys[m.s].items = append(monkeys[m.s].items, val)
				}
			}
			ins[mi] += len(monkeys[mi].items)
			monkeys[mi].items = monkeys[mi].items[:0]
		}
	}

	sort.Ints(ins)
	var ans = ins[len(inputs)/7] * ins[len(inputs)/7-1]
	fmt.Printf("Task 1: %d\n", ans)
}

func problem2(inputs []string) {
	monkeys := makeMonkeys(inputs)

	ins := []int{}
	for i := 0; i <= len(inputs)/7; i++ {
		ins = append(ins, 0)
	}

	for r := 0; r < 10000; r++ {
		for mi, m := range monkeys {
			for _, i := range m.items {
				val := m.op(&m, i) // 3
				if val%int64(m.test) != 0 {
					monkeys[m.f].items = append(monkeys[m.f].items, val%9699690)
				} else {
					monkeys[m.s].items = append(monkeys[m.s].items, val%9699690)
				}
			}
			ins[mi] += len(monkeys[mi].items)
			monkeys[mi].items = monkeys[mi].items[:0]
		}
	}

	sort.Ints(ins)
	var ans = ins[len(inputs)/7] * ins[len(inputs)/7-1]
	fmt.Printf("Task 2: %d\n", ans)
}

func makeMonkeys(inputs []string) []Monkey {
	var monkeys []Monkey
	for i := 0; i*7 < len(inputs); i++ {
		stringStarters := strings.Split(inputs[i*7+1][18:], ", ")
		var intStarters []int64
		for _, x := range stringStarters {
			val, _ := strconv.ParseInt(x, 10, 64)
			intStarters = append(intStarters, val)
		}

		operation := inputs[i*7+2][23:24]
		var opFunc func(*Monkey, int64) int64
		opVal := 0
		switch operation {
		case "*":
			if inputs[i*7+2][25:26] == "o" {
				opFunc = (*Monkey).Square
			} else {
				opFunc = (*Monkey).Multiply
				opVal, _ = strconv.Atoi(inputs[i*7+2][25:])
			}
		case "+":
			opFunc = (*Monkey).Add
			opVal, _ = strconv.Atoi(inputs[i*7+2][25:])
		}

		testValue, _ := strconv.Atoi(inputs[i*7+3][21:])
		success, _ := strconv.Atoi(inputs[i*7+4][29:])
		failure, _ := strconv.Atoi(inputs[i*7+5][30:])

		monkeys = append(monkeys, Monkey{
			items: intStarters,
			op:    opFunc,
			val:   opVal,
			test:  testValue,
			s:     success,
			f:     failure,
			ins:   0,
		})
	}
	return monkeys
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
