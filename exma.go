package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var requiredAttributes = []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}

func isValid(passport map[string]string) bool {
	for _, attribute := range requiredAttributes {
		if _, exists := passport[attribute]; !exists {
			return false
		}
	}

	return true
}

func exists(slice []string, val string) bool {
	for _, item := range slice {
		if item == val {
			return true
		}
	}
	return false
}

func validateRange(val string, lower, upper int) bool {
	num, err := strconv.Atoi(val)
	return err == nil && num >= lower && num <= upper
}

func isValidV2(passport map[string]string) bool {
	for attr, val := range passport {
		switch attr {
		case "byr":
			if !validateRange(val, 1920, 2002) {
				return false
			}
		case "iyr":
			if !validateRange(val, 2010, 2020) {
				return false
			}
		case "eyr":
			if !validateRange(val, 2020, 2030) {
				return false
			}

		case "hgt":
			if strings.HasSuffix(val, "cm") {
				if !validateRange(strings.TrimSuffix(val, "cm"), 150, 193) {
					return false
				}
			} else if strings.HasSuffix(val, "in") {
				if !validateRange(strings.TrimSuffix(val, "in"), 59, 76) {
					return false
				}
			} else {
				return false
			}

		case "hcl":
			if match, _ := regexp.MatchString("^#[0-9a-f]{6}$", val); !match {
				return false
			}

		case "ecl":
			eyeColors := []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}
			if !exists(eyeColors, val) {
				return false
			}

		case "pid":
			if match, _ := regexp.MatchString("^[0-9]{9}$", val); !match {
				return false
			}
		}
	}

	return true
}

func main() {
	file, err := os.Open("input.txt")
	defer file.Close()
	if err != nil {
		log.Println(err)
		return
	}
	reader := bufio.NewReader(file)
	passport := make(map[string]string, 8)

	numValid, numValidV2 := 0, 0
	for {
		var line string
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}

		line = strings.TrimSpace(line)

		if len(line) == 0 {
			// passport complete
			if isValid(passport) {
				numValid++
				if isValidV2(passport) {
					numValidV2++
				}
			}
			// reset passport
			passport = make(map[string]string, 8)
		} else {
			parts := strings.Split(line, " ")
			for _, part := range parts {
				attribute := strings.Split(part, ":")
				passport[attribute[0]] = attribute[1]
			}
		}
	}

	// last one is special snowflake, as file does not end with 2 newlines
	if isValid(passport) {
		numValid++
		if isValidV2(passport) {
			numValidV2++
		}
	}

	fmt.Println(numValid, numValidV2)
}
