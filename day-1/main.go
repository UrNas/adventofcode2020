package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	var numbers []int
	f, err := os.Open("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	s := bufio.NewScanner(f)
	for s.Scan() {
		n, err := strconv.Atoi(s.Text())
		if err != nil {
			log.Fatalln(err)
		}
		numbers = append(numbers, n)
	}
	for _, v := range numbers {
		n, ok := isTotal2010(numbers, v)
		if ok {
			fmt.Println(n)
		}
	}
	for _, v := range numbers {
		n, ok := isTotalThree2010(numbers, v)
		if ok {
			fmt.Println(n)
		}
	}
}
func isTotal2010(list []int, n int) (int, bool) {
	for _, v := range list {
		if (n + v) == 2020 {
			return (n * v), true
		}
	}
	return 0, false
}
func isTotalThree2010(list []int, n int) (int, bool) {
	var product []int
	var total int
	for i, v := range list {
		if (v + n) < 2020 {
			total = v + n
			for _, s := range list[i+1:] {
				if total+s == 2020 {
					product = append(product, v, n, s)
					return product[0] * product[1] * product[2], true
				}
			}
		}
	}
	return 0, false
}
