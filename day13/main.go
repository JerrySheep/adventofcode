package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func min(first int, second int) int {
	if first < second {
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

	//timestap, ok := strconv.Atoi(result[0])
	//if ok != nil {
	//	fmt.Fprintf(os.Stderr, "%v\n", ok)
	//}
	busNumber := make([]int, 0)

	for _, line := range strings.Split(result[1], ",") {
		if line != "x" {
			number, ok := strconv.Atoi(line)
			if ok == nil{
				busNumber = append(busNumber, number)
			}
		} else if line == "x" {
			busNumber = append(busNumber, 0)
		}
	}

	index := 0
	temp := 1

	for i := 1; ; i++ {
		judge := true
		division := i + index
		for j := index; j < len(busNumber); j++ {
			if busNumber[j] != 0 {
				if (division % busNumber[j]) != 0 {
					judge = false
					break
				} else {
					temp *= busNumber[j]
					index = j + 1
				}
			}
			division++
		}
		if judge {
			return i
		} else {
			i += temp - 1
		}
	}
}

func main(){
	answer := solution()
	fmt.Println(answer)
}
