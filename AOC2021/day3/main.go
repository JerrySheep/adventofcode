package main

import (
	"fmt"
	"io/ioutil"
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

	oxygen_result = dealResult(oxygen_result, '1', '0')
	co2_result = dealResult(co2_result, '0', '1')

	oxygen := 0
	co2 := 0
	count := 0.0
	for i := 0; i < len(oxygen_result[0]); i++ {
		oxygen <<= 1
		co2 <<= 1
		if oxygen_result[0][i] == '1' {
			oxygen += 1
		}
		if co2_result[0][i] == '1' {
			co2 += 1
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

func dealResult(result []string, staySymbol byte, deleteSymbol byte) []string{
	index := 0
	for len(result) > 1 {
		count := 0
		for i := 0; i < len(result); i++{
			if result[i][index] == '1' {
				count += 1
			} else {
				count -= 1
			}
		}

		for i := 0; i < len(result); i++{
			if count >= 0 {
				if result[i][index] == deleteSymbol {
					result = append(result[:i], result[i + 1:]...)
					i--
				}
			} else if count < 0 {
				if result[i][index] == staySymbol {
					result = append(result[:i], result[i + 1:]...)
					i--
				}
			}
		}
		index++
	}

	return result
}

func main(){
	answer := solution()
	fmt.Println(answer)
}