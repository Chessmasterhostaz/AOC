package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"path/filepath"
	"runtime"
)

func main() {
	inputs := readInput()
	generateDroplets(inputs)
	problem1()
	problem2()
}

type Point struct {
	x, y, z int
}

func (p *Point) Add(p2 Point) Point {
	nx := p.x + p2.x
	ny := p.y + p2.y
	nz := p.z + p2.z
	return Point{nx, ny, nz}
}

var neightbours = []Point{
	{-1, 0, 0},
	{1, 0, 0},
	{0, -1, 0},
	{0, 1, 0},
	{0, 0, -1},
	{0, 0, 1},
}

var minx, miny, minz = math.MaxInt, math.MaxInt, math.MaxInt
var maxx, maxy, maxz = 0, 0, 0
var droplets = make(map[Point]bool)

func generateDroplets(inputs []string) {
	for _, v := range inputs {
		var x, y, z int
		fmt.Sscanf(v, "%d,%d,%d", &x, &y, &z)
		droplets[Point{x, y, z}] = true
		if x > maxx {
			maxx = x
		} else if x < minx {
			minx = x
		}

		if y > maxy {
			maxy = y
		} else if y < miny {
			miny = y
		}

		if z > maxz {
			maxz = z
		} else if z < minz {
			minz = z
		}
	}
}

var knownPockets = make(map[Point]bool)

func problem1() {
	var ans int
	for d := range droplets {
		for _, n := range neightbours {
			if _, ok := droplets[d.Add(n)]; !ok {
				ans++
			}
		}
	}
	fmt.Printf("Task 1: %d\n", ans)
}

func problem2() {
	var ans int

	for d := range droplets {
		for _, n := range neightbours {
			if _, ok := droplets[d.Add(n)]; !ok {
				ans++
			}
		}
	}

	for x := minx; x < maxx; x++ {
		for y := miny; y < maxy; y++ {
			for z := minz; z < maxz; z++ {
				p := Point{x, y, z}
				if _, ok := droplets[p]; !ok {

					pocket := pocketFinder(p)
					if pocket != nil {
						vis := false
						for _, p := range pocket {
							if _, ok := knownPockets[p]; ok {
								vis = true
							}
							knownPockets[p] = true
						}
						if !vis {
						}
						ans -= getPocketSize(pocket)
					}
				}
			}
		}
	}

	fmt.Printf("Task 2: %d\n", ans)
}

func pocketFinder(p Point) []Point {
	var visited = make(map[Point]bool)
	var pts = []Point{}
	pts = append(pts, p)
	for idx := 0; idx < len(pts); idx++ {
		if idx == len(pts) {
			return pts
		}

		p := pts[idx]
		if idx == 0 {
			if _, ok := knownPockets[p]; ok {
				return nil
			}

			visited[p] = true
		}

		if p.x > maxx || p.x < minx || p.y > maxy || p.y < miny || p.z > maxz || p.z < minz {
			return nil
		}

		for _, n := range neightbours {
			np := p.Add(n)
			if _, ok := visited[np]; !ok {
				if _, ok := droplets[np]; !ok {
					pts = append(pts, np)
				}
				visited[np] = true
			}
		}
	}
	return pts
}

func getPocketSize(pts []Point) int {

	retval := 0
	for _, p := range pts {
		for _, n := range neightbours {
			if _, ok := droplets[p.Add(n)]; ok {
				retval++
			} else if len(pts) == 1 {
				fmt.Println(p.Add(n))
			}
		}
	}
	return retval
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
