package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	countAnswers := 0
	answers := make(map[rune]int)
	file, err := os.Open("input.txt")
	defer file.Close()
	if err != nil {
		log.Println(err)
		return
	}
	scanner := bufio.NewScanner(file)
	n := 0
	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			for _, v := range answers {
				if v == n {
					countAnswers++
				}
			}

			answers = make(map[rune]int)
			n = 0
		}
		for _, v := range text {
			answers[v]++
		}
		if text != "" {
			n++
		}
	}
	for _, v := range answers {
		if v == n {
			countAnswers++
		}
	}
	fmt.Println(countAnswers)
}
