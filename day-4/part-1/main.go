package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	passport := make(map[string]string)
	validPass := 0
	file, err := os.Open("input.txt")
	if err != nil {
		log.Println(err)
		return
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if len(line) == 0 {
			if len(passport) >= 7 {
				if isValidPassPort(passport) {
					validPass++
				}
			}
			passport = make(map[string]string)
			continue
		}
		for _, val := range strings.Split(line, " ") {
			data := strings.Split(val, ":")
			passport[data[0]] = data[1]
		}
	}
	if isValidPassPort(passport) {
		validPass++
	}
	fmt.Println(validPass)
}
func isValidPassPort(pp map[string]string) bool {
	requiredKeys := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}
	if len(pp) == 8 {
		return true
	}
	_, ok := pp["cid"]
	if ok {
		return false
	}
	for _, k := range requiredKeys {
		_, ok := pp[k]
		if ok {
			return true
		}
	}
	return false
}
