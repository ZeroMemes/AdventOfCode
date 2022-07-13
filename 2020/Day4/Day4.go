package main

import (
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	bytes, _ := ioutil.ReadFile("2020/Day4/input.txt")
	split := strings.Split(string(bytes), "\n\n")

	handle(split, func(m map[string]string) bool {
		return hasRequiredKeys(m)
	})

	handle(split, func(m map[string]string) bool {
		if hasRequiredKeys(m) {
			if match, _ := regexp.MatchString("^(amb|blu|brn|gry|grn|hzl|oth)$", m["ecl"]); !match {
				return false
			}
			if match, _ := regexp.MatchString("^[0-9]{9}$", m["pid"]); !match {
				return false
			}
			if match, _ := regexp.MatchString("^#[0-9a-f]{6}$", m["hcl"]); !match {
				return false
			}

			i := m["hgt"][0:(len(m["hgt"]) - 2)]

			if strings.HasSuffix(m["hgt"], "cm") {
				if !checkRange(i, 150, 193) {
					return false
				}
			} else if strings.HasSuffix(m["hgt"], "in") {
				if !checkRange(i, 59, 76) {
					return false
				}
			} else {
				return false
			}

			return checkRange(m["byr"], 1920, 2002) &&
				checkRange(m["iyr"], 2010, 2020) &&
				checkRange(m["eyr"], 2020, 2030)
		}
		return false
	})
}

func checkRange(str string, min int, max int) bool {
	i, err := strconv.Atoi(str)
	return err == nil && i >= min && i <= max
}

func hasRequiredKeys(m map[string]string) bool {
	return contains(m, "byr") &&
		contains(m, "iyr") &&
		contains(m, "eyr") &&
		contains(m, "hgt") &&
		contains(m, "hcl") &&
		contains(m, "ecl") &&
		contains(m, "pid")
}

func handle(passports []string, validate func(map[string]string) bool) {
	valid := 0

	for _, passport := range passports {
		m := parsePassport(passport)
		if validate(m) {
			valid++
		}
	}

	println(valid)
}

func parsePassport(raw string) map[string]string {
	m := make(map[string]string)
	raw = strings.ReplaceAll(raw, "\n", " ")
	data := strings.Split(raw, " ")
	for _, d := range data {
		s := strings.Split(d, ":")
		m[s[0]] = s[1]
	}
	return m
}

func contains(m map[string]string, key string) bool {
	if _, ok := m[key]; ok {
		return true
	}
	return false
}
