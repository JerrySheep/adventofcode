package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func answerNumberCal(s string, number int) int {
	count := 0
	store := make([]int, 26)
	for i := 0; i < len(s); i++ {
		store[s[i] - 'a'] += 1
	}

	for i := 0; i < 26; i++ {
		if store[i] == number {
			count += 1
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

		for _, line := range strings.Split(string(data), "\n") {
			result = append(result, line)
		}
	}

	store := make([]string, 0)
	personStore := make([]int, 0)

	for i := 0; i < len(result); i++ {
		personStore = append(personStore, 1)
		store = append(store, result[i])
		for result[i + 1] != "" && i + 1 < len(result){
			i++
			store[len(store) - 1] += result[i]
			personStore[len(personStore) - 1] += 1
		}
		i++
	}

	countNumber := 0
	for i := 0; i < len(store); i++ {
		countNumber += answerNumberCal(store[i], personStore[i])
	}

	return countNumber
}

func main(){
	answer := solution()
	fmt.Println(answer)
}
