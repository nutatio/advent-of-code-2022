package main

import (
	"advc"
	"fmt"
)

func DayTwoPartOne() {
	lines := advc.GetLines()
	// 1 for rock, 2 for paper, 3 for scissors
	// 0 for lose, 3 for draw, 6 for w

	results := map[string]int{
		"A X": 4,
		"A Y": 8,
		"A Z": 3,
		"B X": 1,
		"B Y": 5,
		"B Z": 9,
		"C X": 7,
		"C Y": 2,
		"C Z": 6,
	}

	var ans int

	for _, line := range lines {
		ans += results[line]
	}
	fmt.Println(ans)
}
