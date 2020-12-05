package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type password struct {
	content  string
	min, max int
	char     string
}

func (p password) isValidPass() bool {
	existance := strings.Count(p.content, p.char)
	return existance <= p.max && existance >= p.min
}
func (p password) isValidPosition() bool {
	var fposition, sposition string
	existance := strings.Count(p.content, p.char)
	if existance <= 0 {
		return false
	}
	fposition = string(p.content[p.min-1])
	sposition = string(p.content[p.max-1])
	if fposition == p.char && sposition == p.char {
		return false
	}
	if fposition != p.char && sposition != p.char {
		return false
	}
	return true
}
func main() {
	var vPass int
	var vPositionPass int
	var passwords []password
	f, err := os.Open("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	s := bufio.NewScanner(f)
	for s.Scan() {
		d := strings.Split(s.Text(), " ")
		lenghtPass := strings.Split(d[0], "-")
		min, err := strconv.Atoi(lenghtPass[0])
		if err != nil {
			log.Fatal(err)
		}
		max, err := strconv.Atoi(lenghtPass[1])
		if err != nil {
			log.Fatal(err)
		}
		char := strings.Split(d[1], ":")[0]
		passwords = append(passwords, password{
			content: d[2],
			min:     min,
			max:     max,
			char:    char,
		})
	}
	for _, v := range passwords {
		if v.isValidPass() {
			vPass++
		}
	}
	for _, v := range passwords {
		if v.isValidPosition() {
			vPositionPass++
		}
	}
	fmt.Println(vPass)
	fmt.Println(vPositionPass)
}
