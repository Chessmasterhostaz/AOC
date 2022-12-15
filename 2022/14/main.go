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
	rockMap := generateRockMap(inputs)
	println("Task 1:", problem1(rockMap))
	println("Task 2:", problem2(rockMap))
}

type Point struct {
	x, y int
}

func (p *Point) Add(p2 Point) Point {
	ax := p.x + p2.x
	ay := p.y + p2.y
	return Point{ax, ay}
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

var pts = map[byte]Point{'D': {0, 1}, 'R': {1, 1}, 'L': {-1, 1}}

var maxY = 0

func generateRockMap(inputs []string) map[Point]bool {
	var rocks = make(map[Point]bool)
	for _, v := range inputs {
		pl := strings.Split(v, " -> ")
		for i := 0; i < len(pl)-1; i++ {
			ep := Point{0, 0}
			sp := Point{0, 0}
			fmt.Sscanf(pl[i], "%d,%d", &sp.x, &sp.y)
			fmt.Sscanf(pl[i+1], "%d,%d", &ep.x, &ep.y)

			if i == 0 {
				rocks[sp] = true
			}

			dp := ep.Diff(sp)
			for ; dp.x != 0; dp.x -= dp.Sign().x {
				sp = sp.Add(dp.Sign())
				rocks[sp] = true
			}

			for ; dp.y != 0; dp.y -= dp.Sign().y {
				sp = sp.Add(dp.Sign())
				rocks[sp] = true

				if sp.y > maxY {
					maxY = sp.y
				}
			}
		}
	}
	return rocks
}
func problem1(rocks map[Point]bool) int {
	var sand = make(map[Point]bool)
	var path []Point
	path = append(path, Point{500, 0})
	for len(path) > 0 {
		head := path[len(path)-1]
		if head.y > maxY+2 {
			break
		}

		next := head.Add(pts['D'])
		if rocks[next] || sand[next] {
			next = head.Add(pts['L'])
		} else {
			path = append(path, next)
			continue
		}

		if rocks[next] || sand[next] {
			next = head.Add(pts['R'])
		} else {
			path = append(path, next)
			continue
		}

		if rocks[next] || sand[next] {
			path = path[:len(path)-1]
			sand[head] = true
		} else {
			path = append(path, next)
		}
	}

	return len(sand)
}

func problem2(rocks map[Point]bool) int {
	for x := -(maxY * 4); x < maxY*4; x++ {
		rocks[Point{x, maxY + 2}] = true
	}

	return problem1(rocks)
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
