//copy from https://github.com/lizthegrey/adventofcode/blob/main/2020/day19.go
package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type Rule struct {
	Literal string
	Children [][]int
}

type RuleSet map[int]Rule

func match(rules RuleSet, ruleNo int, s string) []int {
	rule := rules[ruleNo]
	if len(rule.Children) == 0 {
		if len(s) < len(rule.Literal) {
			return nil
		}
		if s[:len(rule.Literal)] == rule.Literal {
			return []int{len(rule.Literal)}
		}
	}

	matchedChars := make([]int, 0)
	for _, child := range rule.Children {
		potentialMatches := []int{0}
		for _, singleChild := range child {
			newPotentialMatches := make([]int, 0)
			for _, matchLength := range potentialMatches {
				matches := match(rules, singleChild, s[matchLength:len(s)])

				for _, value := range matches {
					newPotentialMatches = append(newPotentialMatches, value + matchLength)
				}
			}
			potentialMatches = newPotentialMatches
		}
		matchedChars = append(matchedChars, potentialMatches...)
	}

	return matchedChars
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

	result = result[:len(result) - 1]

	rules := make(RuleSet)
	messages := make([]string, 0)
	judge := false

	for _, s := range result {
		if s == "" {
			judge = true
			continue
		}
		if judge {
			messages = append(messages, s)
			continue
		}

		index := strings.Index(s, ":")
		ruleNo, err := strconv.Atoi(s[0:index])
		if err != nil {
			fmt.Printf("Failed to parse %s\n", s)
			break
		}

		rule := Rule{}
		if s[index + 2] == '"' {
			rule.Literal = s[index + 3 : len(s) - 1]
		} else {
			matches := strings.Split(s[index + 2 : len(s)], " | ")
			for _, m := range matches {
				elements := strings.Split(m, " ")
				group := make([]int, 0)
				for _, ele := range elements {
					child, err := strconv.Atoi(ele)
					if err != nil {
						fmt.Printf("Failed to parse %s\n", s)
						break
					}
					group = append(group, child)
				}
				rule.Children = append(rule.Children, group)
			}
		}
		rules[ruleNo] = rule
	}

	rules[8] = Rule{Children: [][]int{{42}, {42, 8}}}
	rules[11] = Rule{Children: [][]int{{42, 31}, {42, 11, 31}}}

	matches := 0
	for _, m := range messages {
		symbolMatches := match(rules, 0, m)
		for _, times := range symbolMatches {
			if times == len(m){
				matches++
				break
			}
		}
	}

	return matches
}

func main(){
	answer := solution()
	fmt.Println(answer)
}