package main

import (
	"advc"
	"fmt"
	"unicode"
)

type stack []rune

func (s *stack) Push(item rune) {
	*s = append(*s, item)
}

func (s *stack) Pop() (item rune) {
	if len(*s) == 0 {
		return
	}

	*s, item = (*s)[:len(*s)-1], (*s)[len(*s)-1]
	return item
}

// returns field and commands
func setup() ([]stack, []string) {
	field := make([]stack, 0)
	for i := 0; i < 9; i++ {
		s := make([]rune, 0)
		field = append(field, s)
	}
	lines := advc.GetLines()
	for i := 7; i >= 0; i-- {
		for j := 1; j < 35; j += 4 {
			s := rune(lines[i][j])
			if !unicode.IsSpace(s) {
				field[(j-1)/4].Push(s)
			}
		}
	}
	return field, lines[10:]
}

func partOne() string {
	field, lines := setup()
	for _, line := range lines {
		var amount, from, to int
		fmt.Sscanf(line, "move %d from %d to %d", &amount, &from, &to)
		from--
		to--
		for j := 0; j < amount; j++ {
			t := field[from].Pop()
			field[to].Push(t)
		}
	}

	ans := make([]rune, 0)
	for _, st := range field {
		ans = append(ans, st.Pop())
	}

	return string(ans)
}

func partTwo() string {
	field, lines := setup()
	for _, line := range lines {
		var amount, from, to int
		fmt.Sscanf(line, "move %d from %d to %d", &amount, &from, &to)
		from--
		to--
		tmp := stack{}

		for j := 0; j < amount; j++ {
			t := field[from].Pop()
			tmp.Push(t)

		}
		for j := 0; j < amount; j++ {
			t := tmp.Pop()
			field[to].Push(t)
		}

	}
	ans := make([]rune, 0)
	for _, st := range field {
		ans = append(ans, st.Pop())
	}

	return string(ans)
}
func main() {
	fmt.Println("part one ans is:", partOne())
	fmt.Println("part two ans is:", partTwo())
}
