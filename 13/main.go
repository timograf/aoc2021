package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

var folds []fold
var mapp [][]bool
var coords []coord

type coord struct {
	x int
	y int
}

type fold struct {
	direction string
	line      int
}

func main() {
	fillMapp()
	fmt.Println(len(mapp), len(mapp[0]))
	/*for i, f := range folds {
		foldIt(f)
		if i == 0 {
			c := count()
			fmt.Println("one", c)
		}
	}*/
	foldIt(folds[0])
	foldIt(folds[1])

	/*for _, y := range mapp {
		for _, x := range y {
			if x {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Print("\n")
	}*/
}

func foldIt(f fold) {
	if f.direction == "y" {
		newMapp := make([][]bool, f.line)

		for i := 0; i < f.line; i++ {
			newMapp[i] = make([]bool, len(mapp[i]))
		}

		fmt.Println(len(newMapp), len(newMapp[0]))

		for y := 0; y < f.line; y++ {
			for x := 0; x < len(newMapp[y]); x++ {
				if x == 0 {
					fmt.Println("one", len(mapp)-y-1, x)
					fmt.Println("two", y, x)
				}
				if mapp[y][x] || mapp[len(mapp)-y-1][x] {
					newMapp[y][x] = true
				}
			}
		}

		mapp = newMapp
	} else {
		newMapp := make([][]bool, len(mapp))
		for i := 0; i < len(mapp); i++ {
			newMapp[i] = make([]bool, f.line)
		}

		fmt.Println(len(newMapp), len(newMapp[0]))
		for y := 0; y < len(newMapp); y++ {
			for x := 0; x < len(newMapp[y]); x++ {
				if y == 0 && x < 100000 {
					//fmt.Println("one", y, len(mapp[y])-x-1)
					//fmt.Println("two", y, x)
				}
				if mapp[y][x] || mapp[y][len(mapp[y])-x-1] {
					newMapp[y][x] = true
				}
			}
		}

		mapp = newMapp
	}
}

func count() int {
	count := 0
	for _, x := range mapp {
		for _, y := range x {
			if y {
				count++
			}
		}
	}
	return count
}

func fillMapp() {
	for _, c := range coords {
		mapp[c.x][c.y] = true
	}
}

func init() {
	input, err := ioutil.ReadFile("./13/input.txt")
	if err != nil {
		panic(err)
	}

	lns := strings.Split(string(input), "\n")
	coords = make([]coord, 0, len(lns))
	highestX := 0
	highestY := 0

	for _, l := range lns {
		nums := strings.Split(l, ",")
		var coord coord
		num2, err := strconv.Atoi(nums[0])
		num1, err := strconv.Atoi(nums[1])
		if err != nil {
			panic(err)
		}

		if num1 > highestX {
			highestX = num1
		}
		if num2 > highestY {
			highestY = num2
		}

		coord.x, coord.y = num1, num2
		coords = append(coords, coord)
	}

	mapp = make([][]bool, highestX+1)
	for i := 0; i < (highestX + 1); i++ {
		mapp[i] = make([]bool, highestY+1)
	}

	folds = []fold{
		{
			direction: "x",
			line:      655,
		},
		{
			direction: "y",
			line:      447,
		},
		{
			direction: "x",
			line:      327,
		},
		{
			direction: "y",
			line:      223,
		},
		{
			direction: "x",
			line:      163,
		},
		{
			direction: "y",
			line:      111,
		},
		{
			direction: "x",
			line:      81,
		},
		{
			direction: "y",
			line:      55,
		},
		{
			direction: "x",
			line:      40,
		},
		{
			direction: "y",
			line:      27,
		},
		{
			direction: "y",
			line:      13,
		},
		{
			direction: "y",
			line:      6,
		},
	}
}
