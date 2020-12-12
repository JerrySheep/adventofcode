package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func solution() int {
	result := make([]string, 0)

	for _, filename := range os.Args[1:] {
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			continue
		}

		for _, line := range strings.Split(string(data), "\n") {
			if line != "" {
				result = append(result, line)
			}
		}
	}

	store := make([]int, len(result))
	index := 0
	count := 0

	operation := make([]string, 0)
	offset := make([]string, 0)

	for i := 0; i < len(result); i++ {
		temp := make([]string, 0)
		for _, line := range strings.Split(result[i], " "){
			temp = append(temp, line)
		}
		operation = append(operation, temp[0])
		offset = append(offset, temp[1])
	}

	indexStore := make([]int, 0)

	for {
		if store[index] == 1 {
			break
		}
		indexStore = append(indexStore, index)
		store[index] = 1
		number, _ := strconv.Atoi(offset[index][1:])

		if operation[index] == "acc" {
			if offset[index][0] == '+' {
				count += number
			} else {
				count -= number
			}
			index++
		} else if operation[index] == "jmp" {
			if offset[index][0] == '+' {
				index += number
			} else {
				index -= number
			}
		} else {
			index++
		}
	}

	for i:= 0; i < len(indexStore); i++ {
		if operation[indexStore[i]] == "acc" {
			continue
		} else if operation[indexStore[i]] == "jmp" {
			operation[indexStore[i]] = "nop"
		} else if operation[indexStore[i]] == "nop" {
			operation[indexStore[i]] = "jmp"
		}

		store = make([]int, len(result))
		index = 0
		count = 0

		for {
			if index >= len(result){
				return count
			}

			if store[index] == 1 {
				break
			}
			store[index] = 1
			number, _ := strconv.Atoi(offset[index][1:])

			if operation[index] == "acc" {
				if offset[index][0] == '+' {
					count += number
				} else {
					count -= number
				}
				index++
			} else if operation[index] == "jmp" {
				if offset[index][0] == '+' {
					index += number
				} else {
					index -= number
				}
			} else {
				index++
			}
		}

		if operation[indexStore[i]] == "jmp" {
			operation[indexStore[i]] = "nop"
		} else if operation[indexStore[i]] == "nop" {
			operation[indexStore[i]] = "jmp"
		}
	}

	return count
}

func main(){
	answer := solution()
	fmt.Println(answer)
}