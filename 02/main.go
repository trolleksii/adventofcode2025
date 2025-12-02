package main

import (
	"fmt"
	"strconv"
	"strings"
)

func isValid(num int) bool {
	r := strconv.Itoa(num)
	l := len(r)
	if l % 2 != 0 {
		return true
	}
	return r[:l / 2] != r[l / 2:]
}

func sumInvalidIDs(from, to int) int {
	result := 0
	for num := from; num <= to; num++ {
		if !isValid(num) {
			result += num
		}
	}
	return result
}

func main() {
	//testInput := "11-22,95-115,998-1012,1188511880-1188511890,222220-222224,1698522-1698528,446443-446449,38593856-38593862,565653-565659,824824821-824824827,2121212118-2121212124"
	input := "655-1102,2949-4331,885300-1098691,1867-2844,20-43,4382100-4484893,781681037-781860439,647601-734894,2-16,180-238,195135887-195258082,47-64,4392-6414,6470-10044,345-600,5353503564-5353567532,124142-198665,1151882036-1151931750,6666551471-6666743820,207368-302426,5457772-5654349,72969293-73018196,71-109,46428150-46507525,15955-26536,65620-107801,1255-1813,427058-455196,333968-391876,482446-514820,45504-61820,36235767-36468253,23249929-23312800,5210718-5346163,648632326-648673051,116-173,752508-837824"
	result := 0
	for r := range strings.SplitSeq(input, ",") {
		mm := strings.Split(r, "-")
		from, err := strconv.Atoi(mm[0])
		if err != nil {
			panic("non-numeric input")
		}
		to, err := strconv.Atoi(mm[1])
		if err != nil {
			panic("non-numeric input")
		}
		result += sumInvalidIDs(from, to)
	}
	fmt.Println(result)
}
