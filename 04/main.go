package main

import (
	"bufio"
	"fmt"
	"os"
)

const GRID_SIZE = 10

func main() {
	file, err := os.Open("test.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var grid [GRID_SIZE + 2][GRID_SIZE + 2]int

	scanner := bufio.NewScanner(file)
	i := 1
	for scanner.Scan() {
		j := 1
		line := scanner.Text()
		for _, cell := range line {
			if cell == '@' {
				grid[i-1][j-1]++
				grid[i-1][j]++
				grid[i-1][j+1]++
				grid[i][j-1]++
				grid[i][j+1]++
				grid[i+1][j-1]++
				grid[i+1][j]++
				grid[i+1][j+1]++
			} else {
				grid[i][j] += 99
			}
			j++
		}
		i++
	}
	answer := 0
	for i := 1; i < GRID_SIZE+1; i++ {
		for j := 1; j < GRID_SIZE+1; j++ {
			if grid[i][j] < 4 {
				answer++
			}
		}
	}
	fmt.Println(answer)
}
