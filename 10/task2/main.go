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

type Puzzle struct {
	joltageReq int
	seqLen     int
	buttons    []int
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
		line := strings.Split(scanner.Text(), " ")
		joltageReq := strings.ReplaceAll(line[len(line)-1], ",", "")
		joltageReq = joltageReq[1 : len(joltageReq)-1]
		joltageNum := 0
		for i, ch := range joltageReq {
			n, _ := strconv.Atoi(string(ch))
			joltageNum += n * int(math.Pow10(len(joltageReq)-i-1))
		}

		buttons := []int{}
		for _, b := range line[1 : len(line)-1] {
			button := 0
			for i := range strings.SplitSeq(b[1:len(b)-1], ",") {
				lightIndex, _ := strconv.Atoi(i)
				button += int(math.Pow10(len(joltageReq) - lightIndex - 1))
			}
			buttons = append(buttons, button)
		}
		slices.SortFunc(buttons, func(a, b int) int { return b - a })
		puzzles = append(puzzles, Puzzle{joltageReq: joltageNum, buttons: buttons})
	}
	total := 0
	for _, puzzle := range puzzles {
		// like recursion but iterative
		total += getShortestSequence(puzzle)
	}
	fmt.Println(total)
}

func getShortestSequence(puzzle Puzzle) int {
	return 0
}
