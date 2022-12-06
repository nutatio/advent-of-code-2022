package main

import (
	"advc"
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// TODO add input.txt
// TODO refactor to have main
func stackGet(s [3]int) int {
	res := s[0]
	for i := 0; i < len(s); i++ {
		if res > s[i] {
			res = s[i]
		}
	}
	return res
}

func stackReplace(s *[3]int, toDel int, toIns int) {
	for i := 0; i < len(s); i++ {
		if s[i] == toDel {
			s[i] = toIns
			return
		}
	}
}

func DayOnePartTwo() {
	file, err := os.Open("input.txt")

	advc.Check(err)
	defer file.Close()

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)
	var lines []string

	for fileScanner.Scan() {
		lines = append(lines, fileScanner.Text())
	}

	var stack [3]int
	var cur int

	for _, line := range lines {
		if line == "" {
			n := stackGet(stack)
			if n < cur {
				stackReplace(&stack, n, cur)
			}
			cur = 0
		} else {
			n, err := strconv.Atoi(line)
			advc.Check(err)
			cur += n
		}
	}
	fmt.Println(stack[0] + stack[1] + stack[2])

}
