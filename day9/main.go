package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strconv"
	"strings"
)

func solution() int {
	result := make([]int, 0)

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

	var falseNumber int
	var index int
	for i := 25; i < len(result); i++ {
		temp := make([]int, 0)
		for j := i - 1; j > i - 26; j-- {
			for k := j - 1; k > i - 26; k-- {
				temp = append(temp, result[j] + result[k])
			}
		}

		judge := false

		for j := 0; j < len(temp); j++ {
			if result[i] == temp[j] {
				judge = true
				break
			}
		}

		if judge {
			continue
		} else {
			falseNumber = result[i]
			index = i
			break
		}
	}

	for i := 0; i < index; i++ {
		temp := make([]int, 0)
		sum := 0
		for j := i; j < index; j++ {
			sum += result[j]
			if sum == falseNumber {
				sort.Ints(temp)
				return temp[0] + temp[len(temp) - 1]
			}
			if sum > falseNumber {
				break
			}
			temp = append(temp, result[j])
		}
	}

	return falseNumber
}

func main(){
	answer := solution()
	fmt.Println(answer)
}
