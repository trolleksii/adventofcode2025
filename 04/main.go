package main

import (
	"bufio"
	"fmt"
	"os"
)

const GRID_SIZE = 137

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var grid [GRID_SIZE][GRID_SIZE]bool
	var neighbours [GRID_SIZE + 2][GRID_SIZE + 2]int

	scanner := bufio.NewScanner(file)
	i := 1
	for scanner.Scan() {
		j := 1
		line := scanner.Text()
		for _, cell := range line {
			if cell == '@' {
				grid[i-1][j-1] = true
				neighbours[i-1][j-1]++
				neighbours[i-1][j]++
				neighbours[i-1][j+1]++
				neighbours[i][j-1]++
				neighbours[i][j+1]++
				neighbours[i+1][j-1]++
				neighbours[i+1][j]++
				neighbours[i+1][j+1]++
			} else {
				grid[i-1][j-1] = false
			}
			j++
		}
		i++
	}
	answer := 0
	i = 1
	for i < GRID_SIZE+1 {
		rollRemoved := false
		for j := 1; j < GRID_SIZE+1; j++ {
			if grid[i-1][j-1] && neighbours[i][j] < 4 {
				answer++
				rollRemoved = true
				grid[i-1][j-1] = false
				neighbours[i-1][j-1]--
				neighbours[i-1][j]--
				neighbours[i-1][j+1]--
				neighbours[i][j-1]--
				neighbours[i][j+1]--
				neighbours[i+1][j-1]--
				neighbours[i+1][j]--
				neighbours[i+1][j+1]--
			}
		}
		if rollRemoved {
			i = max(1, i-1) // if a roll was removed go back 1 row if possible
		} else {
			i++
		}
	}
	fmt.Println(answer)
}
