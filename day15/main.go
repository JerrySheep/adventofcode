package main

import "fmt"

func solution() int {
	result := []int{0,14,1,3,7,9}
	count := len(result)

	keyValue := make(map[int]int)
	for i := 0; i < len(result); i++ {
		keyValue[result[i]] = i
	}

	for true {
		if count >= 30000000 {
			break
		}
		temp := result[len(result) - 1]
		value, ok := keyValue[temp]

		if !ok {
			result = append(result, 0)
		} else {
			result = append(result, len(result) - 1 - value)
		}
		keyValue[temp] = len(result) - 2

		count++
	}
	return result[len(result) - 1]
}

func main(){
	answer := solution()
	fmt.Println(answer)
}
