package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func idCalculate(sentence string) int {
	start := 0
	end := 127
	var row int
	for i := 0; i < 6; i++ {
		if sentence[i] == 'F' {
			end = (start + end) / 2
		} else {
			start = (start + end) / 2 + 1
		}
	}

	if sentence[6] == 'F' {
		row = start
	} else {
		row = end
	}

	start = 0
	end = 7
	var col int

	for i := 7; i < 9; i++ {
		if sentence[i] == 'L' {
			end = (start + end) / 2
		} else {
			start = (start + end) / 2 + 1
		}
	}

	if sentence[9] == 'L' {
		col = start
	} else {
		col = end
	}

	return row * 8 + col
}

func max(first int, second int) int {
	if first >= second {
		return first
	} else {
		return second
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
			if line != "" {
				result = append(result, line)
			}
		}
	}

	maxBoardNumber := 0
	store := make([]int, 995)

	for i := 0; i < len(result); i++ {
		//maxBoardNumber = max(maxBoardNumber, idCalculate(result[i]))
		store[idCalculate(result[i])] = 1
	}

	for i := 994 - len(result); i < len(result); i++ {
		if store[i] == 0 {
			fmt.Println(i)
		}
	}

	return maxBoardNumber
}

func main(){
	answer := solution()
	fmt.Println(answer)
}
