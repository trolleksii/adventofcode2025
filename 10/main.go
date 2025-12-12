package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Puzzle struct {
	target     int
	components []int
	seqLen     int
}

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()

	puzzles := []Puzzle{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")
		target := 0
		diagram := line[0][1 : len(line[0])-1]
		// traversing in the opposite direction to compensate the inverse
		for i := len(diagram) - 1; i >= 0; i-- {
			switch diagram[i] {
			case '#':
				target = (target << 1) + 1
			case '.':
				target = (target << 1)
			}
		}
		components := []int{}
		// knowing which bits/buttons are in the combo, we can do bit shifting and sum to get the binary representation
		// I couldn't figure how to do in proper order so I reversed the diagram instead and that
		for _, b := range line[1 : len(line)-1] {
			inverse := 0
			for i := range strings.SplitSeq(b[1:len(b)-1], ",") {
				num, _ := strconv.Atoi(i)
				inverse += 1 << num
			}
			components = append(components, inverse)
		}
		puzzles = append(puzzles, Puzzle{target: target, components: components, seqLen: 1})
	}
	total := 0
	for _, puzzle := range puzzles {
		// like recursion but iterative
		total += getShortestSequence(puzzle)
	}
	fmt.Println(total)
}

func getShortestSequence(puzzle Puzzle) int {
	queue := []Puzzle{puzzle}
	// we go through all button combos and figure if pressing that yields the desired result using XOR
	// return with a first match, throw new subPuzleis which excludes already pressed buttons to the end of the queue
	for len(queue) > 0 {
		subPuzzle := queue[0]
		queue = queue[1:]
		// see if we can get to target in one press
		for i, comp := range subPuzzle.components {
			//skip button combos which don't touch interesting bits
			if comp&subPuzzle.target == 0 {
				continue
			}
			if comp == subPuzzle.target {
				return subPuzzle.seqLen
			} else {
				// put sub-puzzle in a queue with new target, longer sequence, smaller component list
				newTarget := subPuzzle.target ^ comp
				queue = append(queue, Puzzle{seqLen: subPuzzle.seqLen + 1, target: newTarget, components: slices.Concat(subPuzzle.components[:i], subPuzzle.components[i+1:])})
			}
		}
	}
	// assuming solution always exist this should never run
	return len(puzzle.components)
}
