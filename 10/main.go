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
	lightsDiagram int
	buttons       []int
	seqLen        int
}

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()

	puzzles := []Puzzle{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")
		lightsDiagram := 0
		diagramString := line[0][1 : len(line[0])-1]
		// traversing in the opposite direction to compensate the buttonInversed
		for i := len(diagramString) - 1; i >= 0; i-- {
			switch diagramString[i] {
			case '#':
				lightsDiagram = (lightsDiagram << 1) + 1
			case '.':
				lightsDiagram = (lightsDiagram << 1)
			}
		}

		buttons := []int{}
		// knowing which bits/lights are affected by the button, we can do bit shifting and sum to get the binary representation
		// I couldn't figure how to do in proper order so I reversed the diagramString instead
		for _, b := range line[1 : len(line)-1] {
			buttonInversed := 0
			for i := range strings.SplitSeq(b[1:len(b)-1], ",") {
				num, _ := strconv.Atoi(i)
				buttonInversed += 1 << num
			}
			buttons = append(buttons, buttonInversed)
		}
		puzzles = append(puzzles, Puzzle{lightsDiagram: lightsDiagram, buttons: buttons, seqLen: 1})
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
		// see if we can get to lightsDiagram in one press
		for i, comp := range subPuzzle.buttons {
			//skip button combos which don't touch interesting bits
			if comp&subPuzzle.lightsDiagram == 0 {
				continue
			}
			if comp == subPuzzle.lightsDiagram {
				return subPuzzle.seqLen
			} else {
				// put sub-puzzle in a queue with new lightsDiagram, longer sequence, smaller component list
				newTarget := subPuzzle.lightsDiagram ^ comp
				queue = append(queue, Puzzle{seqLen: subPuzzle.seqLen + 1, lightsDiagram: newTarget, buttons: slices.Concat(subPuzzle.buttons[:i], subPuzzle.buttons[i+1:])})
			}
		}
	}
	// assuming solution always exist the worst case scenario is pressing all buttons once
	return len(puzzle.buttons)
}
