package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
)

func main() {
	inputs := readInput()
	fmt.Printf("Task 1: %d\n", problem(inputs, 2))
	fmt.Printf("Task 1: %d\n", problem(inputs, 10))
}

type Point struct {
	x, y int
}

func (p *Point) Add(p2 Point) {
	p.x = p.x + p2.x
	p.y = p.y + p2.y
}

func (p *Point) Diff(p2 Point) Point {
	dx := p.x - p2.x
	dy := p.y - p2.y
	return Point{dx, dy}
}

func (p *Point) Sign() Point {
	sx := 0
	if p.x > 0 {
		sx = 1
	} else if p.x < 0 {
		sx = -1
	}
	sy := 0
	if p.y > 0 {
		sy = 1
	} else if p.y < 0 {
		sy = -1
	}
	return Point{sx, sy}
}

func (p *Point) Euclidean() float64 {
	return math.Sqrt(float64(p.x*p.x + p.y*p.y))
}

var pts = map[byte]Point{'U': {0, 1}, 'R': {1, 0}, 'D': {0, -1}, 'L': {-1, 0}}

func problem(inputs []string, knots int) int {
	visits := make(map[Point]bool)
	var pos []Point
	for i := 0; i < knots; i++ {
		pos = append(pos, Point{0, 0})
	}

	for _, v := range inputs {
		steps, _ := strconv.Atoi(v[2:])
		for s := 0; s < steps; s++ {
			pos[0].Add(pts[v[0]])
			for i := 1; i < knots; i++ {
				diff := pos[i-1].Diff(pos[i])
				if diff.Euclidean() >= 2 {
					pos[i].Add(diff.Sign())
				}
				if i == knots-1 {
					visits[pos[i]] = true
				}
			}
		}
	}

	return len(visits)
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
