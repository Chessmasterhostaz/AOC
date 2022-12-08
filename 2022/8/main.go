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

type Point struct {
	x, y int
}

func problem1(inputs []string) {
	rows := len(inputs)
	cols := len(inputs[0])

	visibleTrees := make(map[Point]bool)

	northMax := []byte(inputs[0])
	southMax := []byte(inputs[cols-1])
	for y := 1; y < rows-1; y++ {
		row1 := inputs[y]
		row2 := inputs[len(inputs)-1-y]

		westMax := row1[0]
		eastMax := row1[len(row1)-1]
		for x := 1; x < cols-1; x++ {
			if row1[x] > westMax {
				westMax = row1[x]
				visibleTrees[Point{y, x}] = true
			}

			if row1[x] > northMax[x] {
				northMax[x] = row1[x]
				visibleTrees[Point{y, x}] = true
			}

			revX := len(row1) - x - 1
			if row1[revX] > eastMax {
				eastMax = row1[revX]
				visibleTrees[Point{y, revX}] = true
			}

			revY := rows - y - 1
			if row2[x] > southMax[x] {
				southMax[x] = row2[x]
				visibleTrees[Point{revY, x}] = true
			}
		}
	}
	ans := len(visibleTrees) + rows + rows + cols + cols - 4

	fmt.Printf("Task 1: %d\n", ans)
}

func problem2(inputs []string) {
	ans := 0

	for i, row := range inputs {
		for j, v := range row {
			treeHeight := int(v)

			valW := 0
			max := 0
			for west := j - 1; west >= 0; west-- {
				valW++
				if int(row[west]) > max {
					max = int(row[west])
					if max >= treeHeight {
						break
					}
				}
			}

			valE := 0
			max = 0
			for east := j + 1; east < len(row); east++ {
				valE++
				if int(row[east]) > max {
					max = int(row[east])
					if max >= treeHeight {
						break
					}
				}
			}

			valN := 0
			max = 0
			for north := i - 1; north >= 0; north-- {
				valN++
				if int(inputs[north][j]) > max {
					max = int(inputs[north][j])
					if max >= treeHeight {
						break
					}
				}
			}

			valS := 0
			max = 0
			for south := i + 1; south < len(inputs); south++ {
				valS++
				if int(inputs[south][j]) > max {
					max = int(inputs[south][j])
					if max >= treeHeight {
						break
					}
				}
			}

			val := valW * valE * valN * valS
			if val > ans {
				ans = val
			}
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
