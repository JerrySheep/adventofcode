package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func bagCalculate(result []string, bag string) int{
	count := 0
	for j := 0; j < len(result); j++ {
		if strings.HasPrefix(result[j], bag) && !strings.Contains(result[j], " 0 "){
			temp := make([]string, 0)
			bagSeries := make([]string, 0)
			bagNumber := make([]int, 0)

			for _, line := range strings.Split(result[j], " ") {
				temp = append(temp, line)
			}

			for k := 4; k < len(temp); k += 4 {
				number, _ := strconv.Atoi(temp[k])
				bagNumber = append(bagNumber, number)
				bagSeries = append(bagSeries, temp[k + 1])
				bagSeries[len(bagSeries) - 1] += " " + temp[k + 2]
			}

			for k := 0; k < len(bagNumber); k++ {
				count += bagNumber[k] + bagNumber[k] * bagCalculate(result, bagSeries[k])
			}
			break
		} else if strings.HasPrefix(result[j], bag) && strings.Contains(result[j], " 0 ") {
			break
		}
	}

	return count
}

func solution() int{
	result := make([]string, 0)

	for _, filename := range os.Args[1:] {
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			continue
		}

		for _, line := range strings.Split(string(data), "\n"){
			if line != "" {
				result = append(result, line)
			}
		}
	}

	//store := make([]string, 0)
	//store = append(store, "shiny gold")
	//s := ""
	//count := 0

	//for i := 0; i < len(store); i++ {
	//	s += store[i] + " "
	//	for j := 0; j < len(result); j++ {
	//		if strings.Contains(result[j], store[i]) && !strings.HasPrefix(result[j], store[i]){
	//			if !strings.Contains(s, result[j][0: strings.Index(result[j], "bags") - 1]) {
	//				count++
	//				store = append(store, result[j][0: strings.Index(result[j], "bags") - 1])
	//				s += result[j][0: strings.Index(result[j], "bags") - 1] + " "
	//				//fmt.Println(result[j][0: strings.Index(result[j], "bags") - 1])
	//			}
	//		}
	//	}
	//}
	return bagCalculate(result, "shiny gold")
}

func main(){
	answer := solution()
	fmt.Println(answer)
}
