package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	line := scanner.Text()
	beamPos := make(map[int]struct{})
	waysToReach := make(map[int]int)
	for i := range line {
		if line[i] == 'S' {
			beamPos[i] = struct{}{}
			waysToReach[i] = 1
			break
		}
	}

	for scanner.Scan() {
		line = scanner.Text()
		for i, c := range line {
			if c == '^' {
				if _, ok := beamPos[i]; ok {
					beamPos[i-1] = struct{}{}
					beamPos[i+1] = struct{}{}
					waysToReach[i-1] = waysToReach[i-1] + waysToReach[i]
					waysToReach[i+1] = waysToReach[i+1] + waysToReach[i]
					delete(beamPos, i)
					delete(waysToReach, i)
				}
			}
		}
	}
	total := 0
	for _, w := range waysToReach {
		total +=w
	}
	fmt.Println(total)
}
