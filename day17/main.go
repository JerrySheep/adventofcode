package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type Coordinate struct {
	x int
	y int
	z int
	w int
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

	coordinates := make([]Coordinate, 0)
	for x := -1; x <= 1; x++ {
		for y := -1; y <= 1; y++ {
			for z := -1; z <= 1; z++ {
				for w := -1; w <= 1; w++ {
					if x == 0 && y == 0 && z == 0 && w == 0 {
						continue
					}
					coordinate := Coordinate{x, y, z, w}
					coordinates = append(coordinates, coordinate)
				}
			}
		}
	}

	matrix := make(map[Coordinate]int, 0)
	for i := 0; i < len(result); i++ {
		for j := 0; j < len(result[i]); j++ {
			if result[i][j] == '#' {
				coordinate := Coordinate{i, j, 0, 0}
				matrix[coordinate] = 1
			}
		}
	}

	count := 0
	for count < 6 {
		neighbors := make(map[Coordinate]int, 0)

		for key, _ := range matrix {
			for j := 0; j < len(coordinates); j++ {
				temp := Coordinate{key.x + coordinates[j].x, key.y + coordinates[j].y, key.z + coordinates[j].z, key.w + coordinates[j].w}
				neighbors[temp] += 1
			}
		}

		newMatrix := make(map[Coordinate]int, 0)

		for key, value := range neighbors {
			_, ok := matrix[key]
			if value == 3 {
				newMatrix[key] = 1
			} else if value == 2 && ok {
				newMatrix[key] = 1
			}
		}
		matrix = newMatrix
		count++
	}

	return len(matrix)
}

func main(){
	answer := solution()
	fmt.Println(answer)
}

