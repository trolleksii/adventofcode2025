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

func Parse(filename string) []Puzzle {
	file, _ := os.Open(filename)
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
			button := make(Button, len(joltageReqs))
			for _, i := range lights {
				lightIndex, _ := strconv.Atoi(i)
				button[lightIndex] = 1
			}
			buttons = append(buttons, button)
		}
		puzzles = append(puzzles, Puzzle{joltageReq: joltageReqs, buttons: buttons})
	}
	return puzzles
}

func main() {
	puzzles := Parse("../input.txt")
	for _, p := range puzzles {
		fmt.Println(p)
	}
	total := 0
	out := make(chan int)
	defer close(out)
	for _, puzzle := range puzzles {
		result := getShortestSequence(puzzle)
		fmt.Println(result)
		total += result
	}
	fmt.Println("Total:", total)
}

func sub(joltage, button []int, k int) []int {
	c := make([]int, len(joltage))
	for i := range joltage {
		c[i] = joltage[i] - button[i]*k
	}
	return c
}

func isZero(joltage []int) bool {
	for _, v := range joltage {
		if v != 0 {
			return false
		}
	}
	return true
}

func getMaxPresses(joltage, button []int) int {
	// to find the max number of times a button can be pressed considering current joltage req
	// we go through all the bits button affects and find the lowest joltage
	maxP := 999
	for _, v := range button {
		if v != 0 {
			maxP = min(maxP, joltage[v])
		}
	}
	return maxP
}

func sliceToKey(j []int, bt []Button) string {
	b := strings.Builder{}
	for i, v := range j {
		if i > 0 {
			b.WriteByte(',')
		}
		b.WriteString(strconv.Itoa(v))
	}
	b.WriteByte('+')
	for i, bb := range bt {
		if i > 0 {
			b.WriteByte('|')
		}
		for j, bbt := range bb {
			if j > 0 {
				b.WriteByte(',')
			}
			b.WriteString(strconv.Itoa(bbt))
		}
	}
	return b.String()
}

func getShortestSequence(puzzle Puzzle) int {
	queue := &PriorityQueue{puzzle}
	heap.Init(queue)
	visited := make(map[string]bool)
	for queue.Len() > 0 {
		subPuzzle := heap.Pop(queue).(Puzzle)
		key := sliceToKey(subPuzzle.joltageReq, subPuzzle.buttons)
		if visited[key] {
			continue
		}
		visited[key] = true
		for buttonIndex, button := range subPuzzle.buttons {
			if isZero(subPuzzle.joltageReq) {
				return subPuzzle.seqLen
			}
			for i := 0; i <= getMaxPresses(subPuzzle.joltageReq, button); i++ {
				newJoltageReq := sub(subPuzzle.joltageReq, button, i)
				newSubPuzzle := Puzzle{seqLen: subPuzzle.seqLen + i, joltageReq: newJoltageReq, buttons: slices.Concat(subPuzzle.buttons[:buttonIndex], subPuzzle.buttons[buttonIndex+1:])}
				heap.Push(queue, newSubPuzzle)
			}
		}
	}
	return -1
}
