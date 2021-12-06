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

	num := make([]int, 0)
	for _, num_str := range strings.Split(result[0], ",") {
		number, _ := strconv.Atoi(num_str)
		num = append(num, number)
	}

	count := make([]int, 9)
	for i := 0; i < len(num); i++ {
		count[num[i]] += 1
	}

	for i := 0; i < 256; i++ {
		eightCount := 0
		sixCount := 0
		for j := 0; j < len(count); j++ {
			temp := count[j]
			if j == 0 {
				eightCount += temp
				sixCount += temp
			} else {
				count[j - 1] += temp
			}
			count[j] = 0
		}
		count[8] += eightCount
		count[6] += sixCount
	}

	total := 0
	for i := 0; i < len(count); i++ {
		total += count[i]
	}

	return total
}

func main(){
	answer := solution()
	fmt.Println(answer)
}