package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	var areaMap []string
	trakStep := 0
	counterTree := 0
	f, err := os.Open("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	s := bufio.NewScanner(f)
	for s.Scan() {
		areaMap = append(areaMap, s.Text())
	}
	for _, v := range areaMap[1:] {
		trakStep += 3
		if trakStep >= len(v) {
			n := trakStep % len(v)
			step := slope(n, v)
			if isTree(step) {
				counterTree++
			}
			continue
		}
		step := slope(trakStep, v)
		if isTree(step) {
			counterTree++

		}
	}
	fmt.Println(counterTree)
}
func slope(steps int, pattern string) byte {
	item := pattern[steps]
	return item
}
func isTree(tree byte) bool {
	if tree == byte(35) {
		return true
	}
	return false
}
