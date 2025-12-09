package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Coordinate struct {
	x, y int
}

type Pair struct {
	c1, c2 Coordinate
	area   float64
}

func cmpCoord(a, b Coordinate) int {
    if a.y != b.y {
        return a.y - b.y
    }
    return a.x - b.x
}

func main() {
	file, err := os.Open("test.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	redTiles := []Coordinate{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), ",")
		x, err := strconv.Atoi(line[0])
		if err != nil {
			panic(err)
		}
		y, err := strconv.Atoi(line[1])
		if err != nil {
			panic(err)
		}
		redTiles = append(redTiles, Coordinate{x: x, y: y})
	}
	slices.SortFunc(redTiles, cmpCoord)
	fmt.Println(redTiles)
}
