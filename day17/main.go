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

	matrix := make([]Coordinate, 0)
	for i := 0; i < len(result); i++ {
		for j := 0; j < len(result[i]); j++ {
			if result[i][j] == '#' {
				coordinate := Coordinate{i, j, 0, 0}
				matrix = append(matrix, coordinate)
			}
		}
	}

	count := 0
	for count < 6 {
		neighbors := make(map[Coordinate]int, 0)

		for i := 0; i < len(matrix); i++ {
			for j := 0; j < len(coordinates); j++ {
				temp := Coordinate{matrix[i].x + coordinates[j].x, matrix[i].y + coordinates[j].y, matrix[i].z + coordinates[j].z, matrix[i].w + coordinates[j].w}
				neighbors[temp] += 1
			}
		}

		length := len(matrix)
		for key, value := range neighbors {
			if value == 3 {
				matrix = append(matrix, key)
			} else if value == 2 {
				judge := false
				for i := 0; i < length; i++ {
					if key == matrix[i] {
						judge = true
						break
					}
				}
				if judge {
					matrix = append(matrix, key)
				}
			}
		}

		matrix = matrix[length:]
		count++
	}

	return len(matrix)
}

func main(){
	answer := solution()
	fmt.Println(answer)
}