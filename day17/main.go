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

	store := make([][][][]int, 0)

	for m := 0; m < 30; m++ {
		highTemp := make([][][]int, 0)
		for i := 0; i < 30; i++ {
			temp := make([][]int, 0)
			for j := 0; j < 30; j++ {
				number := make([]int, 30)
				temp = append(temp, number)
			}
			highTemp = append(highTemp, temp)
		}
		store = append(store, highTemp)
	}


	for i := 0; i < len(result); i++ {
		for j := 0; j < len(result[i]); j++ {
			if result[i][j] == '#' {
				store[i + 10][j + 10][10][10] = 1
			}
		}
	}

	count := 0
	for count < 6 {
		newTemp := make([][][][]int, 0)
		//copy(newTemp, store)

		for m := 0; m < 30; m++ {
			highTemp := make([][][]int, 0)
			for i := 0; i < 30; i++ {
				temp := make([][]int, 0)
				for j := 0; j < 30; j++ {
					deepTemp := make([]int, 0)
					for k := 0; k < 30; k++ {
						deepTemp = append(deepTemp, store[m][i][j][k])
					}
					temp = append(temp, deepTemp)
				}
				highTemp = append(highTemp, temp)
			}
			newTemp = append(newTemp, highTemp)
		}

		for i := 3; i < 27; i++ {
			for j := 3; j < 27; j++ {
				for k := 3; k < 27; k++ {
					for n := 3; n < 27; n++ {
						number := 0
						for m := i - 1; m <= i + 1; m++ {
							for p := j - 1; p <= j + 1; p++ {
								for q := k - 1; q <= k + 1; q++ {
									for x := n - 1; x <= n + 1; x++ {
										if newTemp[m][p][q][x] == 1 {
											number++
										}
									}
								}
							}
						}
						if newTemp[i][j][k][n] == 1 && !(number == 3 || number == 4) {
							store[i][j][k][n] = 0
						} else if newTemp[i][j][k][n] == 0 && number == 3 {
							store[i][j][k][n] = 1
						}
					}
				}
			}
		}
		count++
	}

	countNumber := 0
	for i := 0; i < 30; i++ {
		for j := 0; j < 30; j++ {
			for k := 0; k < 30; k++ {
				for n := 0; n < 30; n++ {
					if store[i][j][k][n] == 1 {
						countNumber++
					}
				}
			}
		}
	}

	return countNumber
}

func main(){
	answer := solution()
	fmt.Println(answer)
}