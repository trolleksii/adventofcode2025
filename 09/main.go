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

	reds := make(map[Coordinate]bool)
	greens := make(map[Coordinate]bool)
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	this := getCoordinates(scanner.Text())
	first := this // remember this to connect it with last point
	reds[this] = true
	minX, maxX := this.x, this.x
	minY, maxY := this.y, this.y
	for scanner.Scan() {
		next := getCoordinates(scanner.Text())
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
		// each consecutive pair has same x or y, everything in between is green
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
		} else {
			if next.x > this.x {
				for i := this.x + 1; i < next.x; i++ {
					greens[Coordinate{x: i, y: this.y}] = true
				}
			} else {
				for i := next.x + 1; i < this.x; i++ {
					greens[Coordinate{x: i, y: this.y}] = true
				}
			}
		}
		this = next
	}
	// connect last point with the first point
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
	} else {
		if first.x > this.x {
			for i := this.x + 1; i < first.x; i++ {
				greens[Coordinate{x: i, y: this.y}] = true
			}
		} else {
			for i := first.x + 1; i < this.x; i++ {
				greens[Coordinate{x: i, y: this.y}] = true
			}
		}
	}
	// colour everything inside the boundaries
	var floodFill func(Coordinate)
	floodFill = func(c Coordinate) {
		if reds[c] || greens[c] {
			return
		}
		greens[c] = true
		floodFill(Coordinate{c.x + 1, c.y})
		floodFill(Coordinate{c.x - 1, c.y})
		floodFill(Coordinate{c.x, c.y + 1})
		floodFill(Coordinate{c.x, c.y - 1})
	}
	floodFill(Coordinate{x: 8, y: 2})

	for y := range(9) {
		row := make([]byte, 14)
		for x := range 14 {
			c := Coordinate{x:x, y:y}
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
