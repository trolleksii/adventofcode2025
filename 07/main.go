package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("test.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	line := scanner.Text()
	beamPos := make(map[int]struct{})
	parentBeams := make(map[int]string)
	for i := range line {
		if line[i] == 'S' {
			beamPos[i] = struct{}{}
			parentBeams[i] = "S"
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
					thisParent := fmt.Sprintf("%v-%v", parentBeams[i], i)

					if leftParent, exists := parentBeams[i-1]; exists {
						parentBeams[i-1] = strings.Join([]string{leftParent, thisParent}, ",")
					} else {
						parentBeams[i-1] = thisParent
					}
					if rightParent, exists := parentBeams[i+1]; exists {
						parentBeams[i+1] = strings.Join([]string{rightParent, thisParent}, ",")
					} else {
						parentBeams[i+1] = thisParent
					}
					delete(beamPos, i)
					delete(parentBeams, i)
				}
			}
		}
	}
	total := 0
	for i := range beamPos {
		total +=len(strings.Split(parentBeams[i], ","))
	}
	fmt.Println(total)
}
