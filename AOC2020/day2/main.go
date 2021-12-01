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

	lowestTimes := make([]int, 0)
	highestTimes := make([]int, 0)
	pwdBytes := make([]byte, 0)
	passward := make([]string, 0)
	for i := 0; i < len(result); i++ {
		splitSentence := strings.Split(result[i], " ")
		times := strings.Split(splitSentence[0], "-")
		lowNumber, _ := strconv.Atoi(times[0])
		highNumber, _ := strconv.Atoi(times[1])
		lowestTimes = append(lowestTimes, lowNumber)
		highestTimes = append(highestTimes, highNumber)

		pwdBytes = append(pwdBytes, splitSentence[1][0])
		passward = append(passward, splitSentence[2])
	}

	validPwdNumber := 0

	for i := 0; i < len(result); i++ {
		temp := 0
		if passward[i][lowestTimes[i] - 1] == pwdBytes[i] {
			temp ++
		}
		if passward[i][highestTimes[i] - 1] == pwdBytes[i] {
			temp ++
		}

		if temp == 1 {
			validPwdNumber++
		}
		//for j := 0; j < len(passward[i]); j++ {
		//	if passward[i][j] == pwdBytes[i] {
		//		temp++
		//	}
		//}

		//if temp >= lowestTimes[i] && temp <= highestTimes[i] {
		//	validPwdNumber++
		//}
	}

	return validPwdNumber
}

func main() {
	answer := solution()
	fmt.Println(answer)
}
