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

		for _, line := range strings.Split(string(data), "\n"){
			if line != "" {
				result = append(result, line)
			}
		}
	}

	roadOfEastOrWest := 0
	roadOfSouthOrNorth := 0

	firstDirection := 0
	secondDirection := 3

	waypointOfEastOrWest := 10
	waypointOfSouthOrNorth := -1

	for i := 0; i < len(result); i++ {
		operation := result[i][0]
		number, ok := strconv.Atoi(result[i][1:])
		if ok != nil {
			fmt.Fprintf(os.Stderr, "%v\n", ok)
		}

		if operation == 'F' {
			roadOfEastOrWest += waypointOfEastOrWest * number
			roadOfSouthOrNorth += waypointOfSouthOrNorth * number
		} else if operation == 'R' || operation == 'L'{
			if operation == 'L' {
				number = 360 - number
			}
			formerEastWestNumber := waypointOfEastOrWest
			formerSouthNorthNumber := waypointOfSouthOrNorth
			direction := (firstDirection + (number / 90)) % 4
			if direction == 0 && formerEastWestNumber < 0 {
				waypointOfEastOrWest = 0 - formerEastWestNumber
			} else if direction == 2 && formerEastWestNumber > 0 {
				waypointOfEastOrWest = 0 - formerEastWestNumber
			} else if direction == 1 && formerEastWestNumber < 0 {
				waypointOfSouthOrNorth = 0 - formerEastWestNumber
			} else if direction == 1 {
				waypointOfSouthOrNorth = formerEastWestNumber
			} else if formerEastWestNumber > 0 {
				waypointOfSouthOrNorth = 0 - formerEastWestNumber
			} else {
				waypointOfSouthOrNorth = formerEastWestNumber
			}

			direction = (secondDirection + (number / 90)) % 4
			if direction == 0 && formerSouthNorthNumber < 0 {
				waypointOfEastOrWest = 0 - formerSouthNorthNumber
			} else if direction == 0 {
					waypointOfEastOrWest = formerSouthNorthNumber
			} else if direction == 2 && formerSouthNorthNumber > 0 {
					waypointOfEastOrWest = 0 - formerSouthNorthNumber
			} else if direction == 2 {
					waypointOfEastOrWest = formerSouthNorthNumber
			} else if direction == 1 && formerSouthNorthNumber < 0 {
				waypointOfSouthOrNorth = 0 - formerSouthNorthNumber
			} else if formerSouthNorthNumber > 0 {
				waypointOfSouthOrNorth = 0 - formerSouthNorthNumber
			}
		} else if operation == 'S' {
			waypointOfSouthOrNorth += number
		} else if operation == 'N' {
			waypointOfSouthOrNorth -= number
		} else if operation == 'E' {
			waypointOfEastOrWest += number
		} else if operation == 'W'{
			waypointOfEastOrWest -= number
		}

		if waypointOfEastOrWest >= 0 {
			firstDirection = 0
		} else {
			firstDirection = 2
		}

		if waypointOfSouthOrNorth >= 0 {
			secondDirection = 1
		} else {
			secondDirection = 3
		}
	}

	if roadOfEastOrWest < 0 {
		roadOfEastOrWest = 0 - roadOfEastOrWest
	}
	if roadOfSouthOrNorth < 0 {
		roadOfSouthOrNorth = 0 - roadOfSouthOrNorth
	}

	return roadOfSouthOrNorth + roadOfEastOrWest
}

func main(){
	answer := solution()
	fmt.Println(answer)
}
