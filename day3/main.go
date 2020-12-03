package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func tree_calculate(result []string, first int, second int) int {
	n := len(result[0])
	start := 0
	treeNumber := 0
	for i := second; i < len(result); i += second {
		start += first
		if start >= n {
			start -= n
		}
		if result[i][start] == '#' {
			treeNumber++
		}
	}

	return treeNumber
}
func solution() int {
	result := make([]string, 0)
	for _, filename := range os.Args[1:] {
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
		}

		for _, line := range strings.Split(string(data), "\n") {
			if line != "" {
				result = append(result, line)
			}
		}
	}

	treeNumber := 1
	firstList := []int{1, 3, 5, 7, 1}
	secondList := []int{1, 1, 1, 1, 2}
	for i := 0; i < len(firstList); i++ {
		treeNumber *= tree_calculate(result, firstList[i] ,secondList[i])
	}

	return treeNumber
}

func main() {
	answer := solution()
	fmt.Println(answer)
}
