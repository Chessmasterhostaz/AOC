package main

import (
	"bufio"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

type Shape struct {
	ps []Point
}

type Point struct {
	x, y int
}

func (p *Point) Add(p2 Point) Point {
	ax := p.x + p2.x
	ay := p.y + p2.y
	return Point{ax, ay}
}

func (p *Point) Sub(p2 Point) Point {
	ax := p.x - p2.x
	ay := p.y - p2.y
	return Point{ax, ay}
}

var s1 = Shape{
	[]Point{
		{0, 0},
		{1, 0},
		{2, 0},
		{3, 0},
	},
}

var s2 = Shape{
	[]Point{
		{1, 0},
		{0, 1},
		{1, 1},
		{2, 1},
		{1, 2},
	},
}

var s3 = Shape{
	[]Point{
		{0, 0},
		{1, 0},
		{2, 0},
		{2, 1},
		{2, 2},
	},
}

var s4 = Shape{
	[]Point{
		{0, 0},
		{0, 1},
		{0, 2},
		{0, 3},
	},
}

var s5 = Shape{
	[]Point{
		{0, 0},
		{0, 1},
		{1, 0},
		{1, 1},
	},
}

var moves []Point
var immovables map[Point]bool

func main() {
	inputs := readInput()

	for i := 0; i < len(inputs[0]); i++ {
		moves = append(moves, Point{int(inputs[0][i]) - '=', 0})
	}

	println("Task 1:", problem(inputs, 2022))
	println("Task 2:", problem(inputs, 1000_000_000_000))
}

func problem(inputs []string, maxRounds int) int {
	var shapes = []Shape{s1, s2, s3, s4, s5}
	immovables = make(map[Point]bool)

	for i := 0; i < 9; i++ {
		immovables[Point{i, 0}] = true
	}

	height := 0
	prevHeight := 0

	shapeNr := 0
	moveNr := 0

	lastShapeNr := 0
	lastMoveNr := 0
	foundRepeat := false
	actedOnRepeat := false
	repeatCounter := 0
	repeat := 5
	repeatHeight := 0

	roundsLeft := 0
	for round := 0; round < maxRounds; round++ {

		// Increase wall h
		for h := height; h < height+10; h++ {
			immovables[Point{0, h}] = true
			immovables[Point{8, h}] = true
		}

		// Select shape and start position
		shape := shapes[shapeNr]
		shapeNr = (shapeNr + 1) % len(shapes)
		pos := Point{3, height + 4}

		collision := false
		for !collision {
			// Try moving left or right based on input
			np := pos.Add(moves[moveNr])
			for _, p := range shape.ps {
				nps := np.Add(p)
				if _, ok := immovables[nps]; ok {
					np = pos
					break
				}
			}
			moveNr = (moveNr + 1) % len(moves)
			pos = np

			// Try moving down
			np = np.Add(Point{0, -1})
			for _, s := range shape.ps {
				nps := np.Add(s)
				if _, ok := immovables[nps]; ok {
					collision = true
					break
				}
			}

			if !collision {
				pos = np
			}
		}

		// Add to saved structures
		for _, s := range shape.ps {
			nps := pos.Add(s)
			if nps.y > height {
				height = nps.y
			}
			immovables[nps] = true
		}

		// Try different values for repeat (Mutlitples of 5 since there are 5 shapes) until both moveNr
		// and shapeNr are the same after repeating
		if !foundRepeat && round > 1000 && round%repeat == 0 {
			if repeatCounter == 1 {
				if lastShapeNr == shapeNr && lastMoveNr == moveNr {
					repeatHeight = height - prevHeight
					foundRepeat = true
				} else {
					repeat += 5
					repeatCounter = 0
				}
			} else {
				repeatCounter++
			}
			lastShapeNr = shapeNr
			lastMoveNr = moveNr
			prevHeight = height
		}

		// Since we know repeats we can cut the number of rounds short
		if foundRepeat && !actedOnRepeat {
			actedOnRepeat = true
			roundsLeft = maxRounds - round
			round = maxRounds - roundsLeft%repeat
		}

	}

	height += roundsLeft / 1735 * repeatHeight
	return height
}

func contains(l []Point, p Point) bool {
	for _, v := range l {
		if v == p {
			return true
		}
	}
	return false
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
