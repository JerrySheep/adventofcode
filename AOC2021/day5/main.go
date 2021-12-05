package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"strconv"
	"strings"
)

type Point struct {
	xPosition int
	yPosition int
}

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

	startPosition := make([]Point, 0)
	endPosition := make([]Point, 0)

	for i := 0; i < len(result); i++ {
		sentence := strings.Split(result[i], " -> ")
		start := strings.Split(sentence[0], ",")
		end := strings.Split(sentence[1], ",")

		temp := make([]int, 0)
		for j := 0; j < len(start); j++ {
			num, _ := strconv.Atoi(start[j])
			temp = append(temp, num)
		}
		startPosition = append(startPosition, Point{temp[0], temp[1]})
		temp = make([]int, 0)
		for j := 0; j < len(end); j++ {
			num, _ := strconv.Atoi(end[j])
			temp = append(temp, num)
		}
		endPosition = append(endPosition, Point{temp[0], temp[1]})
	}

	Points := make(map[Point]int)

	for i := 0; i < len(startPosition); i++ {
		xChange := endPosition[i].xPosition - startPosition[i].xPosition
		yChange := endPosition[i].yPosition - startPosition[i].yPosition


		//for xChange == 0 || yChange == 0{
		for {
			Points[Point{startPosition[i].xPosition, startPosition[i].yPosition}] += 1
			if startPosition[i].xPosition == endPosition[i].xPosition && startPosition[i].yPosition == endPosition[i].yPosition {
				break
			}

			if xChange != 0 {
				startPosition[i].xPosition += xChange / int(math.Abs(float64(xChange)))
			}
			if yChange != 0 {
				startPosition[i].yPosition += yChange / int(math.Abs(float64(yChange)))
			}
		}
	}

	count := 0
	for _, value := range Points {
		if value > 1 {
			count++
		}
	}

	return count
}

func main(){
	answer := solution()
	fmt.Println(answer)
}