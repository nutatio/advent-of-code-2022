package advc

import (
	"bufio"
	"os"
	"strings"
)

func Check(e error) {
	if e != nil {
		panic(e)
	}
}
func GetLines() []string {
	file, err := os.Open("input.txt")

	Check(err)
	defer file.Close()

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)
	var lines []string

	for fileScanner.Scan() {
		lines = append(lines, fileScanner.Text())
	}
	return lines
}

// day3 stuff
func GetPriority(letter rune) int {
	dict := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	return strings.IndexRune(dict, letter) + 1
}

func GetLetter(s1, s2 string) rune {
	for _, ch := range s1 {
		for _, ch2 := range s2 {
			if ch == ch2 {
				return ch
			}
		}
	}
	return '0'
}
