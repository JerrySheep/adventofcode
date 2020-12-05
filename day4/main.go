package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func byrCheck(s string) bool {
	year, ok := strconv.Atoi(s)
	if ok == nil {
		if year >= 1920 && year <= 2002 {
			return true
		} else {
			return false
		}
	} else {
		return false
	}
}

func iyrCheck(s string) bool {
	year, ok := strconv.Atoi(s)
	if ok == nil {
		if year >= 2010 && year <= 2020 {
			return true
		} else {
			return false
		}
	} else {
		return false
	}
}

func eyrCheck(s string) bool {
	year, ok := strconv.Atoi(s)
	if ok == nil {
		if year >= 2020 && year <= 2030 {
			return true
		} else {
			return false
		}
	} else {
		return false
	}
}

func hgtCheck(s string) bool {
	if s[len(s) - 2:] == "cm" {
		height, ok := strconv.Atoi(s[0:len(s) - 2])
		if ok == nil {
			if height >= 150 && height <= 193 {
				return true
			} else {
				return false
			}
		} else {
			return false
		}
	} else if s[len(s) - 2:] == "in" {
		height, ok := strconv.Atoi(s[0 : len(s) - 2])
		if ok == nil {
			if height >= 59 && height <= 76 {
				return true
			} else {
				return false
			}
		} else {
			return false
		}
	} else {
		return false
	}
}

func hclCheck(s string) bool {
	if strings.HasPrefix(s, "#") {
		if len(s) == 7 {
			for i := 1; i < 7; i++ {
				if (s[i] >= '0' && s[i] <= '9') || (s[i] >= 'a' && s[i] <= 'f') {
					continue
				} else {
					return false
				}
			}
			return true
		} else {
			return false
		}
	} else {
		return false
	}
}

func eclCheck(s string) bool {
	if s == "amb" || s == "blu" || s == "brn" || s == "gry" || s == "grn" || s == "hzl" || s == "oth" {
		return true
	} else {
		return false
	}
}

func pidCheck(s string) bool {
	if len(s) == 9 {
		for i := 0; i < 9; i++ {
			if s[i] >= '0' && s[i] <= '9' {
				continue
			} else {
				return false
			}
		}
		return true
	} else {
		return false
	}
}

func solution() int {
	result := make([]string, 0)

	for _, filename := range os.Args[1:] {
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			continue
		}
		for _, line := range strings.Split(string(data), "\n") {
			result = append(result, line)
		}
	}

	store := make([]string, 0)

	for i := 0; i < len(result); i++ {
		store = append(store, result[i])
		for result[i + 1] != "" && i + 1 < len(result) {
			i++
			store[len(store) - 1] = store[len(store) - 1] + " " + result[i]
		}
		i++
	}

	validNumber := 0

	for i := 0; i < len(store); i++ {
		splitSentence := strings.Split(store[i], " ")
		if len(splitSentence) == 8 || (len(splitSentence) == 7 && !strings.Contains(store[i], "cid")){
			var byr string
			var iyr string
			var eyr string
			var hgt string
			var hcl string
			var ecl string
			var pid string
			for j := 0; j < len(splitSentence); j++ {
				if strings.HasPrefix(splitSentence[j], "byr") {
					byr = splitSentence[j][4:]
				} else if strings.HasPrefix(splitSentence[j], "iyr") {
					iyr = splitSentence[j][4:]
				} else if strings.HasPrefix(splitSentence[j], "eyr") {
					eyr = splitSentence[j][4:]
				} else if strings.HasPrefix(splitSentence[j], "hgt") {
					hgt = splitSentence[j][4:]
				} else if strings.HasPrefix(splitSentence[j], "hcl") {
					hcl = splitSentence[j][4:]
				} else if strings.HasPrefix(splitSentence[j], "ecl") {
					ecl = splitSentence[j][4:]
				} else if strings.HasPrefix(splitSentence[j], "pid") {
					pid = splitSentence[j][4:]
				}
			}

			if byrCheck(byr) && iyrCheck(iyr) && eyrCheck(eyr) && hgtCheck(hgt) && hclCheck(hcl) && eclCheck(ecl) && pidCheck(pid) {
				validNumber++
			}
		}
	}

	return validNumber
}

func main(){
	answer := solution()
	fmt.Println(answer)
}
