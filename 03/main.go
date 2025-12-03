package main

import (
	"bufio"
	"strconv"
	"fmt"
	"os"
)

func getMaxJoltage(bank string) int {
	removals := len(bank) - 12
	cells := bank[removals:]
	j := 0
	for i := removals - 1; i >= 0; i-- {
		if bank[i] < cells[0] {
			continue
		}
		for j <11 {
			if cells[j] < cells[j+1] {
				break
			}
			j++
		} 
		cells = bank[i:i+1] + cells[0:j] + cells[j+1:]
	}
	res, err := strconv.Atoi(string(cells))
	if err != nil {
		panic("could not convert cell joltage to number")
	}
	return res
}

func main() {
	total := 0
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		bank := scanner.Text()
		total += getMaxJoltage(bank)
	}
	fmt.Println(total)
}
