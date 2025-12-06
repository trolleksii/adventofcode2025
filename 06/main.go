package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	lines := []string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}
	numericRows := len(lines) - 1
	opRow := lines[numericRows]
	var op rune
	result := 0
	subResult := 0
	for i, c := range opRow {
		switch c {
		case '+':
			op = c
			result += subResult
			subResult = 0
			fmt.Println("Result: ", result)
		case '*':
			op = c
			result += subResult
			subResult = 1
			fmt.Println("Result: ", result)
		}
		if i == len(opRow)-1 || opRow[i+1] == ' ' {
			s := ""
			for j := range numericRows {
				s += string(lines[j][i])
			}
			num, err := strconv.Atoi(strings.Trim(s, " "))
			if err != nil {
				panic(err)
			}
			switch op {
			case '+':
				subResult += num
			case '*':
				subResult *= num
			}
			fmt.Println(string(op), num, ", subresult: ", subResult)
		}
		if i == len(opRow)-1 {
			result += subResult
		}
	}
	fmt.Println(result)
}
