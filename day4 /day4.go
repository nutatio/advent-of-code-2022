package main

import (
	"advc"
	"fmt"
	"strconv"
	"strings"
)

func contains(leftX, rightX, leftY, rightY int) bool {
	if leftY >= leftX && rightY <= rightX ||
		leftX >= leftY && rightX <= rightY {
		return true
	}
	return false
}

func overlaps(leftX, rightX, leftY, rightY int) bool {
	if leftX >= leftY && leftX <= rightY ||
		leftY >= leftX && leftY <= rightX {
		return true
	}
	return false
}
func atoi(s string) int {
	a, _ := strconv.Atoi(s)
	return a
}

func main() {
	lines := advc.GetLines()
	var ansA, ansB int
	for _, line := range lines {
		ids := strings.Split(line, ",")
		x, y := strings.Split(ids[0], "-"), strings.Split(ids[1], "-")
		leftX, rightX, leftY, rightY := atoi(x[0]), atoi(x[1]), atoi(y[0]), atoi(y[1])
		if contains(leftX, rightX, leftY, rightY) {
			ansA++
		}
		if overlaps(leftX, rightX, leftY, rightY) {
			ansB++
		}
	}

	fmt.Println("contains:", ansA, "\noverlaps:", ansB)
}
