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

type ll struct {
	val  int
	prev *ll
	next *ll
}

func main() {
	inputs := readInput()
	problem1(inputs)
	problem2(inputs)
}

func problem1(inputs []string) {
	var nums []*ll
	var zero *ll

	var first, prev *ll
	for i, v := range inputs {
		var val int
		val, _ = strconv.Atoi(v)
		var cur ll
		if i == 0 {
			cur = ll{val, nil, nil}
			first = &cur
		} else if i == len(inputs)-1 {
			cur = ll{val, prev, first}
			prev.next = &cur
			first.prev = &cur
		} else {
			cur = ll{val, prev, nil}
			prev.next = &cur
		}
		nums = append(nums, &cur)
		if val == 0 {
			zero = &cur
		}
		prev = &cur
	}

	for _, p := range nums {
		if p.val != 0 {
			// Slice itself out of ll
			p.prev.next = p.next
			p.next.prev = p.prev

			// Find insert position
			target := p
			for s := p.val; s != 0; s -= Sign(p.val) {
				if Sign(p.val) == 1 {
					target = target.next
				} else {
					target = target.prev
				}
			}

			// Insert itself into ll
			if Sign(p.val) == 1 {
				target.next.prev = p
				p.next = target.next
				target.next = p
				p.prev = target

			} else {
				target.prev.next = p
				p.prev = target.prev
				target.prev = p
				p.next = target
			}

		}
	}

	var ans int
	o := zero
	for i := 1; i <= 3000; i++ {
		o = o.next
		if i%1000 == 0 {
			ans += o.val
		}
	}
	fmt.Printf("Task 1: %d\n", ans)
}

func problem2(inputs []string) {
	var nums []*ll
	var zero *ll

	var first, prev *ll
	for i, v := range inputs {
		var val int
		val, _ = strconv.Atoi(v)
		val *= 811589153
		var cur ll
		if i == 0 {
			cur = ll{val, nil, nil}
			first = &cur
		} else if i == len(inputs)-1 {
			cur = ll{val, prev, first}
			prev.next = &cur
			first.prev = &cur
		} else {
			cur = ll{val, prev, nil}
			prev.next = &cur
		}
		nums = append(nums, &cur)
		if val == 0 {
			zero = &cur
		}
		prev = &cur
	}

	for r := 1; r <= 10; r++ {
		for _, p := range nums {
			if p.val%(len(inputs)-1) == 0 {
				continue
			}

			// Slice element out of ll
			p.prev.next = p.next
			p.next.prev = p.prev

			// Find insert position
			target := p
			for s := p.val % (len(inputs) - 1); s != 0; s -= Sign(p.val) {
				if Sign(p.val) == 1 {
					target = target.next
				} else {
					target = target.prev
				}
			}

			// Insert element into ll
			if Sign(p.val) == 1 {
				target.next.prev = p
				p.next = target.next
				target.next = p
				p.prev = target
			} else {
				target.prev.next = p
				p.prev = target.prev
				target.prev = p
				p.next = target
			}
		}
	}

	var ans int
	o := zero
	for i := 1; i <= 3000; i++ {
		o = o.next
		if i%1000 == 0 {
			ans += o.val
		}
	}
	fmt.Printf("Task 1: %d\n", ans)
}

func Abs(in int) int {
	if in < 0 {
		return in * -1
	}
	return in
}

func Sign(in int) int {
	if in < 0 {
		return -1
	}
	return 1
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
