package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

var mapp [][]int

func main() {
	sum := 0
	for i, x := range mapp {
		for j, num := range x {
			cond1 := true
			cond2 := true
			cond3 := true
			cond4 := true

			xs, ys := surrounding(i, j)

			if len(xs) > 0 {
				cond1 = num < mapp[xs[0]][ys[0]]
			}
			if len(xs) > 1 {
				cond2 = num < mapp[xs[1]][ys[1]]
			}
			if len(xs) > 2 {
				cond3 = num < mapp[xs[2]][ys[2]]
			}
			if len(xs) > 3 {
				cond4 = num < mapp[xs[3]][ys[3]]
			}
			if cond1 && cond2 && cond3 && cond4 {
				sum++
				sum += num

				expand(i, j)
			}

		}
	}
	fmt.Println("one", sum)
}

func surrounding(x, y int) (xs []int, ys []int) {
	if x != 0 {
		xs = append(xs, x-1)
		ys = append(ys, y)
	}
	if x != len(mapp)-1 {
		xs = append(xs, x+1)
		ys = append(ys, y)
	}
	if y != 0 {
		xs = append(xs, x)
		ys = append(ys, y-1)
	}
	if y != len(mapp[x])-1 {
		xs = append(xs, x)
		ys = append(ys, y+1)
	}

	return
}

func expand(x, y int) (int, bool) {
	xs, ys := surrounding(x, y)

	for i, x := range xs {
		sum, done := expand(x, ys[i])
	}
}

func init() {
	mapp = make([][]int, 100)
	for i, _ := range mapp {
		mapp[i] = make([]int, 0)
	}
	input, err := ioutil.ReadFile("./9/input.txt")
	if err != nil {
		panic(err)
	}

	lns := strings.Split(string(input), "\n")

	for i, ln := range lns {
		nums := strings.Split(ln, "")
		for _, num := range nums {
			num, err := strconv.Atoi(num)
			if err != nil {
				panic(err)
			}
			mapp[i] = append(mapp[i], num)
		}
	}
}
