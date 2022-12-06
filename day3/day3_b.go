package main

import (
	"advc"
	"fmt"
)

func getMatch(s1, s2, s3 string) rune {
	t := make([]rune, 0)

	for _, c1 := range s1 {
		for _, c2 := range s2 {
			if c1 == c2 {
				t = append(t, c1)
			}
		}
	}
	return advc.GetLetter(string(t), s3)

}

func main() {
	lines := advc.GetLines()
	var sum int
	for i := 0; i < len(lines); i += 3 {
		letter := getMatch(lines[i], lines[i+1], lines[i+2])
		sum += advc.GetPriority(letter)
	}

	fmt.Println(sum)
}
