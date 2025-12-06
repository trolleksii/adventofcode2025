package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func add(arr [][]int, i int) int {
	res := 0
	for j := range arr {
		res += arr[j][i]
	}
	return res
}

func mul(arr [][]int, i int) int {
	res := 1
	for j := range arr {
		res *= arr[j][i]
	}
	return res
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	elements := [][]int{}
	scanner := bufio.NewScanner(file)
	result := 0
	for scanner.Scan() {
		line := strings.Fields(scanner.Text())
		row := []int{}
		for i, e := range line {
			switch e {
			case "+":
				result += add(elements, i)
			case "*":
				result += mul(elements, i)
			default:
				num, err := strconv.Atoi(e)
				if err != nil {
					panic(err)
				}
				row = append(row, num)
			}
		}
		elements = append(elements, row)
	}
	fmt.Println(result)
}
