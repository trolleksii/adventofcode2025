package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Coordinate struct {
	x, y int
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
		redTiles = append(redTiles, Coordinate{x: x, y:y})
	}
	fmt.Println(redTiles)
}
