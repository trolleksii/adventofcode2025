package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Coordinate struct {
	x, y float64
}

type Pair struct {
	c1, c2 Coordinate
	area   float64
}

func NewPair(c1, c2 Coordinate) Pair {
	a := squareArea(c1, c2)
	return Pair{c1: c1, c2: c2, area: a}
}

func squareArea(a, b Coordinate) float64 {
	w := math.Abs(a.x - b.x) + 1
	h := math.Abs(a.y - b.y) + 1
	return w * h
}

func main() {
	file, err := os.Open("input.txt")
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
		redTiles = append(redTiles, Coordinate{x: float64(x), y: float64(y)})
	}

	coordinatePairs := []Pair{}
	var maxAreaPair Pair
	l := len(redTiles)
	for i := range redTiles {
		for j := i + 1; j < l; j++ {
			p := NewPair(redTiles[i], redTiles[j])
			if p.area > maxAreaPair.area {
				maxAreaPair = p
			}
			coordinatePairs = append(coordinatePairs, p)
		}
	}

	fmt.Println(int(maxAreaPair.area))
}
