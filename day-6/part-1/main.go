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
	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			countAnswers += len(answers)
			answers = make(map[rune]int)
		}
		for _, v := range text {
			answers[v]++
		}
	}
	countAnswers += len(answers)
	fmt.Println(countAnswers)
}
