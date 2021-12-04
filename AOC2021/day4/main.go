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

	squares := make([][]int, 0)
	numbers := make([]int, 0)

	for _, numStr := range strings.Split(result[0], ",") {
		num, _ := strconv.Atoi(numStr)
		numbers = append(numbers, num)
	}

	for i := 1; i < len(result); i ++ {
		temp := make([]int, 0)
		for _, numStr := range strings.Fields(result[i]){
			num, _ := strconv.Atoi(numStr)
			temp = append(temp, num)
		}
		squares = append(squares, temp)
	}

	mark := make([][]int, 0)
	for i := 0; i < len(squares); i++ {
		temp := make([]int, 5)
		copy(temp, squares[i])
		mark = append(mark, temp)
	}

	total := len(squares) / 5
	count := 0
	winIndex := make([]int, total)
	for i := 0; i < len(numbers); i++ {
		for j := 0; j < len(squares); j++ {
			if winIndex[j / 5] == 1 {
				continue
			}
			for k := 0; k < 5; k++ {
				if mark[j][k] != -1 && squares[j][k] == numbers[i]{
					mark[j][k] = -1
					if ifComplete(mark, j, k){
						count++
						winIndex[j / 5] = 1
						if count == total {
							return calculate(squares, mark, j) * numbers[i]
						}
					}
				}
			}
		}
	}

	return 0
}

func calculate(squares [][]int, mark [][]int, row int) int {
	startRow := row - (row % 5)

	unMarkedNumber := 0
	for i := startRow; i < startRow + 5; i++ {
		for j := 0; j < 5; j++ {
			if mark[i][j] != -1 {
				unMarkedNumber += squares[i][j]
			}
		}
	}
	return unMarkedNumber
}

func ifComplete(mark [][]int, row int, column int) bool{
	rowJudge := true
	columnJudge := true
	for i := 0; i < 5; i++ {
		if mark[row][i] != -1 {
			rowJudge = false
			break
		}
	}

	startRow := row - (row % 5)
	for i := startRow; i < startRow + 5; i++ {
		if mark[i][column] != -1 {
			columnJudge = false
			break
		}
	}

	return rowJudge || columnJudge
}

func main(){
	answer := solution()
	fmt.Println(answer)
}