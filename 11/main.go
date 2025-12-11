package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	graph := make(map[string][]string)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), ":")
		name := line[0]
		neighbours := strings.Fields(line[1])
		graph[name] = neighbours
	}
	paths := 0
	queue := []string{"you"}
	visited := make(map[string]bool)
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		if current == "out" {
			paths++
		}
		visited[current] = true
		queue = append(queue, graph[current]...)
	}
	fmt.Println(paths)
}
