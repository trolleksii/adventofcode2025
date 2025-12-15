package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Vertex struct {
	x, y float64
}

type Edge struct {
	fixed, from, to float64
}

func squareArea(a, b Vertex) int {
	w := math.Abs(a.x-b.x) + 1
	h := math.Abs(a.y-b.y) + 1
	return int(w * h)
}

func getCoordinates(line string) Vertex {
	p := strings.Split(line, ",")
	x, _ := strconv.Atoi(p[0])
	y, _ := strconv.Atoi(p[1])
	return Vertex{x: float64(x), y: float64(y)}
}

func Parse(filename string) ([]Vertex, []Edge, []Edge) {
	file, _ := os.Open(filename)
	defer file.Close()

	redTiles := []Vertex{}
	vedges := []Edge{}
	hedges := []Edge{}
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	this := getCoordinates(scanner.Text())
	redTiles = append(redTiles, this)
	for scanner.Scan() {
		next := getCoordinates(scanner.Text())
		redTiles = append(redTiles, next)
		if next.x == this.x {
			vedges = append(vedges, Edge{fixed: this.x, from: min(this.y, next.y), to: max(this.y, next.y)})
		} else {
			hedges = append(hedges, Edge{fixed: this.y, from: min(this.x, next.x), to: max(this.x, next.x)})
		}
		this = next
	}
	next := redTiles[0]
	if this.x == next.x {
		vedges = append(vedges, Edge{fixed: this.x, from: min(this.y, next.y), to: max(this.y, next.y)})
	} else {
		hedges = append(hedges, Edge{fixed: this.y, from: min(this.x, next.x), to: max(this.x, next.x)})
	}
	return redTiles, vedges, hedges
}

func noEdgesInside(minx, miny, maxx, maxy float64, vedges []Edge, hedges []Edge) bool {
	for _, e := range vedges {
		if e.fixed > minx && e.fixed < maxx && e.to > miny && e.from < maxy {
			return false
		}
	}

	for _, e := range hedges {
		if e.fixed > miny && e.fixed < maxy && e.to > minx && e.from < maxx {
			return false
		}
	}
	return true
}

func isInside(x, y float64, edges []Edge) bool {
	intersections := 0
	for _, e := range edges {
		if e.fixed > x {
			if e.from <= y && y < e.to {
				intersections++
			}
		}
	}
	return intersections%2 == 1
}

func main() {
	redTiles, vedges, hedges := Parse("input.txt")
	maxArea := 0
	for i, v1 := range redTiles {
		for j := i + 1; j < len(redTiles); j++ {
			v2 := redTiles[j]
			if area := squareArea(v1, v2); area < maxArea {
				continue
			} else {
				minx := min(v1.x, v2.x)
				miny := min(v1.y, v2.y)
				maxx := max(v1.x, v2.x)
				maxy := max(v1.y, v2.y)
				if isInside(minx+0.5, miny+0.5, vedges) {
					if noEdgesInside(minx, miny, maxx, maxy, vedges, hedges) {
						maxArea = area
					}
				}
			}
		}
	}
	fmt.Println(maxArea)
}
