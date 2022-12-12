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
	findStartAndEnd(inputs)
	problem1(inputs)
	problem2(inputs)
}

var pts = []Point{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

type Node struct {
	p      Point
	parent *Node
}

type Point struct {
	x, y int
}

func (p *Point) Add(p2 Point) Point {
	p.x = p.x + p2.x
	p.y = p.y + p2.y
	return *p
}

var start Point
var end Point

func findStartAndEnd(inputs []string) {
	for i, v := range inputs {
		idx := strings.Index(v, "S")
		if idx >= 0 {
			start = Point{idx, i}
			inputs[i] = strings.Replace(inputs[i], "S", "a", 1)
		}
		idx = strings.Index(v, "E")
		if idx >= 0 {
			end = Point{idx, i}
			inputs[i] = strings.Replace(inputs[i], "E", "z", 1)
		}
	}
}

func problem1(inputs []string) {
	var queue []Node
	var added = make(map[Point]bool)
	var node Node

	s := Node{start, nil}
	added[start] = true
	queue = append(queue, s)

foundEnd:
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]

		for _, p := range pts {
			n := Node{p.Add(cur.p), &cur}
			if n.p.x < 0 || n.p.x >= len(inputs[0]) || n.p.y < 0 || n.p.y >= len(inputs) {
				continue
			}
			if _, ok := added[n.p]; !ok {
				if (int(inputs[n.p.y][n.p.x]) - int(inputs[cur.p.y][cur.p.x])) <= 1 {
					added[n.p] = true
					queue = append(queue, n, cur)
					if n.p == end {
						node = n
						break foundEnd
					}
				}
			}
		}
	}

	var ans int
	for {
		if node.parent == nil {
			break
		}
		node = *node.parent
		ans++
	}

	fmt.Printf("Task 1: %d\n", ans)
}

func problem2(inputs []string) {
	var queue []Node
	var added = make(map[Point]bool)
	var node Node

	s := Node{end, nil}
	added[end] = true
	queue = append(queue, s)

foundEnd:
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]

		for _, p := range pts {
			n := Node{p.Add(cur.p), &cur}
			if n.p.x < 0 || n.p.x >= len(inputs[0]) || n.p.y < 0 || n.p.y >= len(inputs) {
				continue
			}
			if _, ok := added[n.p]; !ok {
				if (int(inputs[n.p.y][n.p.x]) - int(inputs[cur.p.y][cur.p.x])) >= -1 {
					added[n.p] = true
					queue = append(queue, n, cur)
					if inputs[n.p.y][n.p.x] == 'a' {
						node = n
						break foundEnd
					}
				}
			}
		}
	}

	var ans int
	for {
		if node.parent == nil {
			break
		}
		node = *node.parent
		ans++
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
