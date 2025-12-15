package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
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
	visited := make(map[string]bool)
	required := map[string]bool{"fft": true, "dac": true}

	memo := make(map[State]int)
	fmt.Println(countPathsWithRequired(graph, "svr", visited, required, memo))
}

type State struct {
    node string
    reqKey string
}

func countPathsWithRequired(graph map[string][]string, current string, visited map[string]bool, required map[string]bool, memo map[State]int) int {
 	reqNodes := make([]string, 0, len(required))
    for k := range required {
        reqNodes = append(reqNodes, k)
    }
    slices.Sort(reqNodes)
    state := State{current, strings.Join(reqNodes, ",")}
    
    if val, ok := memo[state]; ok {
        return val
    }
    
    if current == "out" {
        if len(required) == 0 {
            return 1
        }
        return 0
    }

    visited[current] = true
    wasRequired := required[current]
    if wasRequired {
        delete(required, current)
    }

    count := 0
    for _, neighbor := range graph[current] {
        if !visited[neighbor] {
            count += countPathsWithRequired(graph, neighbor, visited, required, memo)
        }
    }

    visited[current] = false
    if wasRequired {
        required[current] = true
    }
    
    memo[state] = count
    return count
}
