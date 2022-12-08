package main

import (
	"advc"
	"fmt"
	"strconv"
	"strings"
)

const (
	MAX_ALLOWED_SIZE = 100000
	NEEDED_SPACE     = 30000000
	TOTAL_SPACE      = 70000000
)

type node struct {
	name   string
	size   int
	isFile bool
	root   *node
	leafs  map[string]*node
}

func getDirSize(root node) int {
	var size int
	if root.isFile {
		return root.size
	}
	for _, d := range root.leafs {
		size += getDirSize(*d)
	}
	return size
}

func partOne() int {
	lines := advc.GetLines()
	var curDir *node

	dirs := make([]*node, 0)

	for _, line := range lines {
		lineSlice := strings.Fields(line)
		switch line[:4] {
		case "$ cd":
			if lineSlice[2] == ".." {
				curDir = curDir.root
			} else if lineSlice[2] == "/" {
				curDir = &node{"/", 0, false, nil, make(map[string]*node)}
			} else {
				curDir = curDir.leafs[lineSlice[2]]
			}

		case "dir ":
			curDir.leafs[lineSlice[1]] = &node{lineSlice[1], 0, false, curDir, make(map[string]*node)}
			dirs = append(dirs, curDir.leafs[lineSlice[1]])
		default:
			if lineSlice[0] != "$" {
				size, err := strconv.Atoi(lineSlice[0])
				if err != nil {
					fmt.Println(err)
				}
				curDir.leafs[lineSlice[1]] = &node{lineSlice[1], size, true, curDir, nil}
			}
		}
	}

	var ans int

	for _, dir := range dirs {
		size := getDirSize(*dir)
		if size <= MAX_ALLOWED_SIZE {
			ans += size
		}
	}
	return ans
}

func partTwo() int {
	lines := advc.GetLines()
	var curDir *node

	dirs := make([]*node, 0)

	for _, line := range lines {
		lineSlice := strings.Fields(line)
		switch line[:4] {
		case "$ cd":
			if lineSlice[2] == ".." {
				curDir = curDir.root
			} else if lineSlice[2] == "/" {
				curDir = &node{"/", 0, false, nil, make(map[string]*node)}
				dirs = append(dirs, curDir)

			} else {
				curDir = curDir.leafs[lineSlice[2]]
			}

		case "dir ":
			curDir.leafs[lineSlice[1]] = &node{lineSlice[1], 0, false, curDir, make(map[string]*node)}
			dirs = append(dirs, curDir.leafs[lineSlice[1]])
		default:
			if lineSlice[0] != "$" {
				size, err := strconv.Atoi(lineSlice[0])
				if err != nil {
					fmt.Println(err)
				}
				curDir.leafs[lineSlice[1]] = &node{lineSlice[1], size, true, curDir, nil}
			}
		}
	}

	toFree := NEEDED_SPACE - (TOTAL_SPACE - getDirSize(*dirs[0]))
	ans := getDirSize(*dirs[0])

	for _, dir := range dirs {
		size := getDirSize(*dir)
		if size > toFree && size-toFree < ans-toFree {
			ans = size
		}
	}
	return ans
}
func main() {
	fmt.Println("part one ans:", partOne())
	fmt.Println("part two ans:", partTwo())
}
