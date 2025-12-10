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
	x, y int
}

func squareArea(a, b Coordinate) int {
	w := math.Abs(float64(a.x-b.x)) + 1
	h := math.Abs(float64(a.y-b.y)) + 1
	return int(w * h)
}

func getCoordinates(line string) Coordinate {
	p := strings.Split(line, ",")
	x, err := strconv.Atoi(p[0])
	if err != nil {
		panic(err)
	}
	y, err := strconv.Atoi(p[1])
	if err != nil {
		panic(err)
	}
	return Coordinate{x: x, y: y}
}

func main() {
	file, err := os.Open("test.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	redTiles := []Coordinate{}
	reds := make(map[Coordinate]bool)
	greens := make(map[Coordinate]bool)
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	this := getCoordinates(scanner.Text())
	redTiles = append(redTiles, this)
	reds[this] = true
	minX, maxX := this.x, this.x
	minY, maxY := this.y, this.y
	for scanner.Scan() {
		next := getCoordinates(scanner.Text())
		redTiles = append(redTiles, next)
		if next.x < minX {
			minX = next.x
		}
		if next.x > maxX {
			maxX = next.x
		}
		if next.y < minY {
			minY = next.y
		}
		if next.y > maxY {
			maxY = next.y
		}
		reds[next] = true
		// paint the vertical borders
		if next.x == this.x {
			if next.y > this.y {
				for i := this.y + 1; i < next.y; i++ {
					greens[Coordinate{x: this.x, y: i}] = true
				}
			} else {
				for i := next.y + 1; i < this.y; i++ {
					greens[Coordinate{x: this.x, y: i}] = true
				}
			}
		}
		this = next
	}
	// connect last point with the first point
	first := redTiles[0]
	if first.x == this.x {
		if first.y > this.y {
			for i := this.y + 1; i < first.y; i++ {
				greens[Coordinate{x: this.x, y: i}] = true
			}
		} else {
			for i := first.y + 1; i < this.y; i++ {
				greens[Coordinate{x: this.x, y: i}] = true
			}
		}
	}
	for y := minY; y <= maxY; y++ {
		inside := false
		for x := minX; x <= maxX; x++ {
			c := Coordinate{x: x, y: y}

			if greens[c] {
				// It's an edge - but is it a vertical edge?
				// Check if it connects tiles above AND below
				above := Coordinate{x: x, y: y - 1}
				below := Coordinate{x: x, y: y + 1}
				if (greens[above] || reds[above]) && (greens[below] || reds[below]) {
					inside = !inside
				}
				continue
			}

			if reds[c] {
				// Same check for red tiles
				above := Coordinate{x: x, y: y - 1}
				below := Coordinate{x: x, y: y + 1}
				if (greens[above] || reds[above]) && (greens[below] || reds[below]) {
					inside = !inside
				}
				continue
			}

			if inside {
				greens[c] = true
			}
		}
	}

	for y := range 9 {
		row := make([]byte, 14)
		for x := range 14 {
			c := Coordinate{x: x, y: y}
			if reds[c] {
				row[x] = '#'
			} else if greens[c] {
				row[x] = 'X'
			} else {
				row[x] = '.'
			}
		}
		fmt.Println(string(row))
	}
}
