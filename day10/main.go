package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strconv"
	"strings"
)

func sliceCompare(first []int, second []int) bool {
	for i := 0; i < len(first); i++ {
		if first[i] != second[i] {
			return false
		}
	}

	return true
}

func numberCalculate(result []int, index int, validNumber *int, depth int, lastResult []int){
	if index != len(result) - 1 {
		temp := make([]int, len(result))
		copy(temp, result)
		if result[index + 1] - result[index] < 4 {
			if !sliceCompare(lastResult, result) {
				*validNumber++
				copy(lastResult, result)
			}
			numberCalculate(temp, index + 1, validNumber, depth + 1, lastResult)
		}

		result[index] = result[index - 1]
		copy(temp, result)
		if result[index + 1] - result[index] < 4 {
			if !sliceCompare(lastResult, result) {
				*validNumber++
				copy(lastResult, result)
			}
			numberCalculate(temp, index + 1, validNumber, depth + 1, lastResult)
		}
	}
}

func solution() int {
	result := make([]int, 0)
	result = append(result, 0)

	for _, filename := range os.Args[1:] {
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			continue
		}

		for _, line := range strings.Split(string(data), "\n"){
			number, ok := strconv.Atoi(line)
			if ok == nil {
				result = append(result, number)
			}
		}
	}

	sort.Ints(result)
	//onePlusNumber := 1
	//threePlusNumber := 1

	//for i := 1; i < len(result); i++ {
	//	if result[i] - result[i - 1] == 1 {
	//		onePlusNumber++
	//	} else if result[i] - result[i - 1] == 3{
	//		threePlusNumber++
	//	}
	//}

	mutliNumber := 1

	for i := 0; i < len(result) - 1; i++ {
		temp := make([]int, 0)

		for i + 1 < len(result) && result[i + 1] - result[i] == 1 {
			temp = append(temp, result[i])
			i++
		}
		if len(temp) != 0{
			temp = append(temp, result[i])
		} else{
			temp = append(temp, result[i])
		}

		if len(temp) > 2 {
			validNumber := 0
			lastResult := make([]int, len(temp))
			numberCalculate(temp, 1, &validNumber, 1, lastResult)
			mutliNumber *= validNumber
		}
	}

	return mutliNumber
	//return onePlusNumber * threePlusNumber
}

func main(){
	answer := solution()
	fmt.Println(answer)
}

