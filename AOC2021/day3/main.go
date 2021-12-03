package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"os"
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

	oxygen_result := make([]string, len(result))
	co2_result := make([]string, len(result))
	copy(oxygen_result, result)
	copy(co2_result, result)

	index := 0
	for len(oxygen_result) > 1 {
		count := 0
		for i := 0; i < len(oxygen_result); i++{
			if oxygen_result[i][index] == '0' {
				count -= 1
			} else {
				count += 1
			}
		}

		for i := 0; i < len(oxygen_result); i++{
			if count >= 0 {
				if oxygen_result[i][index] == '0' {
					oxygen_result = append(oxygen_result[:i], oxygen_result[i + 1:]...)
					i--
				}
			} else {
				if oxygen_result[i][index] == '1' {
					oxygen_result = append(oxygen_result[:i], oxygen_result[i + 1:]...)
					i--
				}
			}
		}
		index++
	}

	index = 0
	for len(co2_result) > 1 {
		count := 0
		for i := 0; i < len(co2_result); i++{
			if co2_result[i][index] == '0' {
				count -= 1
			} else {
				count += 1
			}
		}

		for i := 0; i < len(co2_result); i++{
			if count >= 0 {
				if co2_result[i][index] == '1' {
					co2_result = append(co2_result[:i], co2_result[i + 1:]...)
					i--
				}
			} else {
				if co2_result[i][index] == '0' {
					co2_result = append(co2_result[:i], co2_result[i + 1:]...)
					i--
				}
			}
		}
		index++
	}

	oxygen := 0
	co2 := 0
	count := 0.0
	for i := len(oxygen_result[0]) - 1; i >= 0; i-- {
		if oxygen_result[0][i] == '1' {
			oxygen += int(math.Pow(2, count))
		}
		if co2_result[0][i] == '1' {
			co2 += int(math.Pow(2, count))
		}
		count++
	}

	return oxygen * co2



	//store := make([]int, len(result[0]))
	//for i := 0; i < len(result); i++ {
	//	for j := 0; j < len(result[i]); j++ {
	//		if result[i][j] == '0' {
	//			store[j] -= 1
	//		} else {
	//			store[j] += 1
	//		}
	//	}
	//}
	//
	//gamma := 0
	//epsilon := 0
	//count := 0.0
	//for i := len(store) - 1; i >= 0; i-- {
	//	if store[i] > 0 {
	//		gamma += int(math.Pow(2, count))
	//	} else {
	//		epsilon += int(math.Pow(2, count))
	//	}
	//	count++
	//}

	//return gamma * epsilon
}

func main(){
	answer := solution()
	fmt.Println(answer)
}