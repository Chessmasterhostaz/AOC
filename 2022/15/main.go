package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"sort"
)

func main() {
	inputs := readInput()
	problem1(inputs)
	problem2(inputs)
}

func Sign(in int) int {
	if in < 0 {
		return -1
	}
	return 1
}

func Abs(in int) int {
	if in < 0 {
		return -1 * in
	}
	return in
}

type sensor struct {
	x, y, w int
}
type point struct {
	x, y int
}
type line struct {
	x1, x2 int
}

func problem1(inputs []string) {
	var ans int
	row := 2000000

	var sx, sy, bx, by int
	var sns []sensor
	var bcs = make(map[point]int)
	for _, v := range inputs {
		fmt.Sscanf(v, "Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d", &sx, &sy, &bx, &by)
		sns = append(sns, sensor{sx, sy, 1 + 2*(Abs(sx-bx)+Abs(sy-by))})
		bcs[point{bx, by}] = by
	}

	var lines []line
	min := 0
	max := 0
	for _, v := range sns {
		vw := v.width(row)
		if vw > 0 {
			lines = append(lines, line{v.x - vw/2, v.x + vw/2})
			if v.x+vw/2 > max {
				max = v.x + vw/2
			}
			if v.x-vw/2 < min {
				min = v.x - vw/2
			}
		}
	}

	sort.Slice(lines, func(i, j int) bool {
		return lines[i].x1 < lines[j].x1
	})

	ts := make([]bool, max-min+1)
	for i := min; i < max+1; i++ {
		for _, l := range lines {
			if i >= l.x1 && i <= l.x2 {
				if !ts[i-min] {
					ans++
				}
				ts[i-min] = true
				continue
			}
		}
	}

	for _, b := range bcs {
		if b == row {
			ans--
		}
	}

	fmt.Printf("Task 1: %d\n", ans)
}

func problem2(inputs []string) {
	var ans int

	max := 4000000

	var sx, sy, bx, by int
	var sns []sensor
	for _, v := range inputs {
		fmt.Sscanf(v, "Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d", &sx, &sy, &bx, &by)
		sns = append(sns, sensor{sx, sy, 1 + 2*(Abs(sx-bx)+Abs(sy-by))})
	}

main:
	for r := 0; r < max; r++ {
		var lines []line
		for _, v := range sns {
			vw := v.width(r)
			if vw > 0 {
				lines = append(lines, line{v.x - vw/2, v.x + vw/2})
			}
		}

		sort.Slice(lines, func(i, j int) bool {
			return lines[i].x1 < lines[j].x1
		})

		ma := 0
		for _, l := range lines {
			if l.x1 > ma {
				ans = (l.x1-1)*max + r
				break main
			} else if l.x2 > ma {
				ma = l.x2
			}

		}
	}

	fmt.Printf("Task 2: %d\n", ans)
}

func (s sensor) width(row int) int {
	if 2*Abs(s.y-row) < s.w {
		return s.w - 2*Abs(s.y-row)
	} else {
		return 0
	}
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
