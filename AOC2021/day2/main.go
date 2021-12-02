package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func solution() int{
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

	move_type := make([]string, 0)
	move_value := make([]int, 0)

	for i := 0; i < len(result); i++ {
		sentence := strings.Split(result[i], " ")
		move_type = append(move_type, sentence[0])
		num, _ := strconv.Atoi(sentence[1])
		move_value = append(move_value, num)
	}

	horizon := 0
	depth := 0
	aim := 0

	for i := 0; i < len(result); i++ {
		if move_type[i] == "forward" {
			horizon += move_value[i]
			depth += aim * move_value[i]
		} else if move_type[i] == "down" {
			//depth += move_value[i]
			aim += move_value[i]
		} else {
			//depth -= move_value[i]
			aim -= move_value[i]
		}
	}

	return horizon * depth
}

func main(){
	answer := solution()
	fmt.Println(answer)
}