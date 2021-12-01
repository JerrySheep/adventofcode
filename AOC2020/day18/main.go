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

	calNumber := make([]int, 0)

	for i := 0; i < len(result); i++ {
		calculateSentence := result[i]
		temp := make([]string, 0)

		for j := 0; j < len(calculateSentence); j++ {
			if calculateSentence[j] != ' ' && (calculateSentence[j] < '0' || calculateSentence[j] > '9') {
				temp = append(temp, string(calculateSentence[j]))
			} else if calculateSentence[j] != ' ' {
				numberSentence := ""
				for j < len(calculateSentence) && calculateSentence[j] >= '0' && calculateSentence[j] <= '9' {
					numberSentence += string(calculateSentence[j])
					j++
				}
				temp = append(temp, numberSentence)
				j--
			}
		}

		stack := make([]string, 0)
		postExpression := make([]string, 0)
		for j := 0; j < len(temp); j++ {
			if temp[j] == "+" ||  temp[j] == "*" || temp[j] == "("{
				//if len(stack) == 0 || temp[j] == "("{
				//	stack = append(stack, temp[j])
				//} else {
				//	for len(stack) > 0 && stack[len(stack) - 1] != "(" {
				//		postExpression = append(postExpression, stack[len(stack) - 1])
				//		stack = stack[0 : len(stack) - 1]
				//	}
				//	stack = append(stack, temp[j])
				//}
				if len(stack) == 0 || temp[j] == "("{
					stack = append(stack, temp[j])
				} else if temp[j] == "+" {
					for len(stack) > 0 && stack[len(stack) - 1] != "(" && stack[len(stack) - 1] != "*" {
						postExpression = append(postExpression, stack[len(stack) - 1])
						stack = stack[0 : len(stack) - 1]
					}
					stack = append(stack, temp[j])
				} else {
					for len(stack) > 0 && stack[len(stack) - 1] != "(" {
						postExpression = append(postExpression, stack[len(stack) - 1])
						stack = stack[0 : len(stack) - 1]
					}
					stack = append(stack, temp[j])
				}
			} else if temp[j] == ")" {
				for len(stack) > 0 && stack[len(stack) - 1] != "(" {
					postExpression = append(postExpression, stack[len(stack)-1])
					stack = stack[0 : len(stack) - 1]
				}
				if len(stack) == 0 {
					continue
				} else {
					stack = stack[0 : len(stack) - 1]
				}
			} else {
				postExpression = append(postExpression, temp[j])
			}
		}
		for len(stack) > 0 {
			postExpression = append(postExpression, stack[len(stack)-1])
			stack = stack[0 : len(stack)-1]
		}

		numbers := make([]int, 0)
		for j := 0; j < len(postExpression); j++ {
			if postExpression[j] == "+" || postExpression[j] == "*" {
				second := numbers[len(numbers) - 1]
				numbers = numbers[0:len(numbers) - 1]
				first := numbers[len(numbers) - 1]
				numbers = numbers[0:len(numbers) - 1]

				if postExpression[j] == "+" {
					numbers = append(numbers, first + second)
				} else if postExpression[j] == "*" {
					numbers = append(numbers, first * second)
				}
			} else {
				number, _ := strconv.Atoi(postExpression[j])
				numbers = append(numbers, number)
			}
		}

		calNumber = append(calNumber, numbers[0])
	}

	sumNumber := 0
	for i := 0; i < len(calNumber); i++ {
		sumNumber += calNumber[i]
	}

	return sumNumber
}

func main(){
	answer := solution()
	fmt.Println(answer)
}
