package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
)

func main() {
	var seatIDS []int
	file, err := os.Open("input.txt")
	defer file.Close()
	if err != nil {
		log.Println(err)
		return
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		row := getRow(text[7:])
		col := getColumn(text[:7])
		seatIDS = append(seatIDS, col*8+row)

	}
	sort.Ints(seatIDS)
	fmt.Println("Max SeatID: ", seatIDS[len(seatIDS)-1])
	fmt.Println("Your SeatID: ", getMySeat(seatIDS))
}

// Part one Section
func getColumn(s string) int {
	var region []int
	firstRegion := s[:1]
	if firstRegion == "B" {
		region = []int{int(math.Ceil(float64(127) / float64(2))), 127}
	} else {
		region = []int{0, 127 / 2}
	}
	for _, v := range s[1:] {
		if v == 70 {
			diff := float64(region[1] - region[0])
			region = []int{region[0], int(math.Floor((diff / 2))) + region[0]}
			continue
		}
		if v == 66 {
			diff := float64(region[1] - region[0])
			region = []int{region[0] + int(math.Ceil(diff/2)), region[1]}
			continue
		}
	}
	return region[0]
}
func getRow(s string) int {

	var row []int
	firstRow := s[:1]
	if firstRow == "R" {
		row = []int{4, 7}
	} else {
		row = []int{0, 3}
	}

	for _, v := range s[1:] {
		if v == 76 {
			diff := float64(row[1] - row[0])
			row = []int{row[0], int(math.Floor((diff / 2))) + row[0]}
			continue
		} else {
			diff := float64(row[1] - row[0])
			row = []int{row[0] + int(math.Ceil(diff/2)), row[1]}
			continue
		}
	}
	return row[0]
}

// part 2
func getMySeat(seats []int) int {
	var seat int
	for i, v := range seats {
		if v+1 != seats[i+1] {
			seat = v + 1
			break
		}
	}
	return seat
}
