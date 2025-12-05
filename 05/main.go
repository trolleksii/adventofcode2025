package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Range struct {
	From, To int
}

func (r Range) size() int {
	return r.To - r.From + 1
}

func cmpRange(a, b Range) int {
    if a.From != b.From {
        return a.From - b.From
    }
    return a.To - b.To
}

func main() {
	file, err := os.Open("test.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	ranges := []Range{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		rs := strings.Split(line, "-")
		from, err := strconv.Atoi(rs[0])
		if err != nil {
			panic(err)
		}
		to, err := strconv.Atoi(rs[1])
		if err != nil {
			panic(err)
		}
		ranges = append(ranges, Range{From: from, To: to})
	}
	slices.SortFunc(ranges, cmpRange)

	joinedRanges := []Range{}
	current := ranges[0]
	for i := 1; i < len(ranges); i++ {
		next := ranges[i]
		if next.From <= current.To+1 {
			current = Range{From: current.From, To: max(current.To, next.To)}
		} else {
			joinedRanges = append(joinedRanges, current)
			current = next
		}
	}
	joinedRanges = append(joinedRanges, current)

	fresh := 0
	for _, x := range joinedRanges {
		fresh += x.size()
	}
	fmt.Println(fresh)
}
