package main

import (
	"fmt"
	"os"
	"bufio"
	"strconv"
)

func getMaxJoltage(bank string) int {
	cells := []byte{'0', '0'}
	l := len(bank)
	for i := 0; i < l-1; i++ {
		if bank[i] > cells[0] {
			cells[0] = bank[i]
			cells[1] = 0
		} else if bank[i] > cells[1] {
			cells[1] = bank[i]
		}
	}
	if bank[l-1] > cells[1] {
		cells[1] = bank[l-1]
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
