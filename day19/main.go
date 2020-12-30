package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func storeMessages(ruleKeyValue map[int]string, store []string, index int) []string{
	rule := ruleKeyValue[index]
	ruleNumbers := strings.Split(rule, " ")
	if ruleNumbers[0] == "\"a\"" {
		for i := 0; i < len(store); i++ {
			store[i] += "a"
		}
		return store
	} else if ruleNumbers[0] == "\"b\"" {
		for i := 0; i < len(store); i++ {
			store[i] += "b"
		}
		return store
	} else {
		if !strings.Contains(rule, "|") {
			ruleOnly := make([]int, 0)
			count := 0
			for count < len(ruleNumbers) {
				number, _ := strconv.Atoi(ruleNumbers[count])
				ruleOnly = append(ruleOnly, number)
				count++
			}

			for i := 0; i < len(ruleOnly); i++ {
				store = storeMessages(ruleKeyValue, store, ruleOnly[i])
			}
		} else {
			ruleOne := make([]int, 0)
			ruleTwo := make([]int, 0)
			count := 0
			for ruleNumbers[count] != "|" {
				number, _ := strconv.Atoi(ruleNumbers[count])
				ruleOne = append(ruleOne, number)
				count++
			}
			count++

			for count < len(ruleNumbers) {
				number, _ := strconv.Atoi(ruleNumbers[count])
				ruleTwo = append(ruleTwo, number)
				count++
			}

			storeLeft := make([]string, len(store))
			storeRight := make([]string, len(store))
			copy(storeLeft, store)
			copy(storeRight, store)

			for i := 0; i < len(ruleOne); i++ {
				storeLeft = storeMessages(ruleKeyValue, storeLeft, ruleOne[i])
			}
			for i := 0; i < len(ruleTwo); i++ {
				storeRight = storeMessages(ruleKeyValue, storeRight, ruleTwo[i])
			}

			store = store[0:0]
			for i := 0; i < len(storeLeft); i++ {
				store = append(store, storeLeft[i])
			}
			for i := 0; i < len(storeRight); i++ {
				store = append(store, storeRight[i])
			}
		}
	}

	return store
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

	rules := make([]string, 0)
	candidates := make([]string, 0)

	index := 0
	for result[index] != "" {
		rules = append(rules, result[index])
		index++
	}

	ruleKeyValue := make(map[int]string, 0)
	for i := 0; i < len(rules); i++ {
		sentence := strings.Split(rules[i], ": ")
		number, _ := strconv.Atoi(sentence[0])
		ruleKeyValue[number] = sentence[1]
	}

	index++
	for result[index] != "" {
		candidates = append(candidates, result[index])
		index++
	}

	message := ""
	store := make([]string, 0)
	store = append(store, message)

	store = storeMessages(ruleKeyValue, store, 0)

	keyValue := make(map[string]int, 0)

	for i := 0; i < len(store); i++ {
		keyValue[store[i]] = 1
	}

	validNumber := 0
	for i := 0; i < len(candidates); i++ {
		if _, ok := keyValue[candidates[i]]; ok {
			validNumber++
		}
	}

	return validNumber
}

func main(){
	answer := solution()
	fmt.Println(answer)
}

