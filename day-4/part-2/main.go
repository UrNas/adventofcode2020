package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
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
		return validWithRules(pp)
	}
	if len(pp) == 7 {
		_, ok := pp["cid"]
		if ok {
			return false
		}
		for _, k := range requiredKeys {
			_, ok := pp[k]
			if ok {
				return validWithRules(pp)
			}
		}
	}
	return false
}
func validWithRules(pp map[string]string) bool {
	requiredKeys := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}
	for _, rk := range requiredKeys {
		value := pp[rk]
		switch rk {
		case "byr":
			byr, err := strconv.Atoi(value)
			if err != nil {
				return false
			}
			if byr < 1920 || byr > 2002 {
				return false
			}
		case "iyr":
			iyr, err := strconv.Atoi(value)
			if err != nil {
				return false
			}
			if iyr < 2010 || iyr > 2020 {
				return false
			}
		case "eyr":
			eyr, err := strconv.Atoi(value)
			if err != nil {
				return false
			}
			if eyr < 2020 || eyr > 2030 {
				return false
			}
		case "hgt":
			matched, _ := regexp.MatchString(`^[0-9]+(cm|in)$`, value)
			if !matched {
				return false
			}
			if matched {
				scm := strings.Split(value, "cm")
				if len(scm) > 1 {
					if !matchHieght(value, "cm") {
						return false
					}
				} else {
					if !matchHieght(value, "in") {
						return false
					}
				}
			}
		case "hcl":
			if !matchHColor(value) {
				return false
			}
		case "ecl":
			if !matchEyeCol(value) {
				return false
			}
		case "pid":
			if !matchPostID(value) {
				return false
			}
		}
	}
	return true
}
func matchHieght(v string, key string) bool {
	hieght := map[string][]int{
		"cm": []int{150, 193},
		"in": []int{59, 76},
	}
	d := strings.Split(v, key)[0]
	nd, err := strconv.Atoi(d)
	if err != nil {
		return false
	}
	if nd < hieght[key][0] || nd > hieght[key][1] {
		return false
	}
	return true
}
func matchHColor(v string) bool {
	matched, err := regexp.MatchString(`^#[0-9a-f]{6}$`, v)
	if err != nil {
		return false
	}
	if !matched {
		return false
	}
	return true
}
func matchEyeCol(v string) bool {
	var eyeCl = []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}
	for _, color := range eyeCl {
		if v == color {
			return true
		}
	}
	return false
}

func matchPostID(v string) bool {
	matched, err := regexp.MatchString(`^[0-9]{9}$`, v)
	if err != nil {
		return false
	}
	if matched {
		return true
	}
	return false
}
