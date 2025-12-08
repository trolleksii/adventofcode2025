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
	connections := 0
	for _, p := range coordinatePairs {
		fmt.Printf("Checking points %v and %v\n", p.c1, p.c2)
		chainA, aIsInChain := coordChain[p.c1]
		chainB, bIsInChain := coordChain[p.c2]
		if !aIsInChain && !bIsInChain {
			// both points are unconnected
			chains[chainIndex] = 2
			coordChain[p.c1] = chainIndex
			coordChain[p.c2] = chainIndex
			fmt.Printf("  neither point is chained, creating chain %v with size %v\n", chainIndex, 2)
			chainIndex++
		} else if !aIsInChain {
			// if a is chained and b is not, get the intex of the chain a and include b into it
			chains[chainB]++
			coordChain[p.c1] = chainB
			fmt.Printf("  point %v is unchained, attaching it to chain %v with new size %v\n", p.c1, chainB, chains[chainB])
		} else if !bIsInChain {
			// if a is chained and b is not, get the intex of the chain a and include b into it
			chains[chainA]++
			coordChain[p.c2] = chainA
			fmt.Printf("  point %v is unchained, attaching it to chain %v with new size %v\n", p.c2, chainA, chains[chainA])
		} else if chainA != chainB {
			for k, v := range coordChain {
				if v == chainB {
					coordChain[k] = chainA
				}
			}
			fmt.Printf("  point %v is in chain %v(size=%v), point %v is in chain %v(size=%v); merging as %v size %v\n", p.c1, chainA, chains[chainA], p.c2, chainB, chains[chainB], chainA, chains[chainA]+chains[chainB])
			chains[chainA] += chains[chainB]
			delete(chains, chainB)
		}
		connections++
		if connections == 1000 {
			break
		}
	}
	x := []int{}
	for _, v := range chains {
		x = append(x, v)
	}
	slices.SortFunc(x, func(a, b int) int { return b - a })
	fmt.Println("Top 3 chain sizes: ", x[:3])
}
