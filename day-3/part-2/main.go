package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// 5522401584
type slope struct {
	right int
	down  int
}

func main() {
	var areaMap []string
	multiTree := 1
	slopes := []slope{
		slope{
			right: 1,
			down:  1,
		},
		slope{
			right: 3,
			down:  1,
		},
		slope{
			right: 5,
			down:  1,
		},
		slope{
			right: 7,
			down:  1,
		},
		slope{
			right: 1,
			down:  2,
		},
	}
	f, err := os.Open("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	s := bufio.NewScanner(f)
	for s.Scan() {
		areaMap = append(areaMap, s.Text())
	}
	for _, v := range slopes {
		multiTree *= run(v.right, v.down, areaMap)

	}
	fmt.Println(multiTree)

}
func nextStep(steps int, pattern string) byte {
	step := pattern[steps]
	return step
}
func isTree(tree byte) bool {
	if tree == byte(35) {
		return true
	}
	return false
}
func run(right int, down int, treeMap []string) int {
	trakStep := 0
	counterTree := 0
	orginDown := down
	for down < len(treeMap) {
		v := treeMap[down]
		trakStep += right

		if trakStep >= len(v) {
			n := trakStep % len(v)
			step := nextStep(n, v)
			if isTree(step) {
				counterTree++
			}
			down += orginDown
			continue
		}
		step := nextStep(trakStep, v)
		if isTree(step) {
			counterTree++
		}
		down += orginDown

	}
	return counterTree
}
