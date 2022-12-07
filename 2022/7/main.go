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

	pwd := "/"
	for i, v := range inputs {
		if v[0] == '$' {
			if v[2] == 'c' {
				if v[5:] == "/" {
					pwd = "/"
				} else if v[5:] == ".." {
					ps := strings.Split(pwd, "/")
					pwd = pwd[:len(pwd)-1-len(ps[len(ps)-1])]
					if pwd == "" {
						pwd = "/"
					}
				} else {
					if pwd == "/" {
						pwd = fmt.Sprintf("/%s", v[5:])
					} else {
						pwd = fmt.Sprintf("%s/%s", pwd, v[5:])
					}
				}

			} else { // ls
				idx := i + 1
				for {
					row := inputs[idx]
					if row[0] == '$' {
						break
					}
					if row[0] == 'd' {
						if pwd == "/" {
							maps[fmt.Sprintf("/%s", row[4:])] = 0
						} else {
							maps[fmt.Sprintf("%s/%s", pwd, row[4:])] = 0
						}
					} else {
						vals := strings.Split(row, " ")
						size, _ := strconv.Atoi(vals[0])
						name := vals[1]

						if pwd == "/" {
							files[fmt.Sprintf("/%s", name)] = size
						} else {
							files[fmt.Sprintf("%s/%s", pwd, name)] = size
						}
					}
					idx++
					if idx == len(inputs) {
						break
					}
				}
			}
		}
	}

	for k, v := range files {
		if v != 0 {
			paths := strings.Split(k, "/")[1:]
			dirpath := k
			for i := len(paths) - 1; i >= 0; i-- {
				dirpath = dirpath[:(len(dirpath) - 1 - len(paths[i]))]
				if dirpath == "" {
					dirpath = "/"
				}
				maps[dirpath] += v
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
