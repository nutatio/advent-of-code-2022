package main

import (
	"advc"
	"fmt"
)

func main() {
	lines := advc.GetLines()
	var sum int
	for _, line := range lines {
		first, second := line[:len(line)/2], line[len(line)/2:]
		letter := advc.GetLetter(first, second)
		sum += advc.GetPriority(letter)
	}

	fmt.Println(sum)
}
