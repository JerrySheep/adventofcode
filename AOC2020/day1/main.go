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
		for _, line := range strings.Split(string(data), "\n") {
			number, ok := strconv.Atoi(line)
			if ok == nil {
				result = append(result, number)
			}
		}
	}

	sort.Ints(result)
	n := len(result)
	for k := 0; k < n; k++ {
		number := 2020 - result[k]
		for i := 0; i < n; i++ {
			if i == k {
				continue
			}
			for j := n - 1; j > i; j-- {
				if j == k {
					continue
				}
				if (result[i] + result[j]) > number {
					continue
				} else if (result[i] + result[j]) < number {
					break
				} else {
					return result[i] * result[j] * result[k]
				}
			}
		}
	}

	return 0
}

func main() {
	answer := solution()
	fmt.Println(answer)
}
