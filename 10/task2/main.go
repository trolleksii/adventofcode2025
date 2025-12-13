package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Button []int

func (b Button) connectLight(num int) {
	b[num] = 1
}

func NewButton(size int) Button {
	return make([]int, size)
}

type Puzzle struct {
	joltageReq []int
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
		joltageReqs := make([]int, len(joltages))
		for i, item := range joltages {
			joltageReqs[i], _ = strconv.Atoi(item)
		}

		buttons := []Button{}
		for _, bs := range data[1 : len(data)-1] {
			lights := strings.Split(bs[1:len(bs)-1], ",")
			button := NewButton(len(joltageReqs))
			for _, i := range lights {
				lightIndex, _ := strconv.Atoi(i)
				button.connectLight(lightIndex)
			}
			buttons = append(buttons, button)
		}
		puzzles = append(puzzles, Puzzle{joltageReq: joltageReqs, buttons: buttons})
	}
	total := 0
	for _, puzzle := range puzzles {
		// like recursion but iterative
		result := getShortestSequence(puzzle)
		total += result
	}
	fmt.Println(total)
}

// subtracts b from a, k times. Returns new slice and a "net balance" which is 0 if resulting slice has all 0, +1 if resulting slice has no negative numbers, -1 if has any negative numbers
func sub(a, b []int, k int) ([]int, int) {
	isEmpty := true
	hasNegatives := false
	c := make([]int, len(a))
	for i := range a {
		c[i] = a[i] - b[i]*k
		if c[i] != 0 {
			isEmpty = false
		}
		if c[i] < 0 {
			hasNegatives = true
		}
	}
	netBalance := 1
	if hasNegatives {
		netBalance = -1
	} else if isEmpty {
		netBalance = 0
	}
	return c, netBalance
}

func getMaxPresses(button, joltage []int) int {
	maxPresses := 999
	for i, v := range button {
		if v != 0 {
			maxPresses = min(maxPresses, joltage[i] / v)
		}
	}
	return maxPresses
}

func getShortestSequence(puzzle Puzzle) int {
	queue := &PriorityQueue{puzzle}
	heap.Init(queue)
	for queue.Len() > 0 {
		subPuzzle := heap.Pop(queue).(Puzzle)
		for buttonIndex, button := range subPuzzle.buttons {
			_, balance := sub(subPuzzle.joltageReq, button, 1)
			if balance < 0 {
				continue
			} else if balance == 0 {
				return subPuzzle.seqLen + 1
			}
			for i := getMaxPresses(button, subPuzzle.joltageReq); i >= 0; i-- {
				newJoltageReq, _ := sub(subPuzzle.joltageReq, button, i)
				newSubPuzzle := Puzzle{seqLen: subPuzzle.seqLen + i, joltageReq: newJoltageReq, buttons: slices.Concat(subPuzzle.buttons[:buttonIndex], subPuzzle.buttons[buttonIndex+1:])}
				heap.Push(queue, newSubPuzzle)
			}
		}
	}
	return -1
}
