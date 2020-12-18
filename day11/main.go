package main

import (
	"fmt"
	"io/ioutil"
	"os"
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

	for {
		temp := make([]string, len(result))
		copy(temp, result)
		for i := 0; i < len(result); i++ {
			for j := 0; j < len(result[i]); j++ {
				occupiedNumber := 0
				//for k := max(i - 1, 0); k <= min(i + 1, len(result) - 1); k++ {
				//	for m := max(j - 1, 0); m <= min(j + 1, len(result[i]) - 1); m++ {
				//		if temp[k][m] == '#' {
				//			occupiedNumber++
				//		}
				//	}
				//}

				for k := i - 1; k >= 0; k-- {
					if temp[k][j] == 'L' {
						break
					}
					if temp[k][j] == '#' {
						occupiedNumber++
						break
					}
				}

				for k := i + 1; k < len(result); k++ {
					if temp[k][j] == 'L' {
						break
					}
					if temp[k][j] == '#' {
						occupiedNumber++
						break
					}
				}

				for k := j - 1; k >= 0; k-- {
					if temp[i][k] == 'L' {
						break
					}
					if temp[i][k] == '#' {
						occupiedNumber++
						break
					}
				}

				for k := j + 1; k < len(result[i]); k++ {
					if temp[i][k] == 'L' {
						break
					}
					if temp[i][k] == '#' {
						occupiedNumber++
						break
					}
				}

				for k := 1; i - k >= 0 && j - k >= 0; k++ {
					if temp[i - k][j - k] == 'L' {
						break
					}
					if temp[i - k][j - k] == '#' {
						occupiedNumber++
						break
					}
				}

				for k := 1; i + k < len(result) && j + k < len(result[i]); k++ {
					if temp[i + k][j + k] == 'L' {
						break
					}
					if temp[i + k][j + k] == '#' {
						occupiedNumber++
						break
					}
				}

				for k := 1; i - k >= 0 && j + k < len(result[i]); k++ {
					if temp[i - k][j + k] == 'L' {
						break
					}
					if temp[i - k][j + k] == '#' {
						occupiedNumber++
						break
					}
				}

				for k := 1; i + k < len(result) && j - k >= 0; k++ {
					if temp[i + k][j - k] == 'L' {
						break
					}
					if temp[i + k][j - k] == '#' {
						occupiedNumber++
						break
					}
				}

				if temp[i][j] == 'L' && occupiedNumber == 0 {
					s := []byte(result[i])
					s[j] = '#'
					result[i] = string(s)
				} else if temp[i][j] == '#' && occupiedNumber > 4 {
					s := []byte(result[i])
					s[j] = 'L'
					result[i] = string(s)
				}
			}
		}

		if sliceCompare(temp, result) {
			break
		}
	}

	occupiedSeats := 0
	for i := 0 ; i < len(result); i++ {
		for j := 0; j < len(result[i]); j++ {
			if result[i][j] == '#'{
				occupiedSeats++
			}
		}
	}

	return occupiedSeats
}

func sliceCompare(temp []string, result []string) bool {
	for i := 0 ; i < len(temp); i++ {
		for j := 0; j < len(temp[i]); j++ {
			if temp[i][j] != result[i][j] {
				return false
			}
		}
	}

	return true
}

func max(first int, second int) int {
	if first > second {
		return first
	} else {
		return second
	}
}

func min(first int, second int) int {
	if first < second {
		return first
	} else {
		return second
	}
}

func main(){
	answer := solution()
	fmt.Println(answer)
}
