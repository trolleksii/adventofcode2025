package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Button struct {
	value      int
	maxPresses int
}

type Puzzle struct {
	joltageReq int
	seqLen     int
	buttons    []Button
}

type PriorityQueue []Puzzle

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].seqLen < pq[j].seqLen
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x any) {
	*pq = append(*pq, x.(Puzzle))
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[:n-1]
	return item
}
func main() {
	file, _ := os.Open("../test.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// first step is to convert button/light data and joltage requirements into numeric representation
	// the entire thing would be an equation a1*b1 + a2*b2 + a3*b3 + ... an*bn = joltageTarget
	// where a1, a2, .. an is a number of presses on a given button one need to make to get to the target
	// it must be only natural numbers greater or equal than 0
	puzzles := []Puzzle{}
	for scanner.Scan() {
		data := strings.Split(scanner.Text(), " ")
		joltageData := data[len(data)-1]
		joltageData = joltageData[1 : len(joltageData)-1]
		joltages := strings.Split(joltageData, ",")
		joltageNum := 0
		maxLightPresses := []int{}
		for i, item := range joltages {
			n, _ := strconv.Atoi(item)
			maxLightPresses = append(maxLightPresses, n)
			joltageNum += n * int(math.Pow10(len(joltages)-i-1))
		}

		buttons := []Button{}
		for _, b := range data[1 : len(data)-1] {
			button := 0
			maxPresses := 999
			for i := range strings.SplitSeq(b[1:len(b)-1], ",") {
				lightIndex, _ := strconv.Atoi(i)
				maxPresses = min(maxPresses, maxLightPresses[lightIndex])
				button += int(math.Pow10(len(joltages) - lightIndex - 1))
			}
			buttons = append(buttons, Button{value: button, maxPresses: maxPresses})
		}
		slices.SortFunc(buttons, func(a, b Button) int { return b.value - a.value })
		puzzles = append(puzzles, Puzzle{joltageReq: joltageNum, buttons: buttons})
	}
	total := 0
	for _, puzzle := range puzzles {
		// like recursion but iterative
		result := getShortestSequence(puzzle)
		total += result
	}
	fmt.Println(total)
}

func getShortestSequence(puzzle Puzzle) int {
	queue := &PriorityQueue{puzzle}
	heap.Init(queue)
	for queue.Len() > 0 {
		subPuzzle := heap.Pop(queue).(Puzzle)
		for buttonIndex, button := range subPuzzle.buttons {
			if button.value > subPuzzle.joltageReq {
				continue
			} else if subPuzzle.joltageReq == button.value {
				return subPuzzle.seqLen + 1
			}
			for i := button.maxPresses; i >= 0; i-- {
				newJoltageReq := subPuzzle.joltageReq - button.value*i
				newSubPuzzle := Puzzle{seqLen: subPuzzle.seqLen + i, joltageReq: newJoltageReq, buttons: slices.Concat(subPuzzle.buttons[:buttonIndex], subPuzzle.buttons[buttonIndex+1:])}
				heap.Push(queue, newSubPuzzle)
			}
		}
	}
	return -1
}
