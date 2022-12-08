package main

import (
	"advc"
	"fmt"
)

func isUnique(s string) bool {
	used := make(map[rune]struct{})
	for _, ch := range s {
		if _, ok := used[ch]; ok {
			return false
		}
		used[ch] = struct{}{}
	}
	return true
}

func partOne() int {
	var ans int
	datastream := advc.GetLines()[0]
	for i := 0; i < len(datastream)-3; i++ {
		if isUnique(datastream[i : i+4]) {
			ans = i + 4
			break
		}
	}

	return ans
}

func partTwo() int {
	var ans int
	datastream := advc.GetLines()[0]
	for i := 0; i < len(datastream)-13; i++ {
		if isUnique(datastream[i : i+14]) {
			ans = i + 14
			break
		}
	}

	return ans
}

func main() {
	fmt.Println("part one answear:", partOne())
	fmt.Println("part two answear:", partTwo())
}
