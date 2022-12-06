package main

import (
	"advc"
	"fmt"
)

func DayTwoPartTwo() {
	lines := advc.GetLines()

	// x - need to lose
	// y - need to draw
	// z - need to win

	results := map[string]int{
		"A X": 3,
		"A Y": 4,
		"A Z": 8,
		"B X": 1,
		"B Y": 5,
		"B Z": 9,
		"C X": 2,
		"C Y": 6,
		"C Z": 7,
	}

	var ans int

	for _, line := range lines {
		ans += results[line]
	}
	fmt.Println(ans)
}
