package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const (
	ROTATE_LEFT  = 'L'
	ROTATE_RIGHT = 'R'
	INSTRUCTIONS = "input.txt"
)

func rotateRight(currentNum int, distance int) (int, int) {
	newNum := currentNum + distance
	currentNum = newNum % 100
	clicks := newNum / 100
	return currentNum, clicks
}

func rotateLeft(currentNum int, distance int) (int, int) {
	clicks := 0
	newNum := currentNum - distance
	if currentNum > 0 && newNum <= 0 {
		clicks += 1
	}
	currentNum = newNum % 100
	clicks -= newNum / 100
	if currentNum < 0 {
		currentNum += 100
	}
	return currentNum, clicks
}

func main() {
	password := 0
	currentNum := 50
	clicks := 0
	file, err := os.Open(INSTRUCTIONS)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		op := scanner.Text()
		direction := op[0]
		distance, err := strconv.Atoi(op[1:])
		if err != nil {
			panic("non numeric distance")
		}

		switch direction {
		case ROTATE_LEFT:
			currentNum, clicks = rotateLeft(currentNum, distance)
		case ROTATE_RIGHT:
			currentNum, clicks = rotateRight(currentNum, distance)
		default:
			panic("unknown direction")
		}
		password += clicks
	}
	fmt.Println(password)
}
