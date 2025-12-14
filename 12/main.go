package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Cell struct {
	row, col int
}

type Shape []Cell

func (s Shape) Hash() string {
	var sb strings.Builder
	for _, c := range s {
		sb.WriteString(fmt.Sprintf("%d,%d|", c.row, c.col))
	}
	return sb.String()
}

func (s Shape) Flip() Shape {
	norm := 0
	newShape := Shape{}
	for _, cell := range s {
		if cell.col > norm {
			norm = cell.col
		}
	}
	for _, cell := range s {
		flippedCell := Cell{row: cell.row, col: norm + (-1 * cell.col)}
		newShape = append(newShape, flippedCell)
	}
	slices.SortFunc(newShape, cmpCoordinates)
	return newShape
}

func (s Shape) Rotate() Shape {
	norm := 0
	newShape := Shape{}
	for _, cell := range s {
		if cell.row > norm {
			norm = cell.row
		}
	}

	for _, cell := range s {
		rotatedCell := Cell{row: cell.col, col: norm + (-1 * cell.row)}
		newShape = append(newShape, rotatedCell)
	}
	slices.SortFunc(newShape, cmpCoordinates)
	return newShape
}

func deduplicate(shapes []Shape) []Shape {
	set := make(map[string]Shape)
	for _, v := range shapes {
		set[v.Hash()] = v
	}
	deduped := []Shape{}
	for _, v := range set {
		deduped = append(deduped, v)
	}
	return deduped
}

func cmpCoordinates(a, b Cell) int {
	if a.row == b.row {
		return a.col - b.col
	}
	return a.row - b.row
}

type Region struct {
	width, height int
	pieces        [6]int
}

func parse(name string) (map[int][]Shape, []Region) {
	file, err := os.Open(name)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// parse 5 figures
	allShapes := make(map[int][]Shape)
	shapes := []Shape{}

	for range 6 {
		scanner.Scan() //skip first line
		shape := Shape{}
		for row := range 3 {
			scanner.Scan() // next 3 lines will be the shape
			for col, ch := range scanner.Text() {
				if ch == '#' {
					shape = append(shape, Cell{row: row, col: col})
				}
			}
		}
		shapes = append(shapes, shape)
		scanner.Scan() // final line is the empty line
	}
	// find permutations
	for k, v := range shapes {
		flip := v.Flip()
		permutations := []Shape{v, flip}
		for range 3 {
			v = v.Rotate()
			flip = flip.Rotate()
			permutations = append(permutations, v, flip)
		}
		allShapes[k] = deduplicate(permutations)
	}

	regions := []Region{}
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), ":")
		region := Region{}
		size := strings.Split(line[0], "x")
		region.width, _ = strconv.Atoi(size[0])
		region.height, _ = strconv.Atoi(size[1])
		for i, v := range strings.Fields(line[1]) {
			region.pieces[i], _ = strconv.Atoi(v)
		}
		regions = append(regions, region)
	}
	return allShapes, regions
}

func main() {
	shapes, regions := parse("test.txt")
	fmt.Println(shapes, regions)
}
