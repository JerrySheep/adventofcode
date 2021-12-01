package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func solution() int{
	result := make([]int, 0)

	for _, filename := range os.Args[1:] {
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			continue
		}

		for _, line := range strings.Split(string(data), "\n") {
			num, ok := strconv.Atoi(line)
			if ok == nil {
				result = append(result, num)
			}
		}
	}

	count := 0

	for i := 1; i < len(result) - 2; i++ {
		if result[i + 2] > result[i - 1] {
			count++
		}
	}

	return count
}

func main(){
	answer := solution()
	fmt.Println(answer)
}
