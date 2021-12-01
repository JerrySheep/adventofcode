package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func isValid(number int, ruleRange []int) bool {
	if number < ruleRange[0] || number > ruleRange[3] || (number > ruleRange[1] && number < ruleRange[2]) {
		return false
	} else {
		return true
	}
}

func recursiveCal(validTickets [][]int, ruleRanges [][]int, isUsed []int, path []int, count int, pathStore []int) {
	if len(path) == len(ruleRanges) {
		copy(pathStore, path)
		return
	}
	for i := 0; i < len(ruleRanges); i++ {
		if isUsed[i] == 1 {
			continue
		}
		isUsed[i] = 1
		path = append(path, i)
		judge := true
		for j := 0; j < len(validTickets); j++ {
			if !isValid(validTickets[j][count], ruleRanges[i]) {
				judge = false
				break
			}
		}
		if judge {
			recursiveCal(validTickets, ruleRanges, isUsed, path, count + 1, pathStore)
		}
		path = path[0:len(path) - 1]
		isUsed[i] = 0
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

	rules := make([]string, 0)
	var myTicket string
	nearbyTickets := make([]string, 0)

	for i := 0; i < len(result); i++ {
		for result[i] != "" {
			rules = append(rules, result[i])
			i++
		}

		i += 2
		myTicket = result[i]

		i += 3
		for result[i] != "" && i < len(result){
			nearbyTickets = append(nearbyTickets, result[i])
			i++
		}
		break
	}

	store := make([]int, 1000)
	ruleRanges := make([][]int, 0)

	for i := 0; i < len(rules); i++ {
		rule := strings.Split(rules[i], " ")
		firstRange := strings.Split(rule[len(rule) - 3], "-")
		secondRange := strings.Split(rule[len(rule) - 1], "-")

		temp := make([]int, 0)
		firstNumber, _ := strconv.Atoi(firstRange[0])
		secondNumber, _ := strconv.Atoi(firstRange[1])
		thirdNumber, _ := strconv.Atoi(secondRange[0])
		forthNumber, _ := strconv.Atoi(secondRange[1])

		temp = append(temp, firstNumber)
		temp = append(temp, secondNumber)
		temp = append(temp, thirdNumber)
		temp = append(temp, forthNumber)


		for j := firstNumber; j <= secondNumber; j++ {
			store[j] = 1
		}
		for j := thirdNumber; j <= forthNumber; j++ {
			store[j] = 1
		}

		ruleRanges = append(ruleRanges, temp)
	}

	sumNumber := 0
	validTickets := make([][]int, 0)

	for i := 0; i < len(nearbyTickets); i++ {
		nearbyTicket := strings.Split(nearbyTickets[i], ",")
		temp := make([]int, 0)
		judge := true
		for j := 0; j < len(nearbyTicket); j++ {
			number, _ := strconv.Atoi(nearbyTicket[j])
			if store[number] != 1 {
				sumNumber += number
				judge = false
			} else {
				temp = append(temp, number)
			}
		}

		if judge {
			validTickets = append(validTickets, temp)
		}
	}

	isUsed := make([]int, len(validTickets[0]))
	count := 0
	path := make([]int, 0)
	pathStore := make([]int, len(ruleRanges))
	recursiveCal(validTickets, ruleRanges, isUsed, path, count, pathStore)

	tickets := strings.Split(myTicket, ",")
	ticketNumber := make([]int, 0)
	for i := 0; i < len(tickets); i++ {
		number, _ := strconv.Atoi(tickets[i])
		ticketNumber = append(ticketNumber, number)
	}

	multiNumber := 1
	for i := 0; i < len(pathStore); i++ {
		if pathStore[i] <= 5 {
			multiNumber *= ticketNumber[i]
		}
	}

	return multiNumber
}

func main(){
	answer := solution()
	fmt.Println(answer)
}
