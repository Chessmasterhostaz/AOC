package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
)

func main() {
	inputs := readInput()
	problem(inputs)
}

func problem(inputs []string) {
	files := make(map[string]int)
	maps := make(map[string]int)

	pwd := ""
	for i, v := range inputs {
		if v[0] == '$' {
			if v[2] == 'c' {
				if v[5:] == "/" {
					pwd = ""
				} else if v[5:] == ".." {
					ps := strings.Split(pwd, "/") // ps= PäronSplit
					pwd = pwd[:len(pwd)-1-len(ps[len(ps)-2])]
				} else {
					pwd = fmt.Sprintf("%s%s/", pwd, v[5:])
				}

			} else { // ls
				for idx := i + 1; idx < len(inputs); idx++ {
					row := inputs[idx]
					if row[0] == '$' {
						break
					} else if row[0] != 'd' {
						vals := strings.Split(row, " ")
						size, _ := strconv.Atoi(vals[0])
						files[fmt.Sprintf("%s%s/", pwd, vals[1])] = size
					}
				}
			}
		}
	}

	for dir, v := range files {
		if v != 0 {
			subdirs := strings.Split(dir[:len(dir)-1], "/")
			for i := len(subdirs) - 1; i >= 0; i-- {
				dir = dir[:(len(dir) - 1 - len(subdirs[i]))]
				if dir == "" {
					dir = "/"
				}
				maps[dir] += v
			}
		}
	}

	var ans int
	for _, v := range maps {
		if v <= 100000 {
			ans += v
		}
	}
	fmt.Printf("Task 1: %d\n", ans)

	// Task 2
	requiredSpace := maps["/"] - 40000000
	min := maps["/"]
	for _, v := range maps {
		if v > requiredSpace {
			if v < min {
				min = v
			}
		}
	}
	fmt.Printf("Task 2: %d\n", min)
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
