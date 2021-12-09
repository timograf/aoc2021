package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

var mapp [][]int
var lines []line

func main() {
	for _, line := range lines {
		if line.fromX == line.toX {
			var biggerY int
			var smallerY int

			if line.fromY > line.toY {
				biggerY = line.fromY
				smallerY = line.toY
			} else {
				biggerY = line.toY
				smallerY = line.fromY
			}

			for i := smallerY; i <= biggerY; i++ {
				mapp[line.fromX][i]++
			}
		} else if line.fromY == line.toY {
			var biggerX int
			var smallerX int

			if line.fromX > line.toX {
				biggerX = line.fromX
				smallerX = line.toX
			} else {
				biggerX = line.toX
				smallerX = line.fromX
			}

			for i := smallerX; i <= biggerX; i++ {
				mapp[i][line.fromY]++
			}
		} else {
			if line.fromX > line.toX {
				y := line.toY
				for i := line.toX; i <= line.fromX; i++ {
					mapp[i][y]++

					if line.toY < line.fromY {
						y++
					} else {
						y--
					}
				}
			} else {
				y := line.fromY
				for i := line.fromX; i <= line.toX; i++ {
					mapp[i][y]++

					if line.toY > line.fromY {
						y++
					} else {
						y--
					}
				}
			}
		}

	}

	count := 0
	for _, x := range mapp {
		for _, y := range x {
			if y > 1 {
				count++
			}
		}
	}

	fmt.Println("PART DUETES:", count)
}

func init() {
	input, err := ioutil.ReadFile("./5/input.txt")
	if err != nil {
		panic(err)
	}

	lns := strings.Split(string(input), "\n")
	lines = make([]line, 0, len(lns))
	highestNum := 0

	for _, l := range lns {
		nums := extractCoords(l)

		for _, num := range nums {
			if num > highestNum {
				highestNum = num
			}
		}

		var line line
		line.fromX, line.fromY, line.toX, line.toY = nums[0], nums[1], nums[2], nums[3]

		lines = append(lines, line)
	}

	mapp = make([][]int, highestNum+1)
	for i := 0; i < (highestNum + 1); i++ {
		mapp[i] = make([]int, highestNum+1)
	}
}

func extractCoords(l string) [4]int {
	var nums [4]int

	sets := strings.Split(l, ";")
	num1, num2 := extractNum(sets[0])
	num3, num4 := extractNum(sets[1])
	nums[0], nums[1], nums[2], nums[3] = num1, num2, num3, num4

	return nums
}

func extractNum(set string) (int, int) {
	nums := strings.Split(set, ",")
	num1, err := strconv.Atoi(nums[0])
	num2, err := strconv.Atoi(nums[1])
	if err != nil {
		panic(err)
	}

	return num1, num2
}

type line struct {
	fromX int
	toX   int
	fromY int
	toY   int
}
