package main

import (
	"bufio"
	"os"
	"fmt"
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
	for i := range line {
		if line[i] == 'S' {
			beamPos[i] = struct{}{}
			break
		}
	}
	splits := 0
	for scanner.Scan() {
		line = scanner.Text()
		for i, c := range line {
			if c == '^' {
				if _, ok := beamPos[i]; ok {
					splits++
					delete(beamPos, i)
					beamPos[i-1] = struct{}{}
					beamPos[i+1] = struct{}{}
				}
			}
		}
	}
	fmt.Println(splits)
}
