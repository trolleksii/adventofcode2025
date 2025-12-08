package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Coordinate struct {
	x, y, z float64
}

type Pair struct {
	c1, c2 Coordinate
	distance float64
}

func NewPair(c1, c2 Coordinate) Pair{
	d := distance(c1, c2)
	return Pair{c1:c1, c2:c2, distance: d}
}


func cmpPair(a, b Pair) int {
	if a.distance > b.distance {
		return 1
	} else if a.distance < b.distance {
		return -1
	}
	return 0
}

func distance(a, b Coordinate) float64 {
	return math.Sqrt((a.x-b.x)*(a.x-b.x) + (a.y-b.y)*(a.y-b.y) + (a.z-b.z)*(a.z-b.z))
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	coordinates := []Coordinate{}
	coordChain := make(map[Coordinate]int)
	chains := make(map[int]int)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		coordLine := strings.Split(scanner.Text(), ",")

		coord := Coordinate{}
		num, err := strconv.Atoi(coordLine[0])
		if err != nil {
			panic(err)
		}
		coord.x = float64(num)
		num, err = strconv.Atoi(coordLine[1])
		if err != nil {
			panic(err)
		}
		coord.y = float64(num)
		num, err = strconv.Atoi(coordLine[2])
		if err != nil {
			panic(err)
		}
		coord.z = float64(num)
		coordinates = append(coordinates, coord)
	}

	coordinatePairs := []Pair{}
	l := len(coordinates)
	for i := range coordinates {
		for j := i + 1; j < l; j++ {
			p := NewPair(coordinates[i], coordinates[j])
			coordinatePairs = append(coordinatePairs, p)
		}
	}

	slices.SortFunc(coordinatePairs, cmpPair)
	chainIndex := 0
	var lastPair Pair
	for _, p := range coordinatePairs {
		chainA, aIsInChain := coordChain[p.c1]
		chainB, bIsInChain := coordChain[p.c2]
		if !aIsInChain && !bIsInChain {
			// both points are unconnected
			chains[chainIndex] = 2
			coordChain[p.c1] = chainIndex
			coordChain[p.c2] = chainIndex
			chainIndex++
			lastPair = p
		} else if !aIsInChain {
			// if a is chained and b is not, get the intex of the chain a and include b into it
			chains[chainB]++
			coordChain[p.c1] = chainB
			lastPair = p
		} else if !bIsInChain {
			// if a is chained and b is not, get the intex of the chain a and include b into it
			chains[chainA]++
			coordChain[p.c2] = chainA
			lastPair = p
		} else if chainA != chainB {
			for k, v := range coordChain {
				if v == chainB {
					coordChain[k] = chainA
				}
			}
			chains[chainA] += chains[chainB]
			delete(chains, chainB)
			lastPair = p
		} 
	}
	fmt.Println(int(lastPair.c1.x * lastPair.c2.x))
}
