package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.Open("test.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	puzzle := make(map[string][][]int)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")
		diagram := line[0][1:len(line[0])-1]
		//joltage := line[len(line)-1]
		buttons := [][]int{}
		for _, b := range line[1:len(line)-1] {
			combo := []int{}
			for _, i := range strings.Split(b[1:len(b)-1], ",") {
				num, _  := strconv.Atoi(i)
				combo = append(combo, num)
			}
			buttons = append(buttons, combo)
		}
		puzzle[diagram] = buttons
	}
	fmt.Println(puzzle)
}
