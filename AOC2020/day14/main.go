package main

import (
	"fmt"
	"io/ioutil"
	"math"
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

	sumNumber := 0
	keyValue := make(map[int]int)
	for i := 0; i < len(result); i++ {
		maskResult := result[i]
		memResult := make([]string, 0)

		for strings.HasPrefix(result[i + 1], "mem") {
			memResult = append(memResult, result[i + 1])
			i++
			if i == len(result) - 1 {
				break
			}
		}

		maskSentence := strings.Split(maskResult, " ")
		mask := maskSentence[2]

		memoryIndex := make([]int, 0)
		memoryNumber := make([]int, 0)

		for j := 0; j < len(memResult); j++ {
			memorySentence := strings.Split(memResult[j], " ")
			number, _ := strconv.Atoi(memorySentence[0][4:len(memorySentence[0]) - 1])
			memoryIndex = append(memoryIndex, number)
			number, _ = strconv.Atoi(memorySentence[2])
			memoryNumber = append(memoryNumber, number)
		}

		//for j := 0; j < len(memResult); j++ {
		//	temp := 0
		//	for k := 35; k >= 0; k-- {
		//		if (mask[k] == 'X' && memoryNumber[j] % 2 != 0) || mask[k] == '1' {
		//			temp += int(math.Pow(2, float64(35-k)))
		//		}
		//
		//		memoryNumber[j] = memoryNumber[j] >> 1
		//	}
		//	memoryNumber[j] = temp
		//}
		//
		//for j := 0; j < len(memResult); j++ {
		//	keyValue[memoryIndex[j]] = memoryNumber[j]
		//}


		for j := 0; j < len(memResult); j++ {
			floatNumber := make([]int, 1)
			for k := 35; k >= 0; k-- {
				if mask[k] == '1' || (mask[k] == '0' && memoryIndex[j] % 2 != 0) {
					for m := 0; m < len(floatNumber); m++ {
						floatNumber[m] += int(math.Pow(2, float64(35-k)))
					}
				} else if mask[k] == 'X' {
					temp := len(floatNumber)
					for m := 0; m < temp; m++ {
						floatNumber = append(floatNumber, floatNumber[m] + int(math.Pow(2, float64(35-k))))
					}
				}

				memoryIndex[j] = memoryIndex[j] >> 1
			}

			for k := 0; k < len(floatNumber); k++ {
				keyValue[floatNumber[k]] = memoryNumber[j]
			}
		}
	}

	for _, value := range keyValue {
		sumNumber += value
	}

	return sumNumber
}

func main(){
	answer := solution()
	fmt.Println(answer)
}
